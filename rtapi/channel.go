package rtapi

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// The type of chat channel.
type ChannelJoin_Type int32

const (
	// Default case. Assumed as ROOM type.
	ChannelJoin_TYPE_UNSPECIFIED ChannelJoin_Type = 0
	// A room which anyone can join to chat.
	ChannelJoin_ROOM ChannelJoin_Type = 1
	// A private channel for 1-on-1 chat.
	ChannelJoin_DIRECT_MESSAGE ChannelJoin_Type = 2
	// A channel for group chat.
	ChannelJoin_GROUP ChannelJoin_Type = 3
)

// Enum value maps for ChannelJoin_Type.
var (
	ChannelJoin_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "ROOM",
		2: "DIRECT_MESSAGE",
		3: "GROUP",
	}
	ChannelJoin_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"ROOM":             1,
		"DIRECT_MESSAGE":   2,
		"GROUP":            3,
	}
)

func (x ChannelJoin_Type) Enum() *ChannelJoin_Type {
	p := new(ChannelJoin_Type)
	*p = x
	return p
}

func (x ChannelJoin_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChannelJoin_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_realtime_proto_enumTypes[0].Descriptor()
}

func (ChannelJoin_Type) Type() protoreflect.EnumType {
	return &file_realtime_proto_enumTypes[0]
}

func (x ChannelJoin_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChannelJoin_Type.Descriptor instead.
func (ChannelJoin_Type) EnumDescriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{2, 0}
}

// A realtime chat channel.
type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the channel.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The users currently in the channel.
	Presences []*UserPresence `protobuf:"bytes,2,rep,name=presences,proto3" json:"presences,omitempty"`
	// A reference to the current user's presence in the channel.
	Self *UserPresence `protobuf:"bytes,3,opt,name=self,proto3" json:"self,omitempty"`
	// The name of the chat room, or an empty string if this message was not sent through a chat room.
	RoomName string `protobuf:"bytes,4,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	// The ID of the group, or an empty string if this message was not sent through a group channel.
	GroupId string `protobuf:"bytes,5,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The ID of the first DM user, or an empty string if this message was not sent through a DM chat.
	UserIdOne string `protobuf:"bytes,6,opt,name=user_id_one,json=userIdOne,proto3" json:"user_id_one,omitempty"`
	// The ID of the second DM user, or an empty string if this message was not sent through a DM chat.
	UserIdTwo string `protobuf:"bytes,7,opt,name=user_id_two,json=userIdTwo,proto3" json:"user_id_two,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{1}
}

func (x *Channel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Channel) GetPresences() []*UserPresence {
	if x != nil {
		return x.Presences
	}
	return nil
}

func (x *Channel) GetSelf() *UserPresence {
	if x != nil {
		return x.Self
	}
	return nil
}

func (x *Channel) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *Channel) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Channel) GetUserIdOne() string {
	if x != nil {
		return x.UserIdOne
	}
	return ""
}

func (x *Channel) GetUserIdTwo() string {
	if x != nil {
		return x.UserIdTwo
	}
	return ""
}

// Join operation for a realtime chat channel.
type ChannelJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user ID to DM with, group ID to chat with, or room channel name to join.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The type of the chat channel.
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"` // one of "ChannelId.Type".
	// Whether messages sent on this channel should be persistent.
	Persistence *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=persistence,proto3" json:"persistence,omitempty"`
	// Whether the user should appear in the channel's presence list and events.
	Hidden *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=hidden,proto3" json:"hidden,omitempty"`
}

func (x *ChannelJoin) Reset() {
	*x = ChannelJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelJoin) ProtoMessage() {}

func (x *ChannelJoin) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelJoin.ProtoReflect.Descriptor instead.
func (*ChannelJoin) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelJoin) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *ChannelJoin) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ChannelJoin) GetPersistence() *wrapperspb.BoolValue {
	if x != nil {
		return x.Persistence
	}
	return nil
}

func (x *ChannelJoin) GetHidden() *wrapperspb.BoolValue {
	if x != nil {
		return x.Hidden
	}
	return nil
}

// Leave a realtime channel.
type ChannelLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the channel to leave.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *ChannelLeave) Reset() {
	*x = ChannelLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelLeave) ProtoMessage() {}

func (x *ChannelLeave) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelLeave.ProtoReflect.Descriptor instead.
func (*ChannelLeave) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{3}
}

func (x *ChannelLeave) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

// A receipt reply from a channel message send operation.
type ChannelMessageAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channel the message was sent to.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The unique ID assigned to the message.
	MessageId string `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// The code representing a message type or category.
	Code *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// Username of the message sender.
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the message was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the message was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// True if the message was persisted to the channel's history, false otherwise.
	Persistent *wrapperspb.BoolValue `protobuf:"bytes,7,opt,name=persistent,proto3" json:"persistent,omitempty"`
	// The name of the chat room, or an empty string if this message was not sent through a chat room.
	RoomName string `protobuf:"bytes,8,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	// The ID of the group, or an empty string if this message was not sent through a group channel.
	GroupId string `protobuf:"bytes,9,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The ID of the first DM user, or an empty string if this message was not sent through a DM chat.
	UserIdOne string `protobuf:"bytes,10,opt,name=user_id_one,json=userIdOne,proto3" json:"user_id_one,omitempty"`
	// The ID of the second DM user, or an empty string if this message was not sent through a DM chat.
	UserIdTwo string `protobuf:"bytes,11,opt,name=user_id_two,json=userIdTwo,proto3" json:"user_id_two,omitempty"`
}

func (x *ChannelMessageAck) Reset() {
	*x = ChannelMessageAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelMessageAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelMessageAck) ProtoMessage() {}

func (x *ChannelMessageAck) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelMessageAck.ProtoReflect.Descriptor instead.
func (*ChannelMessageAck) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{4}
}

func (x *ChannelMessageAck) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ChannelMessageAck) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *ChannelMessageAck) GetCode() *wrapperspb.Int32Value {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *ChannelMessageAck) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChannelMessageAck) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ChannelMessageAck) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ChannelMessageAck) GetPersistent() *wrapperspb.BoolValue {
	if x != nil {
		return x.Persistent
	}
	return nil
}

func (x *ChannelMessageAck) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *ChannelMessageAck) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ChannelMessageAck) GetUserIdOne() string {
	if x != nil {
		return x.UserIdOne
	}
	return ""
}

func (x *ChannelMessageAck) GetUserIdTwo() string {
	if x != nil {
		return x.UserIdTwo
	}
	return ""
}

// Send a message to a realtime channel.
type ChannelMessageSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channel to sent to.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Message content.
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChannelMessageSend) Reset() {
	*x = ChannelMessageSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelMessageSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelMessageSend) ProtoMessage() {}

func (x *ChannelMessageSend) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelMessageSend.ProtoReflect.Descriptor instead.
func (*ChannelMessageSend) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{5}
}

func (x *ChannelMessageSend) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ChannelMessageSend) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// Update a message previously sent to a realtime channel.
type ChannelMessageUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channel the message was sent to.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The ID assigned to the message to update.
	MessageId string `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// New message content.
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChannelMessageUpdate) Reset() {
	*x = ChannelMessageUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelMessageUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelMessageUpdate) ProtoMessage() {}

func (x *ChannelMessageUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelMessageUpdate.ProtoReflect.Descriptor instead.
func (*ChannelMessageUpdate) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{6}
}

func (x *ChannelMessageUpdate) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ChannelMessageUpdate) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *ChannelMessageUpdate) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// Remove a message previously sent to a realtime channel.
type ChannelMessageRemove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channel the message was sent to.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The ID assigned to the message to update.
	MessageId string `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *ChannelMessageRemove) Reset() {
	*x = ChannelMessageRemove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelMessageRemove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelMessageRemove) ProtoMessage() {}

func (x *ChannelMessageRemove) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelMessageRemove.ProtoReflect.Descriptor instead.
func (*ChannelMessageRemove) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{7}
}

func (x *ChannelMessageRemove) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ChannelMessageRemove) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

// A set of joins and leaves on a particular channel.
type ChannelPresenceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channel identifier this event is for.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Presences joining the channel as part of this event, if any.
	Joins []*UserPresence `protobuf:"bytes,2,rep,name=joins,proto3" json:"joins,omitempty"`
	// Presences leaving the channel as part of this event, if any.
	Leaves []*UserPresence `protobuf:"bytes,3,rep,name=leaves,proto3" json:"leaves,omitempty"`
	// The name of the chat room, or an empty string if this message was not sent through a chat room.
	RoomName string `protobuf:"bytes,4,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	// The ID of the group, or an empty string if this message was not sent through a group channel.
	GroupId string `protobuf:"bytes,5,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The ID of the first DM user, or an empty string if this message was not sent through a DM chat.
	UserIdOne string `protobuf:"bytes,6,opt,name=user_id_one,json=userIdOne,proto3" json:"user_id_one,omitempty"`
	// The ID of the second DM user, or an empty string if this message was not sent through a DM chat.
	UserIdTwo string `protobuf:"bytes,7,opt,name=user_id_two,json=userIdTwo,proto3" json:"user_id_two,omitempty"`
}

func (x *ChannelPresenceEvent) Reset() {
	*x = ChannelPresenceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelPresenceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelPresenceEvent) ProtoMessage() {}

func (x *ChannelPresenceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelPresenceEvent.ProtoReflect.Descriptor instead.
func (*ChannelPresenceEvent) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{8}
}

func (x *ChannelPresenceEvent) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ChannelPresenceEvent) GetJoins() []*UserPresence {
	if x != nil {
		return x.Joins
	}
	return nil
}

func (x *ChannelPresenceEvent) GetLeaves() []*UserPresence {
	if x != nil {
		return x.Leaves
	}
	return nil
}

func (x *ChannelPresenceEvent) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *ChannelPresenceEvent) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ChannelPresenceEvent) GetUserIdOne() string {
	if x != nil {
		return x.UserIdOne
	}
	return ""
}

func (x *ChannelPresenceEvent) GetUserIdTwo() string {
	if x != nil {
		return x.UserIdTwo
	}
	return ""
}
