package runtime

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ContractCallParams struct {
	Sender          string
	ContractAddress string
	Abi             string
	Method          string
	Params          []interface{}
	Value           string
}

type TransactionRequest struct {
	From  common.Address
	To    common.Address
	Data  string
	Value *big.Int
}

type OnchainTransactionPayload struct {
	To                   string `json:"to"`
	Value                string `json:"value"`
	Data                 string `json:"data"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
}

type OnchainTransactionRequest struct {
	ProjectID      string                     `json:"projectId"`
	ChainID        int                        `json:"chainId"`
	Sponsor        bool                       `json:"sponsor"`
	TransactionReq *OnchainTransactionPayload `json:"transactionReq"`
}
