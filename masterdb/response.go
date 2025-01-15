package masterdb

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CollectionType string

const (
	CollectionTypeERC721  CollectionType = "ERC721"
	CollectionTypeERC1155 CollectionType = "ERC1155"
	CollectionTypeERC20   CollectionType = "ERC20"
)

type CollectionResponse struct {
	ID                string         `json:"id"`
	ChainID           int32          `json:"chainId"`
	CollectionAddress string         `json:"collectionAddress"`
	Type              CollectionType `json:"type"`
	DecimalData       int            `json:"decimalData"`  // Output as int
	InitialBlock      int64          `json:"initialBlock"` // Output as int64
	LastUpdated       time.Time      `json:"lastUpdated"`  // Output as time.Time
}

type ChainResponse struct {
	ID          int32  `json:"id"`
	Chain       string `json:"chain"`
	Name        string `json:"name"`
	RpcUrl      string `json:"rpcUrl"`
	ChainID     int64  `json:"chainId"`
	Explorer    string `json:"explorer"`
	LatestBlock int64  `json:"latestBlock"`
	BlockTime   int32  `json:"blockTime"`
}

type Erc721CollectionAssetResponse struct {
	ID           uuid.UUID `json:"id"`
	ChainID      int32     `json:"chainId"`
	CollectionID string    `json:"collectionId"`
	TokenID      string    `json:"tokenId"`
	Owner        string    `json:"owner"`
	Attributes   string    `json:"attributes"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    uuid.UUID `json:"updatedBy"`
}

type Erc1155CollectionAssetResponse struct {
	ID           uuid.UUID `json:"id"`
	ChainID      int32     `json:"chainId"`
	CollectionID string    `json:"collectionId"`
	TokenID      string    `json:"tokenId"`
	Owner        string    `json:"owner"`
	Balance      string    `json:"balance"`
	Attributes   string    `json:"attributes"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    uuid.UUID `json:"updatedBy"`
}

type Erc20CollectionAssetResponse struct {
	ID           uuid.UUID `json:"id"`
	ChainID      int32     `json:"chainId"`
	CollectionID string    `json:"collectionId"`
	Owner        string    `json:"owner"`
	Balance      string    `json:"balance"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    uuid.UUID `json:"updatedBy"`
}

type RandomAssetData struct {
	ID           uuid.UUID `json:"id"`
	ChainID      int32     `json:"chainId"`
	CollectionID string    `json:"collectionId"`
	Owner        string    `json:"owner"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    uuid.UUID `json:"updatedBy"`
	IsVerified   bool      `json:"isVerified"`
}

type RandomAssetResponse struct {
	AssetType CollectionType  `json:"assetType"`
	AssetData RandomAssetData `json:"assetData"`
}

type Pagination[T any] struct {
	Items      []T   `json:"items"`
	TotalItems int64 `json:"totalItems"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"totalPages"`
}

// convert type

type Erc721CollectionAssetSql struct {
	ID           uuid.UUID      `json:"id"`
	ChainID      int32          `json:"chainId"`
	CollectionID string         `json:"collectionId"`
	TokenID      string         `json:"tokenId"`
	Owner        string         `json:"owner"`
	Attributes   sql.NullString `json:"attributes"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	UpdatedBy    uuid.UUID      `json:"updatedBy"`
	Signature    string         `json:"signature"`
}

// func ConvertSqlErc721AssetsToResponse(assets []Erc721CollectionAssetSql) []Erc721CollectionAssetResponse {
// 	var response []Erc721CollectionAssetResponse
// 	for _, item := range assets {
// 		r := ConvertSqlErc721AssetToResponse(item)
// 		response = append(response, r)
// 	}
// 	return response
// }

// func ConvertSqlErc721AssetToResponse(asset db.Erc721CollectionAsset) types.Erc721CollectionAssetResponse {
// 	attributes := ""
// 	if asset.Attributes.Valid {
// 		attributes = asset.Attributes.String
// 	}

// 	return types.Erc721CollectionAssetResponse{
// 		ID:           asset.ID,
// 		ChainID:      asset.ChainID,
// 		CollectionID: asset.CollectionID,
// 		TokenID:      asset.TokenID,
// 		Owner:        asset.Owner,
// 		Attributes:   attributes,
// 		CreatedAt:    asset.CreatedAt,
// 		UpdatedAt:    asset.UpdatedAt,
// 		UpdatedBy:    asset.UpdatedBy,
// 	}
// }
