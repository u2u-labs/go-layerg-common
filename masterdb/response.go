package masterdb

import "time"

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
