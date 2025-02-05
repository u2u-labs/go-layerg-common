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
