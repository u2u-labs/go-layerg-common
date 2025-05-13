package runtime

import "google.golang.org/protobuf/types/known/timestamppb"

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

type Notification struct {
	Id         string
	UserID     string
	Subject    string
	Content    map[string]any
	Code       int
	Sender     string
	CreateTime *timestamppb.Timestamp
	Persistent bool
}

type NotificationUpdate struct {
	Id      string
	Subject *string
	Content map[string]any
	Sender  *string
}
