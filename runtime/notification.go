package runtime

type NotificationSend struct {
	UserID     string
	Subject    string
	Content    map[string]interface{}
	Code       int
	Sender     string
	Persistent bool
}

type NotificationDelete struct {
	UserID         string
	NotificationID string
}
