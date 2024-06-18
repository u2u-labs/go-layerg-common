package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Fetch a batch of zero or more users from the server.
type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account id of a user.
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	// The account username of a user.
	Usernames []string `protobuf:"bytes,2,rep,name=usernames,proto3" json:"usernames,omitempty"`
	// The Facebook ID of a user.
	FacebookIds []string `protobuf:"bytes,3,rep,name=facebook_ids,json=facebookIds,proto3" json:"facebook_ids,omitempty"`
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[40]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[40]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{40}
}

func (x *GetUsersRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetUsersRequest) GetUsernames() []string {
	if x != nil {
		return x.Usernames
	}
	return nil
}

func (x *GetUsersRequest) GetFacebookIds() []string {
	if x != nil {
		return x.FacebookIds
	}
	return nil
}

// A user in the server.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the user's account.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The username of the user's account.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// The display name of the user.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// A URL for an avatar image.
	AvatarUrl string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	// The language expected to be a tag which follows the BCP-47 spec.
	LangTag string `protobuf:"bytes,5,opt,name=lang_tag,json=langTag,proto3" json:"lang_tag,omitempty"`
	// The location set by the user.
	Location string `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	// The timezone set by the user.
	Timezone string `protobuf:"bytes,7,opt,name=timezone,proto3" json:"timezone,omitempty"`
	// Additional information stored as a JSON object.
	Metadata string `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The Facebook id in the user's account.
	FacebookId string `protobuf:"bytes,9,opt,name=facebook_id,json=facebookId,proto3" json:"facebook_id,omitempty"`
	// The Google id in the user's account.
	GoogleId string `protobuf:"bytes,10,opt,name=google_id,json=googleId,proto3" json:"google_id,omitempty"`
	// The Apple Game Center in of the user's account.
	GamecenterId string `protobuf:"bytes,11,opt,name=gamecenter_id,json=gamecenterId,proto3" json:"gamecenter_id,omitempty"`
	// The Steam id in the user's account.
	SteamId string `protobuf:"bytes,12,opt,name=steam_id,json=steamId,proto3" json:"steam_id,omitempty"`
	// Indicates whether the user is currently online.
	Online bool `protobuf:"varint,13,opt,name=online,proto3" json:"online,omitempty"`
	// Number of related edges to this user.
	EdgeCount int32 `protobuf:"varint,14,opt,name=edge_count,json=edgeCount,proto3" json:"edge_count,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the user was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the user was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The Facebook Instant Game ID in the user's account.
	FacebookInstantGameId string `protobuf:"bytes,17,opt,name=facebook_instant_game_id,json=facebookInstantGameId,proto3" json:"facebook_instant_game_id,omitempty"`
	// The Apple Sign In ID in the user's account.
	AppleId string `protobuf:"bytes,18,opt,name=apple_id,json=appleId,proto3" json:"apple_id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[92]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[92]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{92}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *User) GetLangTag() string {
	if x != nil {
		return x.LangTag
	}
	return ""
}

func (x *User) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *User) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *User) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *User) GetFacebookId() string {
	if x != nil {
		return x.FacebookId
	}
	return ""
}

func (x *User) GetGoogleId() string {
	if x != nil {
		return x.GoogleId
	}
	return ""
}

func (x *User) GetGamecenterId() string {
	if x != nil {
		return x.GamecenterId
	}
	return ""
}

func (x *User) GetSteamId() string {
	if x != nil {
		return x.SteamId
	}
	return ""
}

func (x *User) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

func (x *User) GetEdgeCount() int32 {
	if x != nil {
		return x.EdgeCount
	}
	return 0
}

func (x *User) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *User) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *User) GetFacebookInstantGameId() string {
	if x != nil {
		return x.FacebookInstantGameId
	}
	return ""
}

func (x *User) GetAppleId() string {
	if x != nil {
		return x.AppleId
	}
	return ""
}

// A list of groups belonging to a user, along with the user's role in each group.
type UserGroupList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Group-role pairs for a user.
	UserGroups []*UserGroupList_UserGroup `protobuf:"bytes,1,rep,name=user_groups,json=userGroups,proto3" json:"user_groups,omitempty"`
	// Cursor for the next page of results, if any.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *UserGroupList) Reset() {
	*x = UserGroupList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[93]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupList) ProtoMessage() {}

func (x *UserGroupList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[93]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupList.ProtoReflect.Descriptor instead.
func (*UserGroupList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{93}
}

func (x *UserGroupList) GetUserGroups() []*UserGroupList_UserGroup {
	if x != nil {
		return x.UserGroups
	}
	return nil
}

func (x *UserGroupList) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// A collection of zero or more users.
type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The User objects.
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[94]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[94]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{94}
}

func (x *Users) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}
