package masterdb

import "github.com/google/uuid"

// Add721Asset represents the structure for adding a 721 asset.
type Add721Asset struct {
	Asset721 Asset721 `json:"asset721"`
	History  History  `json:"history"`
}

type Add721AssetBatch struct {
	Assets721 []Asset721 `json:"assets721"`
	// History   *History      `json:"history,omitempty"`
}

// Add1155Asset represents the structure for adding a 1155 asset.
type Add1155Asset struct {
	Asset1155From Asset1155 `json:"asset1155from"`
	Asset1155To   Asset1155 `json:"asset1155to"`
	History       History   `json:"history"`
}

type Add1155AssetBatch struct {
	Assets1155 []Asset1155 `json:"assets1155"`
	// History   *History      `json:"history,omitempty"`
}

// Add20Asset represents the structure for adding a 20 asset.
type Add20Asset struct {
	Asset20From Asset20 `json:"asset20from"`
	Asset20To   Asset20 `json:"asset20to"`
	History     History `json:"history"`
}

type Add20AssetBatch struct {
	Assets20 []Asset20 `json:"assets20"`
	// History   *History      `json:"history,omitempty"`
}

// History holds the transaction history information.
type History struct {
	From         string `json:"from"`
	To           string `json:"to"`
	CollectionId string `json:"collectionId"`
	TokenId      string `json:"tokenId"`
	Amount       string `json:"amount"`
	TxHash       string `json:"txHash"`
}

// Asset721 holds the information specific to the 721 asset.
type Asset721 struct {
	ChainId      int32  `json:"chainId"`
	CollectionId string `json:"collectionId"`
	TokenId      string `json:"tokenId"`
	Owner        string `json:"owner"`
	Attributes   string `json:"attributes"`
}

// Asset1155 holds the information specific to the 1155 asset.
type Asset1155 struct {
	ChainId      int32  `json:"chainId"`
	CollectionId string `json:"collectionId"`
	TokenId      string `json:"tokenId"`
	Owner        string `json:"owner"`
	Balance      string `json:"balance"`
	Attributes   string `json:"attributes"`
}

// Asset20 holds the information specific to the 20 asset.
type Asset20 struct {
	ChainId      int32  `json:"chainId"`
	CollectionId string `json:"collectionId"`
	Owner        string `json:"owner"`
	Balance      string `json:"balance"`
}

type VerifyRandomAsset struct {
	AssetType CollectionType `json:"assetType"`
	AssetId   uuid.UUID      `json:"assetId"`
}
