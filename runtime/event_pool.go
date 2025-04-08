package runtime

import "context"

type EventQuery struct {
	ChainId         int    `json:"chainId"`
	ContractAddress string `json:"contractAddress"`
	TxHash          string `json:"txHash"`
}

type EventSubscription struct {
	ChainId         int    `json:"chainId"`
	ContractAddress string `json:"contractAddress"`
	EventName       string `json:"eventName"`
}

type EventHandler func(ctx context.Context, eventData EventData) error

type EventData struct {
	BlockNumber uint64 `json:"blockNumber"`
	Data        string `json:"data"`
	TxHash      string `json:"txHash"`
}

type EventResponse struct {
	Data []struct {
		ID          string `json:"id"`
		ContractID  string `json:"contractId"`
		BlockNumber int64  `json:"blockNumber"`
		TxHash      string `json:"txHash"`
		LogIndex    int    `json:"logIndex"`
		Data        string `json:"data"`
		CreatedAt   string `json:"createdAt"`
		Contract    struct {
			ID             string `json:"id"`
			ChainID        int    `json:"chainId"`
			EventName      string `json:"eventName"`
			Address        string `json:"address"`
			EventSignature string `json:"eventSignature"`
			StartBlock     int64  `json:"startBlock"`
			EventABI       string `json:"eventABI"`
			CreatedAt      string `json:"createdAt"`
			UpdatedAt      string `json:"updatedAt"`
		} `json:"contract"`
	} `json:"data"`
	Pagination struct {
		Skip int `json:"skip"`
		Take int `json:"take"`
	} `json:"pagination"`
}
