package masterdb

// import (
// 	"database/sql"
// 	"errors"
// 	"fmt"
// 	"net/url"
// 	"time"
// )

// type AssetQueryBuilderParams struct {
// 	ChainId       int32      `json:"chainId"`
// 	CollectionId  *string    `json:"collectionId,omitempty"`
// 	TokenIds      []string   `json:"tokenIds,omitempty"`
// 	Owner         *string    `json:"owner,omitempty"`
// 	CreatedAtFrom *time.Time `json:"createdAtFrom,omitempty"`
// 	CreatedAtTo   *time.Time `json:"createdAtTo,omitempty"`
// 	Page          int        `json:"page"`
// 	Limit         int        `json:"limit"`
// 	Offset        int        `json:"offset"`
// }

// type masterDbClient struct {
// 	localDb     *sql.DB
// 	masterDbUrl string
// 	useMasterDb bool
// }

// // NewMasterDbConfig creates a new instance of masterDbConfig with validation
// func NewMasterDbConfig(
// 	localDb *sql.DB,
// 	masterDbUrl string,
// 	useMasterDb bool,
// ) (*masterDbClient, error) {

// 	if err := validateDbUrl(masterDbUrl); err != nil {
// 		return nil, fmt.Errorf("invalid master URL: %w", err)
// 	}

// 	return &masterDbClient{
// 		localDb:     localDb,
// 		masterDbUrl: masterDbUrl,
// 		useMasterDb: useMasterDb,
// 	}, nil
// }

// // CreateQueryBuilder creates a new AssetQueryBuilder instance
// func (c *masterDbClient) CreateQueryBuilder() AssetQueryBuilder {
// 	return &assetQueryBuilderParam{config: c}
// }

// // validateDbUrl checks if the provided URL is valid
// func validateDbUrl(dbUrl string) error {
// 	if dbUrl == "" {
// 		return errors.New("database URL cannot be empty")
// 	}

// 	_, err := url.Parse(dbUrl)
// 	if err != nil {
// 		return fmt.Errorf("invalid URL format: %w", err)
// 	}

// 	return nil
// }

// type assetQueryBuilderParam struct {
// 	chainId       int32
// 	collectionId  *string
// 	tokenIds      *[]string
// 	owner         *string
// 	createdAtFrom *time.Time
// 	createdAtTo   *time.Time
// 	page          *int
// 	limit         *int
// 	offset        *int
// 	config        *masterDbClient
// }

// type AssetQueryFunction interface {
// 	GetAssetQueryBuilder() (*assetQueryBuilderParam, error)
// 	GetPaginatedAsset() (any, error)
// }

// type AssetQueryBuilder interface {
// 	WithChainId(chainId int32) AssetQueryBuilder
// 	WithCollectionId(collectionId string) AssetQueryBuilder
// 	WithTokenIds(tokenIds []string) AssetQueryBuilder
// 	WithOwner(owner string) AssetQueryBuilder
// 	WithCreatedAtFrom(createdAtFrom time.Time) AssetQueryBuilder
// 	WithCreatedAtTo(createdAtTo time.Time) AssetQueryBuilder
// 	WithPage(page int) AssetQueryBuilder
// 	WithLimit(limit int) AssetQueryBuilder
// 	Build() AssetQueryFunction
// }

// func (b *assetQueryBuilderParam) getHttpClient() *HttpClient {
// 	client := NewHttpClient(b.config.masterDbUrl)
// 	return client
// }

// // // WithChainId implements AssetQueryBuilder.
// // func (b *assetQueryBuilderParam) getCollectionType() masterDbCommon.CollectionType {
// // 	client := b.getHttpClient()

// // 	var response response.HTTPResponse[masterDbCommon.CollectionResponse]
// // 	path := fmt.Sprintf("/chain/%d/collection/%s", b.chainId, *b.collectionId)
// // 	err := client.DoRequest(context.Background(), "GET", path, nil, &response)
// // 	if err != nil {
// // 		fmt.Println("err", err)
// // 		return masterDbCommon.CollectionType("")
// // 	}

// // 	return response.Data.Type
// // }

// // WithChainId implements AssetQueryBuilder.
// func (b *assetQueryBuilderParam) WithChainId(chainId int32) AssetQueryBuilder {
// 	b.chainId = chainId
// 	return b
// }

// // WithCollectionId implements AssetQueryBuilder.
// func (b *assetQueryBuilderParam) WithCollectionId(collectionId string) AssetQueryBuilder {
// 	b.collectionId = &collectionId
// 	return b
// }

// // WithOwner implements AssetQueryBuilder.
// func (b *assetQueryBuilderParam) WithOwner(owner string) AssetQueryBuilder {
// 	b.owner = &owner
// 	return b
// }

// // WithTokenId implements AssetQueryBuilder.
// func (b *assetQueryBuilderParam) WithTokenIds(tokenIds []string) AssetQueryBuilder {
// 	b.tokenIds = &tokenIds
// 	return b
// }

// // WithCreatedAtFrom implements AssetQueryBuilder.
// func (b *assetQueryBuilderParam) WithCreatedAtFrom(createdAtFrom time.Time) AssetQueryBuilder {
// 	b.createdAtFrom = &createdAtFrom
// 	return b
// }

// // WithCreatedAtTo implements AssetQueryBuilder.
// func (b *assetQueryBuilderParam) WithCreatedAtTo(createdAtTo time.Time) AssetQueryBuilder {
// 	b.createdAtTo = &createdAtTo
// 	return b
// }

// // WithPage implements AssetQueryBuilder.
// func (b *assetQueryBuilderParam) WithPage(page int) AssetQueryBuilder {
// 	b.page = &page
// 	return b
// }

// // WithLimit implements AssetQueryBuilder.
// func (b *assetQueryBuilderParam) WithLimit(limit int) AssetQueryBuilder {
// 	b.limit = &limit
// 	return b
// }

// func (b *assetQueryBuilderParam) Build() AssetQueryFunction {
// 	// Set default values if not provided
// 	defaultPage := 1
// 	defaultLimit := 10

// 	if b.page != nil {
// 		if *b.page < 1 {
// 			return nil
// 		}
// 		defaultPage = *b.page
// 	}

// 	if b.limit != nil {
// 		if *b.limit > 100 {
// 			return nil
// 		}
// 		defaultLimit = *b.limit
// 	}

// 	// Calculate offset
// 	offset := (defaultPage - 1) * defaultLimit
// 	b.offset = &offset
// 	return b
// }
