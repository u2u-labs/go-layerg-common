package runtime

type ContractCallParams struct {
	Sender          string
	ContractAddress string
	Abi             string
	Method          string
	Params          []interface{}
	Value           string
}

type TransactionRequest struct {
	Type                 *int          `json:"type,omitempty"`
	To                   *string       `json:"to,omitempty"`
	From                 *string       `json:"from,omitempty"`
	Nonce                *int64        `json:"nonce,omitempty"`
	GasLimit             *string       `json:"gasLimit,omitempty"`
	GasPrice             *string       `json:"gasPrice,omitempty"`
	MaxPriorityFeePerGas *string       `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerGas         *string       `json:"maxFeePerGas,omitempty"`
	Data                 *string       `json:"data,omitempty"`
	Value                *string       `json:"value,omitempty"`
	ChainID              *string       `json:"chainId,omitempty"`
	AccessList           *[]AccessItem `json:"accessList,omitempty"`
	CustomData           interface{}   `json:"customData,omitempty"`
	BlockTag             *string       `json:"blockTag,omitempty"`
	EnableCcipRead       *bool         `json:"enableCcipRead,omitempty"`
	BlobVersionedHashes  *[]string     `json:"blobVersionedHashes,omitempty"`
	MaxFeePerBlobGas     *string       `json:"maxFeePerBlobGas,omitempty"`
	Blobs                *[]string     `json:"blobs,omitempty"`
	Kzg                  interface{}   `json:"kzg,omitempty"`
}

type AccessItem struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

// type OnchainTransactionPayload struct {
// 	To                   string `json:"to"`
// 	Value                string `json:"value"`
// 	Data                 string `json:"data"`
// 	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
// }

type UATransactionRequest struct {
	ProjectID      int                 `json:"projectId"`
	ChainID        int                 `json:"chainId"`
	Sponsor        bool                `json:"sponsor"`
	TransactionReq *TransactionRequest `json:"transactionReq"`
}
