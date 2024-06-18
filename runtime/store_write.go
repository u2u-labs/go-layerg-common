package runtime

type StorageWrite struct {
	Collection      string
	Key             string
	UserID          string
	Value           string
	Version         string
	PermissionRead  int
	PermissionWrite int
}

type StorageDelete struct {
	Collection string
	Key        string
	UserID     string
	Version    string
}
