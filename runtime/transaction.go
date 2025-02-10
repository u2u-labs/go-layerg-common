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
	ChainID        int                 `json:"chainId"`
	Sponsor        bool                `json:"sponsor"`
	TransactionReq *TransactionRequest `json:"transactionReq"`
}

type UATransactionResponse struct {
	Success bool   `json:"success"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Amount     string `json:"amount"`
	UserOpHash string `json:"userOpHash"`
	User       User   `json:"user"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	ID         string `json:"id"`
}

type User struct {
	ID                string  `json:"id"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         string  `json:"updatedAt"`
	Type              string  `json:"type"`
	EoaWallet         string  `json:"eoaWallet"`
	Username          *string `json:"username"`
	Firstname         *string `json:"firstname"`
	Lastname          *string `json:"lastname"`
	Avatar            *string `json:"avatar"`
	Signature         *string `json:"signature"`
	PrimaryAAWallet   *string `json:"primaryAAWallet"`
	TelegramID        string  `json:"telegramId"`
	TelegramUsername  string  `json:"telegramUsername"`
	TelegramFirstName string  `json:"telegramFirstName"`
	TelegramLastName  string  `json:"telegramLastName"`
	TelegramAvatarURL string  `json:"telegramAvatarUrl"`
	TelegramAuthDate  string  `json:"telegramAuthDate"`
	TelegramVerified  bool    `json:"telegramVerified"`
}

type UARefreshTokenResponse struct {
	Success bool `json:"success"`
	Data    struct {
		RefreshToken       string `json:"refreshToken"`
		RefreshTokenExpire int64  `json:"refreshTokenExpire"`
		AccessToken        string `json:"accessToken"`
		AccessTokenExpire  int64  `json:"accessTokenExpire"`
		UserID             int    `json:"userId"`
		ApiKey             string `json:"apiKey"`
	} `json:"data"`
	Message string `json:"message"`
}

type UARefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
