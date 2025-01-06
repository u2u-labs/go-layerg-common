package interfaces

type InsertErc721Asset struct {
	ChainID      int    `json:"chain_id"`
	CollectionID string `json:"collection_id"`
	TokenID      string `json:"token_id"` // Use string for DECIMAL(78, 0)
	Owner        string `json:"owner"`
	Attributes   string `json:"attributes"`
	Signature    string `json:"signature"`
}

type InsertErc1155Asset struct {
	ChainID      int    `json:"chain_id"`
	CollectionID string `json:"collection_id"`
	TokenID      string `json:"token_id"` // Use string for DECIMAL(78, 0)
	Owner        string `json:"owner"`
	Balance      string `json:"balance"` // Use string for DECIMAL(78, 0)
	Attributes   string `json:"attributes"`
	Signature    string `json:"signature"`
}

type InsertErc20Asset struct {
	ChainID      int    `json:"chain_id"`
	CollectionID string `json:"collection_id"`
	TokenID      string `json:"token_id"` // Use string for DECIMAL(78, 0)
	Owner        string `json:"owner"`
	Balance      string `json:"balance"` // Use string for DECIMAL(78, 0)
	Signature    string `json:"signature"`
}
