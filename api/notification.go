package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Delete one or more notifications for the current user.
type DeleteNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of notifications.
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteNotificationsRequest) Reset() {
	*x = DeleteNotificationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[32]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNotificationsRequest) ProtoMessage() {}

func (x *DeleteNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[32]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNotificationsRequest.ProtoReflect.Descriptor instead.
func (*DeleteNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{32}
}

func (x *DeleteNotificationsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// Get a list of unexpired notifications.
type ListNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of notifications to get. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// A cursor to page through notifications. May be cached by clients to get from point in time forwards.
	CacheableCursor string `protobuf:"bytes,2,opt,name=cacheable_cursor,json=cacheableCursor,proto3" json:"cacheable_cursor,omitempty"` // value from NotificationList.cacheable_cursor.
}

func (x *ListNotificationsRequest) Reset() {
	*x = ListNotificationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[65]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotificationsRequest) ProtoMessage() {}

func (x *ListNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[65]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotificationsRequest.ProtoReflect.Descriptor instead.
func (*ListNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{65}
}

func (x *ListNotificationsRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListNotificationsRequest) GetCacheableCursor() string {
	if x != nil {
		return x.CacheableCursor
	}
	return ""
}

// A notification in the server.
type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Notification.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Subject of the notification.
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// Content of the notification in JSON.
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Category code for this notification.
	Code int32 `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	// ID of the sender, if a user. Otherwise 'null'.
	SenderId string `protobuf:"bytes,5,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the notification was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// True if this notification was persisted to the database.
	Persistent bool `protobuf:"varint,7,opt,name=persistent,proto3" json:"persistent,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[74]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[74]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{74}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Notification) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Notification) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Notification) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Notification) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Notification) GetPersistent() bool {
	if x != nil {
		return x.Persistent
	}
	return false
}

// A collection of zero or more notifications.
type NotificationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Collection of notifications.
	Notifications []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	// Use this cursor to paginate notifications. Cache this to catch up to new notifications.
	CacheableCursor string `protobuf:"bytes,2,opt,name=cacheable_cursor,json=cacheableCursor,proto3" json:"cacheable_cursor,omitempty"`
}

func (x *NotificationList) Reset() {
	*x = NotificationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[75]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationList) ProtoMessage() {}

func (x *NotificationList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[75]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationList.ProtoReflect.Descriptor instead.
func (*NotificationList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{75}
}

func (x *NotificationList) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

func (x *NotificationList) GetCacheableCursor() string {
	if x != nil {
		return x.CacheableCursor
	}
	return ""
}
