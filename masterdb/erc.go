package masterdb

// Add721Asset represents the structure for adding a 721 asset.
type Add721Asset struct {
	Asset721 Asset721 `json:"asset721"`
	History  History  `json:"history"`
}

// Add1155Asset represents the structure for adding a 1155 asset.
type Add1155Asset struct {
	Asset1155From Asset1155 `json:"asset1155from"`
	Asset1155To   Asset1155 `json:"asset1155to"`
	History       History   `json:"history"`
}

// Add20Asset represents the structure for adding a 20 asset.
type Add20Asset struct {
	Asset20From Asset1155 `json:"asset20from"`
	Asset20To   Asset1155 `json:"asset20to"`
	History     History   `json:"history"`
}

// History holds the transaction history information.
type History struct {
	From         string `json:"from"`
	To           string `json:"to"`
	CollectionId string `json:"collectionId"`
	TokenId      string `json:"tokenId"`
	Amount       int    `json:"amount"`
	TxHash       string `json:"txHash"`
	Signature    string `json:"signature"`
}

// Asset721 holds the information specific to the 721 asset.
type Asset721 struct {
	ChainId      int32  `json:"chainId"`
	CollectionId string `json:"collectionId"`
	TokenId      string `json:"tokenId"`
	Owner        string `json:"owner"`
	Attributes   string `json:"attributes"`
	Signature    string `json:"signature"`
}

// Asset1155 holds the information specific to the 1155 asset.
type Asset1155 struct {
	ChainId      int32  `json:"chainId"`
	CollectionId string `json:"collectionId"`
	TokenId      string `json:"tokenId"`
	Owner        string `json:"owner"`
	Balance      string `json:"balance"`
	Attributes   string `json:"attributes"`
	Signature    string `json:"signature"`
}

// Asset20 holds the information specific to the 20 asset.
type Asset20 struct {
	ChainId      int32  `json:"chainId"`
	CollectionId string `json:"collectionId"`
	Owner        string `json:"owner"`
	Balance      string `json:"balance"`
	Signature    string `json:"signature"`
}
