package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// The friendship status.
type Friend_State int32

const (
	// The user is a friend of the current user.
	Friend_FRIEND Friend_State = 0
	// The current user has sent an invite to the user.
	Friend_INVITE_SENT Friend_State = 1
	// The current user has received an invite from this user.
	Friend_INVITE_RECEIVED Friend_State = 2
	// The current user has blocked this user.
	Friend_BLOCKED Friend_State = 3
)

// Enum value maps for Friend_State.
var (
	Friend_State_name = map[int32]string{
		0: "FRIEND",
		1: "INVITE_SENT",
		2: "INVITE_RECEIVED",
		3: "BLOCKED",
	}
	Friend_State_value = map[string]int32{
		"FRIEND":          0,
		"INVITE_SENT":     1,
		"INVITE_RECEIVED": 2,
		"BLOCKED":         3,
	}
)

func (x Friend_State) Enum() *Friend_State {
	p := new(Friend_State)
	*p = x
	return p
}

func (x Friend_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Friend_State) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[3].Descriptor()
}

func (Friend_State) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[3]
}

func (x Friend_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Friend_State.Descriptor instead.
func (Friend_State) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{37, 0}
}

// Add one or more friends to the current user.
type AddFriendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account id of a user.
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	// The account username of a user.
	Usernames []string `protobuf:"bytes,2,rep,name=usernames,proto3" json:"usernames,omitempty"`
}

func (x *AddFriendsRequest) Reset() {
	*x = AddFriendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendsRequest) ProtoMessage() {}

func (x *AddFriendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendsRequest.ProtoReflect.Descriptor instead.
func (*AddFriendsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{11}
}

func (x *AddFriendsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *AddFriendsRequest) GetUsernames() []string {
	if x != nil {
		return x.Usernames
	}
	return nil
}

// Delete one or more friends for the current user.
type DeleteFriendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account id of a user.
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	// The account username of a user.
	Usernames []string `protobuf:"bytes,2,rep,name=usernames,proto3" json:"usernames,omitempty"`
}

func (x *DeleteFriendsRequest) Reset() {
	*x = DeleteFriendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[29]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFriendsRequest) ProtoMessage() {}

func (x *DeleteFriendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[29]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFriendsRequest.ProtoReflect.Descriptor instead.
func (*DeleteFriendsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{29}
}

func (x *DeleteFriendsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *DeleteFriendsRequest) GetUsernames() []string {
	if x != nil {
		return x.Usernames
	}
	return nil
}

// A friend of a user.
type Friend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user object.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The friend status.
	State *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"` // one of "Friend.State".
	// Time of the latest relationship update.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Friend) Reset() {
	*x = Friend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[37]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Friend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Friend) ProtoMessage() {}

func (x *Friend) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[37]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Friend.ProtoReflect.Descriptor instead.
func (*Friend) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{37}
}

func (x *Friend) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Friend) GetState() *wrapperspb.Int32Value {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Friend) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// A collection of zero or more friends of the user.
type FriendList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Friend objects.
	Friends []*Friend `protobuf:"bytes,1,rep,name=friends,proto3" json:"friends,omitempty"`
	// Cursor for the next page of results, if any.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *FriendList) Reset() {
	*x = FriendList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[38]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendList) ProtoMessage() {}

func (x *FriendList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[38]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendList.ProtoReflect.Descriptor instead.
func (*FriendList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{38}
}

func (x *FriendList) GetFriends() []*Friend {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *FriendList) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// A List of friends of friends
type FriendsOfFriendsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User friends of friends.
	FriendsOfFriends []*FriendsOfFriendsList_FriendOfFriend `protobuf:"bytes,1,rep,name=friends_of_friends,json=friendsOfFriends,proto3" json:"friends_of_friends,omitempty"`
	// Cursor for the next page of results, if any.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *FriendsOfFriendsList) Reset() {
	*x = FriendsOfFriendsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[39]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendsOfFriendsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendsOfFriendsList) ProtoMessage() {}

func (x *FriendsOfFriendsList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[39]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendsOfFriendsList.ProtoReflect.Descriptor instead.
func (*FriendsOfFriendsList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{39}
}

func (x *FriendsOfFriendsList) GetFriendsOfFriends() []*FriendsOfFriendsList_FriendOfFriend {
	if x != nil {
		return x.FriendsOfFriends
	}
	return nil
}

func (x *FriendsOfFriendsList) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// List friends for a user.
type ListFriendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max number of records to return. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// The friend state to list.
	State *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	// An optional next page cursor.
	Cursor string `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListFriendsRequest) Reset() {
	*x = ListFriendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[58]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendsRequest) ProtoMessage() {}

func (x *ListFriendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[58]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendsRequest.ProtoReflect.Descriptor instead.
func (*ListFriendsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{58}
}

func (x *ListFriendsRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListFriendsRequest) GetState() *wrapperspb.Int32Value {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *ListFriendsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type ListFriendsOfFriendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max number of records to return. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// An optional next page cursor.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListFriendsOfFriendsRequest) Reset() {
	*x = ListFriendsOfFriendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[59]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFriendsOfFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFriendsOfFriendsRequest) ProtoMessage() {}

func (x *ListFriendsOfFriendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[59]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFriendsOfFriendsRequest.ProtoReflect.Descriptor instead.
func (*ListFriendsOfFriendsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{59}
}

func (x *ListFriendsOfFriendsRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListFriendsOfFriendsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// A friend of a friend.
type FriendsOfFriendsList_FriendOfFriend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user who referred its friend.
	Referrer string `protobuf:"bytes,1,opt,name=referrer,proto3" json:"referrer,omitempty"`
	// User.
	User *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *FriendsOfFriendsList_FriendOfFriend) Reset() {
	*x = FriendsOfFriendsList_FriendOfFriend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[123]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendsOfFriendsList_FriendOfFriend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendsOfFriendsList_FriendOfFriend) ProtoMessage() {}

func (x *FriendsOfFriendsList_FriendOfFriend) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[123]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendsOfFriendsList_FriendOfFriend.ProtoReflect.Descriptor instead.
func (*FriendsOfFriendsList_FriendOfFriend) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{39, 0}
}

func (x *FriendsOfFriendsList_FriendOfFriend) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

func (x *FriendsOfFriendsList_FriendOfFriend) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}
