package runtime

type WalletUpdate struct {
	UserID    string
	Changeset map[string]int64
	Metadata  map[string]interface{}
}

type WalletUpdateResult struct {
	UserID   string
	Updated  map[string]int64
	Previous map[string]int64
}

type WalletNegativeError struct {
	UserID  string
	Path    string
	Current int64
	Amount  int64
}

type WalletLedgerItem interface {
	GetID() string
	GetUserID() string
	GetCreateTime() int64
	GetUpdateTime() int64
	GetChangeset() map[string]int64
	GetMetadata() map[string]interface{}
}
