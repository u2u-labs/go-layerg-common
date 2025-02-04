package masterdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/Masterminds/squirrel"
)

func (b *assetQueryBuilderParam) getHttpClient() *HttpClient {
	client := NewHttpClient(b.config.masterDbUrl)
	return client
}

// // WithChainId implements AssetQueryBuilder.
func (b *assetQueryBuilderParam) getCollectionType() CollectionType {
	client := b.getHttpClient()

	var response HTTPResponse[CollectionResponse]
	path := fmt.Sprintf("/chain/%d/collection/%s", b.chainId, *b.collectionId)
	err := client.DoRequest(context.Background(), "GET", path, nil, &response)
	if err != nil {
		fmt.Println("err", err)
		return CollectionType("")
	}

	return response.Data.Type
}

// Helper function to convert string to snake_case
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

func (b *assetQueryBuilderParam) getFilterConditions() map[string][]string {
	filterConditions := make(map[string][]string)

	if b.collectionId != nil {
		filterConditions["collection_id"] = []string{*b.collectionId}
	}

	if b.tokenIds != nil && len(*b.tokenIds) > 0 {
		filterConditions["token_id"] = *b.tokenIds
	}

	if b.owner != nil {
		filterConditions["owner"] = []string{*b.owner}
	}

	if b.createdAtFrom != nil {
		filterConditions["created_at_from"] = []string{b.createdAtFrom.Format(time.RFC3339)}
	}

	if b.createdAtTo != nil {
		filterConditions["created_at_to"] = []string{b.createdAtTo.Format(time.RFC3339)}
	}

	return filterConditions
}

func (b *AssetQueryBuilderParam) GetFilterConditions() map[string][]string {
	filterConditions := make(map[string][]string)

	if b.CollectionId != "" {
		filterConditions["collection_id"] = []string{b.CollectionId}
	}
	if len(b.TokenIds) > 0 {
		filterConditions["token_id"] = b.TokenIds
	}

	if b.Owner != "" {
		filterConditions["owner"] = []string{b.Owner}
	}

	if !b.CreatedAtFrom.IsZero() {
		filterConditions["created_at_from"] = []string{b.CreatedAtFrom.Format(time.RFC3339)}
	}

	if !b.CreatedAtTo.IsZero() {
		filterConditions["created_at_to"] = []string{b.CreatedAtTo.Format(time.RFC3339)}
	}

	return filterConditions
}

// QueryWithDynamicFilter retrieves a slice of items from the database
func QueryWithDynamicFilter[T any](db *sql.DB, tableName string, limit int, offset int, filterConditions map[string][]string) ([]T, error) {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	queryBuilder := psql.Select("*").From(tableName)

	// Apply dynamic filters
	for column, values := range filterConditions {
		if column == "created_at_from" {
			fmt.Println("created_at_from", values[0])
			queryBuilder = queryBuilder.Where(squirrel.GtOrEq{"created_at": values[0]})
		} else if column == "created_at_to" {
			fmt.Println("created_at_to", values[0])
			queryBuilder = queryBuilder.Where(squirrel.LtOrEq{"created_at": values[0]})
		} else if len(values) == 1 {
			queryBuilder = queryBuilder.Where(squirrel.Eq{column: values[0]})
		} else if len(values) > 1 {
			queryBuilder = queryBuilder.Where(squirrel.Eq{column: values})
		}
	}

	// Apply pagination
	if limit > 0 {
		queryBuilder = queryBuilder.Limit(uint64(limit))
	}
	if offset > 0 {
		queryBuilder = queryBuilder.Offset(uint64(offset))
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building SQL query: %w", err)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("error getting columns: %w", err)
	}

	var items []T
	var item T
	itemType := reflect.TypeOf(item)
	if itemType.Kind() == reflect.Ptr {
		itemType = itemType.Elem()
	}

	// Create a map of DB column names to struct field indices
	columnMap := make(map[string]int)
	for i := 0; i < itemType.NumField(); i++ {
		field := itemType.Field(i)

		// Check for db tag first
		dbTag := field.Tag.Get("db")
		if dbTag != "" {
			columnMap[dbTag] = i
			continue
		}

		// Then check for json tag
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			// Split the json tag to handle options like omitempty
			tagParts := strings.Split(jsonTag, ",")
			columnMap[toSnakeCase(tagParts[0])] = i
			continue
		}

		// If no tags present, use field name in snake_case
		columnMap[toSnakeCase(field.Name)] = i
	}

	for rows.Next() {
		itemValue := reflect.New(itemType).Elem()
		scanArgs := make([]interface{}, len(columns))

		for i, colName := range columns {
			if fieldIndex, ok := columnMap[colName]; ok {
				field := itemValue.Field(fieldIndex)

				// Handle null values for pointer fields
				if field.Kind() == reflect.Ptr {
					if field.IsNil() {
						field.Set(reflect.New(field.Type().Elem()))
					}
					scanArgs[i] = field.Interface()
				} else {
					scanArgs[i] = field.Addr().Interface()
				}
			} else {
				var placeholder interface{}
				scanArgs[i] = &placeholder
			}
		}

		if err := rows.Scan(scanArgs...); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		items = append(items, itemValue.Interface().(T))
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return items, nil
}

// CountItems counts the number of items in the database based on dynamic filters.
func CountItemsWithFilter(db *sql.DB, tableName string, filterConditions map[string][]string) (int, int64, error) {
	// Create a Squirrel query builder for counting items
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	queryBuilder := psql.Select("COUNT(*)").From(tableName)

	// Create a Squirrel query builder for counting items holder
	psql2 := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	holderQueryBuilder := psql2.Select("COUNT(DISTINCT(owner))").From(tableName)

	// Apply dynamic filters
	for column, values := range filterConditions {
		if column == "created_at_from" {
			fmt.Println("created_at_from", values[0])
			queryBuilder = queryBuilder.Where(squirrel.GtOrEq{"created_at": values[0]})
		} else if column == "created_at_to" {
			fmt.Println("created_at_to", values[0])
			queryBuilder = queryBuilder.Where(squirrel.LtOrEq{"created_at": values[0]})
		} else if len(values) == 1 {
			queryBuilder = queryBuilder.Where(squirrel.Eq{column: values[0]})
		} else if len(values) > 1 {
			queryBuilder = queryBuilder.Where(squirrel.Eq{column: values})
		}
	}

	// Convert the query to SQL
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return 0, 0, err
	}

	holderQuery, holderArgs, err := holderQueryBuilder.ToSql()
	if err != nil {
		return 0, 0, err
	}

	// Execute the query
	var itemCount int
	var holderCount int64

	err = db.QueryRow(query, args...).Scan(&itemCount)
	if err != nil {
		return 0, 0, err
	}

	err = db.QueryRow(holderQuery, holderArgs...).Scan(&holderCount)
	if err != nil {
		return 0, 0, err
	}

	return itemCount, holderCount, nil
}

func (b *assetQueryBuilderParam) getLocalAssetQuery() (any, error) {
	collectionType := b.getCollectionType()
	fmt.Println("collectionType", collectionType)

	filterConditions := b.getFilterConditions()

	switch collectionType {
	case CollectionTypeERC721:
		totalAssets, _, err := CountItemsWithFilter(b.config.localDb, "erc_721_collection_assets", filterConditions)
		if err != nil {
			return nil, err
		}

		assets, _ := QueryWithDynamicFilter[Erc721CollectionAssetResponse](b.config.localDb, "erc_721_collection_assets", *b.limit, *b.offset, filterConditions)
		return Pagination[Erc721CollectionAssetResponse]{
			Page:       *b.page,
			Limit:      *b.limit,
			TotalItems: int64(totalAssets),
			TotalPages: int64((int64(totalAssets) + int64(*b.limit) - 1) / int64(*b.limit)),
			Data:       assets,
		}, nil

	case CollectionTypeERC1155:
		totalAssets, _, err := CountItemsWithFilter(b.config.localDb, "erc_1155_collection_assets", filterConditions)
		if err != nil {
			return nil, err
		}

		assets, _ := QueryWithDynamicFilter[Erc1155CollectionAssetResponse](b.config.localDb, "erc_1155_collection_assets", *b.limit, *b.offset, filterConditions)
		return Pagination[Erc1155CollectionAssetResponse]{
			Page:       *b.page,
			Limit:      *b.limit,
			TotalItems: int64(totalAssets),
			TotalPages: int64((int64(totalAssets) + int64(*b.limit) - 1) / int64(*b.limit)),
			Data:       assets,
		}, nil

	case CollectionTypeERC20:
		totalAssets, _, err := CountItemsWithFilter(b.config.localDb, "erc_20_collection_assets", filterConditions)
		if err != nil {
			return nil, err
		}

		assets, _ := QueryWithDynamicFilter[Erc20CollectionAssetResponse](b.config.localDb, "erc_20_collection_assets", *b.limit, *b.offset, filterConditions)
		return Pagination[Erc20CollectionAssetResponse]{
			Page:       *b.page,
			Limit:      *b.limit,
			TotalItems: int64(totalAssets),
			TotalPages: int64((int64(totalAssets) + int64(*b.limit) - 1) / int64(*b.limit)),
			Data:       assets,
		}, nil
	}

	return nil, nil
}

func (b *assetQueryBuilderParam) getMasterDbAsset() (any, error) {
	httpClient := b.getHttpClient()
	collectionType := b.getCollectionType()
	fmt.Println("collectionType", collectionType)

	var requestBody AssetQueryBuilderParam
	requestBody.ChainId = b.chainId
	if b.collectionId != nil {
		requestBody.CollectionId = *b.collectionId
	}
	if b.tokenIds != nil {
		requestBody.TokenIds = *b.tokenIds
	}
	if b.owner != nil {
		requestBody.Owner = *b.owner
	}
	if b.createdAtFrom != nil {
		requestBody.CreatedAtFrom = *b.createdAtFrom
	}
	if b.createdAtTo != nil {
		requestBody.CreatedAtTo = *b.createdAtTo
	}
	if b.limit != nil {
		requestBody.Limit = *b.limit
	}
	if b.offset != nil {
		requestBody.Offset = *b.offset
	}

	switch collectionType {
	case CollectionTypeERC721:
		var response HTTPResponse[Pagination[Erc721CollectionAssetResponse]]
		err := httpClient.DoRequest(context.Background(), "POST", "/query-builder", requestBody, &response)
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		fmt.Println("response", response.Data)
		return response.Data, nil
	case CollectionTypeERC1155:
		var response HTTPResponse[Pagination[Erc1155CollectionAssetResponse]]
		err := httpClient.DoRequest(context.Background(), "POST", "/query-builder", requestBody, &response)
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		fmt.Println("response", response.Data)
		return response.Data, nil
	case CollectionTypeERC20:
		var response HTTPResponse[Pagination[Erc20CollectionAssetResponse]]
		err := httpClient.DoRequest(context.Background(), "POST", "/query-builder", requestBody, &response)
		if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		fmt.Println("response", response.Data)
		return response.Data, nil
	}

	return nil, nil
}

// //
// / IMPLEMENT FUNCTION
// ///

func (b *assetQueryBuilderParam) GetAssetQueryBuilder() (*assetQueryBuilderParam, error) {
	if b.chainId == 0 {
		return nil, errors.New("chainId is required")
	}
	return b, nil
}

func (b *assetQueryBuilderParam) GetPaginatedAsset() (any, error) {

	if !b.config.useMasterDb {
		return b.getLocalAssetQuery()
	}
	return b.getMasterDbAsset()

}
