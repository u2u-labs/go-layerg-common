package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// A message sent on a channel.
type ChannelMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channel this message belongs to.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The unique ID of this message.
	MessageId string `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// The code representing a message type or category.
	Code *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// Message sender, usually a user ID.
	SenderId string `protobuf:"bytes,4,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	// The username of the message sender, if any.
	Username string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	// The content payload.
	Content string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the message was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the message was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// True if the message was persisted to the channel's history, false otherwise.
	Persistent *wrapperspb.BoolValue `protobuf:"bytes,9,opt,name=persistent,proto3" json:"persistent,omitempty"`
	// The name of the chat room, or an empty string if this message was not sent through a chat room.
	RoomName string `protobuf:"bytes,10,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	// The ID of the group, or an empty string if this message was not sent through a group channel.
	GroupId string `protobuf:"bytes,11,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The ID of the first DM user, or an empty string if this message was not sent through a DM chat.
	UserIdOne string `protobuf:"bytes,12,opt,name=user_id_one,json=userIdOne,proto3" json:"user_id_one,omitempty"`
	// The ID of the second DM user, or an empty string if this message was not sent through a DM chat.
	UserIdTwo string `protobuf:"bytes,13,opt,name=user_id_two,json=userIdTwo,proto3" json:"user_id_two,omitempty"`
}

func (x *ChannelMessage) Reset() {
	*x = ChannelMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[26]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelMessage) ProtoMessage() {}

func (x *ChannelMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[26]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelMessage.ProtoReflect.Descriptor instead.
func (*ChannelMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{26}
}

func (x *ChannelMessage) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ChannelMessage) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *ChannelMessage) GetCode() *wrapperspb.Int32Value {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *ChannelMessage) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *ChannelMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChannelMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChannelMessage) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ChannelMessage) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ChannelMessage) GetPersistent() *wrapperspb.BoolValue {
	if x != nil {
		return x.Persistent
	}
	return nil
}

func (x *ChannelMessage) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *ChannelMessage) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ChannelMessage) GetUserIdOne() string {
	if x != nil {
		return x.UserIdOne
	}
	return ""
}

func (x *ChannelMessage) GetUserIdTwo() string {
	if x != nil {
		return x.UserIdTwo
	}
	return ""
}

// A list of channel messages, usually a result of a list operation.
type ChannelMessageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of messages.
	Messages []*ChannelMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	// The cursor to send when retrieving the next page, if any.
	NextCursor string `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	// The cursor to send when retrieving the previous page, if any.
	PrevCursor string `protobuf:"bytes,3,opt,name=prev_cursor,json=prevCursor,proto3" json:"prev_cursor,omitempty"`
	// Cacheable cursor to list newer messages. Durable and designed to be stored, unlike next/prev cursors.
	CacheableCursor string `protobuf:"bytes,4,opt,name=cacheable_cursor,json=cacheableCursor,proto3" json:"cacheable_cursor,omitempty"`
}

func (x *ChannelMessageList) Reset() {
	*x = ChannelMessageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[27]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelMessageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelMessageList) ProtoMessage() {}

func (x *ChannelMessageList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[27]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelMessageList.ProtoReflect.Descriptor instead.
func (*ChannelMessageList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{27}
}

func (x *ChannelMessageList) GetMessages() []*ChannelMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ChannelMessageList) GetNextCursor() string {
	if x != nil {
		return x.NextCursor
	}
	return ""
}

func (x *ChannelMessageList) GetPrevCursor() string {
	if x != nil {
		return x.PrevCursor
	}
	return ""
}

func (x *ChannelMessageList) GetCacheableCursor() string {
	if x != nil {
		return x.CacheableCursor
	}
	return ""
}

// List a channel's message history.
type ListChannelMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channel ID to list from.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Max number of records to return. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// True if listing should be older messages to newer, false if reverse.
	Forward *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=forward,proto3" json:"forward,omitempty"`
	// A pagination cursor, if any.
	Cursor string `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListChannelMessagesRequest) Reset() {
	*x = ListChannelMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[57]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChannelMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChannelMessagesRequest) ProtoMessage() {}

func (x *ListChannelMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[57]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChannelMessagesRequest.ProtoReflect.Descriptor instead.
func (*ListChannelMessagesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{57}
}

func (x *ListChannelMessagesRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ListChannelMessagesRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListChannelMessagesRequest) GetForward() *wrapperspb.BoolValue {
	if x != nil {
		return x.Forward
	}
	return nil
}

func (x *ListChannelMessagesRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}
