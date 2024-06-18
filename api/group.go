package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// The group role status.
type GroupUserList_GroupUser_State int32

const (
	// The user is a superadmin with full control of the group.
	GroupUserList_GroupUser_SUPERADMIN GroupUserList_GroupUser_State = 0
	// The user is an admin with additional privileges.
	GroupUserList_GroupUser_ADMIN GroupUserList_GroupUser_State = 1
	// The user is a regular member.
	GroupUserList_GroupUser_MEMBER GroupUserList_GroupUser_State = 2
	// The user has requested to join the group
	GroupUserList_GroupUser_JOIN_REQUEST GroupUserList_GroupUser_State = 3
)

// The group role status.
type UserGroupList_UserGroup_State int32

const (
	// The user is a superadmin with full control of the group.
	UserGroupList_UserGroup_SUPERADMIN UserGroupList_UserGroup_State = 0
	// The user is an admin with additional privileges.
	UserGroupList_UserGroup_ADMIN UserGroupList_UserGroup_State = 1
	// The user is a regular member.
	UserGroupList_UserGroup_MEMBER UserGroupList_UserGroup_State = 2
	// The user has requested to join the group
	UserGroupList_UserGroup_JOIN_REQUEST UserGroupList_UserGroup_State = 3
)

// Enum value maps for GroupUserList_GroupUser_State.
var (
	GroupUserList_GroupUser_State_name = map[int32]string{
		0: "SUPERADMIN",
		1: "ADMIN",
		2: "MEMBER",
		3: "JOIN_REQUEST",
	}
	GroupUserList_GroupUser_State_value = map[string]int32{
		"SUPERADMIN":   0,
		"ADMIN":        1,
		"MEMBER":       2,
		"JOIN_REQUEST": 3,
	}
)

// Enum value maps for UserGroupList_UserGroup_State.
var (
	UserGroupList_UserGroup_State_name = map[int32]string{
		0: "SUPERADMIN",
		1: "ADMIN",
		2: "MEMBER",
		3: "JOIN_REQUEST",
	}
	UserGroupList_UserGroup_State_value = map[string]int32{
		"SUPERADMIN":   0,
		"ADMIN":        1,
		"MEMBER":       2,
		"JOIN_REQUEST": 3,
	}
)

func (x GroupUserList_GroupUser_State) Enum() *GroupUserList_GroupUser_State {
	p := new(GroupUserList_GroupUser_State)
	*p = x
	return p
}

func (x GroupUserList_GroupUser_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupUserList_GroupUser_State) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[4].Descriptor()
}

func (GroupUserList_GroupUser_State) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[4]
}

func (x GroupUserList_GroupUser_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupUserList_GroupUser_State.Descriptor instead.
func (GroupUserList_GroupUser_State) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{44, 0, 0}
}

func (x UserGroupList_UserGroup_State) Enum() *UserGroupList_UserGroup_State {
	p := new(UserGroupList_UserGroup_State)
	*p = x
	return p
}

func (x UserGroupList_UserGroup_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserGroupList_UserGroup_State) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[5].Descriptor()
}

func (UserGroupList_UserGroup_State) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[5]
}

func (x UserGroupList_UserGroup_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserGroupList_UserGroup_State.Descriptor instead.
func (UserGroupList_UserGroup_State) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{93, 0, 0}
}

// Add users to a group.
type AddGroupUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group to add users to.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The users to add.
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *AddGroupUsersRequest) Reset() {
	*x = AddGroupUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGroupUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGroupUsersRequest) ProtoMessage() {}

func (x *AddGroupUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGroupUsersRequest.ProtoReflect.Descriptor instead.
func (*AddGroupUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{12}
}

func (x *AddGroupUsersRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *AddGroupUsersRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// Create a group with the current user as owner.
type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique name for the group.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description for the group.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The language expected to be a tag which follows the BCP-47 spec.
	LangTag string `protobuf:"bytes,3,opt,name=lang_tag,json=langTag,proto3" json:"lang_tag,omitempty"`
	// A URL for an avatar image.
	AvatarUrl string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	// Mark a group as open or not where only admins can accept members.
	Open bool `protobuf:"varint,5,opt,name=open,proto3" json:"open,omitempty"`
	// Maximum number of group members.
	MaxCount int32 `protobuf:"varint,6,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[28]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[28]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{28}
}

func (x *CreateGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateGroupRequest) GetLangTag() string {
	if x != nil {
		return x.LangTag
	}
	return ""
}

func (x *CreateGroupRequest) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *CreateGroupRequest) GetOpen() bool {
	if x != nil {
		return x.Open
	}
	return false
}

func (x *CreateGroupRequest) GetMaxCount() int32 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

// Delete a group the user has access to.
type DeleteGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of a group.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[30]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[30]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{30}
}

func (x *DeleteGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// A group in the server.
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of a group.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The id of the user who created the group.
	CreatorId string `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	// The unique name of the group.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// A description for the group.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// The language expected to be a tag which follows the BCP-47 spec.
	LangTag string `protobuf:"bytes,5,opt,name=lang_tag,json=langTag,proto3" json:"lang_tag,omitempty"`
	// Additional information stored as a JSON object.
	Metadata string `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// A URL for an avatar image.
	AvatarUrl string `protobuf:"bytes,7,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	// Anyone can join open groups, otherwise only admins can accept members.
	Open *wrapperspb.BoolValue `protobuf:"bytes,8,opt,name=open,proto3" json:"open,omitempty"`
	// The current count of all members in the group.
	EdgeCount int32 `protobuf:"varint,9,opt,name=edge_count,json=edgeCount,proto3" json:"edge_count,omitempty"`
	// The maximum number of members allowed.
	MaxCount int32 `protobuf:"varint,10,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the group was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the group was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[42]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[42]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{42}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Group) GetLangTag() string {
	if x != nil {
		return x.LangTag
	}
	return ""
}

func (x *Group) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Group) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Group) GetOpen() *wrapperspb.BoolValue {
	if x != nil {
		return x.Open
	}
	return nil
}

func (x *Group) GetEdgeCount() int32 {
	if x != nil {
		return x.EdgeCount
	}
	return 0
}

func (x *Group) GetMaxCount() int32 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

func (x *Group) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Group) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// One or more groups returned from a listing operation.
type GroupList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One or more groups.
	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	// A cursor used to get the next page.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *GroupList) Reset() {
	*x = GroupList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[43]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupList) ProtoMessage() {}

func (x *GroupList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[43]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupList.ProtoReflect.Descriptor instead.
func (*GroupList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{43}
}

func (x *GroupList) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *GroupList) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// A list of users belonging to a group, along with their role.
type GroupUserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User-role pairs for a group.
	GroupUsers []*GroupUserList_GroupUser `protobuf:"bytes,1,rep,name=group_users,json=groupUsers,proto3" json:"group_users,omitempty"`
	// Cursor for the next page of results, if any.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *GroupUserList) Reset() {
	*x = GroupUserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[44]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupUserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupUserList) ProtoMessage() {}

func (x *GroupUserList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[44]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupUserList.ProtoReflect.Descriptor instead.
func (*GroupUserList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{44}
}

func (x *GroupUserList) GetGroupUsers() []*GroupUserList_GroupUser {
	if x != nil {
		return x.GroupUsers
	}
	return nil
}

func (x *GroupUserList) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// Immediately join an open group, or request to join a closed one.
type JoinGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group ID to join. The group must already exist.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *JoinGroupRequest) Reset() {
	*x = JoinGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[47]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupRequest) ProtoMessage() {}

func (x *JoinGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[47]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupRequest.ProtoReflect.Descriptor instead.
func (*JoinGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{47}
}

func (x *JoinGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// Kick a set of users from a group.
type KickGroupUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group ID to kick from.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The users to kick.
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *KickGroupUsersRequest) Reset() {
	*x = KickGroupUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[49]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickGroupUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickGroupUsersRequest) ProtoMessage() {}

func (x *KickGroupUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[49]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickGroupUsersRequest.ProtoReflect.Descriptor instead.
func (*KickGroupUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{49}
}

func (x *KickGroupUsersRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *KickGroupUsersRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// List groups based on given filters.
type ListGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List groups that contain this value in their names.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional pagination cursor.
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// Max number of groups to return. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// Language tag filter
	LangTag string `protobuf:"bytes,4,opt,name=lang_tag,json=langTag,proto3" json:"lang_tag,omitempty"`
	// Number of group members
	Members *wrapperspb.Int32Value `protobuf:"bytes,5,opt,name=members,proto3" json:"members,omitempty"`
	// Optional Open/Closed filter.
	Open *wrapperspb.BoolValue `protobuf:"bytes,6,opt,name=open,proto3" json:"open,omitempty"`
}

func (x *ListGroupsRequest) Reset() {
	*x = ListGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[60]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupsRequest) ProtoMessage() {}

func (x *ListGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[60]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListGroupsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{60}
}

func (x *ListGroupsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListGroupsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *ListGroupsRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListGroupsRequest) GetLangTag() string {
	if x != nil {
		return x.LangTag
	}
	return ""
}

func (x *ListGroupsRequest) GetMembers() *wrapperspb.Int32Value {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *ListGroupsRequest) GetOpen() *wrapperspb.BoolValue {
	if x != nil {
		return x.Open
	}
	return nil
}

// List all users that are part of a group.
type ListGroupUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group ID to list from.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// Max number of records to return. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// The group user state to list.
	State *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	// An optional next page cursor.
	Cursor string `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListGroupUsersRequest) Reset() {
	*x = ListGroupUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[61]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupUsersRequest) ProtoMessage() {}

func (x *ListGroupUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[61]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupUsersRequest.ProtoReflect.Descriptor instead.
func (*ListGroupUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{61}
}

func (x *ListGroupUsersRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ListGroupUsersRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListGroupUsersRequest) GetState() *wrapperspb.Int32Value {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *ListGroupUsersRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// List the groups a user is part of, and their relationship to each.
type ListUserGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the user.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Max number of records to return. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// The user group state to list.
	State *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	// An optional next page cursor.
	Cursor string `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListUserGroupsRequest) Reset() {
	*x = ListUserGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[71]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserGroupsRequest) ProtoMessage() {}

func (x *ListUserGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[71]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListUserGroupsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{71}
}

func (x *ListUserGroupsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListUserGroupsRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListUserGroupsRequest) GetState() *wrapperspb.Int32Value {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *ListUserGroupsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// Promote a set of users in a group to the next role up.
type PromoteGroupUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group ID to promote in.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The users to promote.
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *PromoteGroupUsersRequest) Reset() {
	*x = PromoteGroupUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[76]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromoteGroupUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromoteGroupUsersRequest) ProtoMessage() {}

func (x *PromoteGroupUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[76]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromoteGroupUsersRequest.ProtoReflect.Descriptor instead.
func (*PromoteGroupUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{76}
}

func (x *PromoteGroupUsersRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *PromoteGroupUsersRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// Demote a set of users in a group to the next role down.
type DemoteGroupUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group ID to demote in.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The users to demote.
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *DemoteGroupUsersRequest) Reset() {
	*x = DemoteGroupUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[77]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoteGroupUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoteGroupUsersRequest) ProtoMessage() {}

func (x *DemoteGroupUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[77]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoteGroupUsersRequest.ProtoReflect.Descriptor instead.
func (*DemoteGroupUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{77}
}

func (x *DemoteGroupUsersRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *DemoteGroupUsersRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// Update fields in a given group.
type UpdateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the group to update.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// Name.
	Name *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description string.
	Description *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Lang tag.
	LangTag *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=lang_tag,json=langTag,proto3" json:"lang_tag,omitempty"`
	// Avatar URL.
	AvatarUrl *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	// Open is true if anyone should be allowed to join, or false if joins must be approved by a group admin.
	Open *wrapperspb.BoolValue `protobuf:"bytes,6,opt,name=open,proto3" json:"open,omitempty"`
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[91]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[91]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{91}
}

func (x *UpdateGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *UpdateGroupRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UpdateGroupRequest) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *UpdateGroupRequest) GetLangTag() *wrapperspb.StringValue {
	if x != nil {
		return x.LangTag
	}
	return nil
}

func (x *UpdateGroupRequest) GetAvatarUrl() *wrapperspb.StringValue {
	if x != nil {
		return x.AvatarUrl
	}
	return nil
}

func (x *UpdateGroupRequest) GetOpen() *wrapperspb.BoolValue {
	if x != nil {
		return x.Open
	}
	return nil
}

// A single user-role pair.
type GroupUserList_GroupUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Their relationship to the group.
	State *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *GroupUserList_GroupUser) Reset() {
	*x = GroupUserList_GroupUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[124]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupUserList_GroupUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupUserList_GroupUser) ProtoMessage() {}

func (x *GroupUserList_GroupUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[124]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupUserList_GroupUser.ProtoReflect.Descriptor instead.
func (*GroupUserList_GroupUser) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{44, 0}
}

func (x *GroupUserList_GroupUser) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GroupUserList_GroupUser) GetState() *wrapperspb.Int32Value {
	if x != nil {
		return x.State
	}
	return nil
}

// A single group-role pair.
type UserGroupList_UserGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Group.
	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// The user's relationship to the group.
	State *wrapperspb.Int32Value `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *UserGroupList_UserGroup) Reset() {
	*x = UserGroupList_UserGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[125]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroupList_UserGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroupList_UserGroup) ProtoMessage() {}

func (x *UserGroupList_UserGroup) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[125]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroupList_UserGroup.ProtoReflect.Descriptor instead.
func (*UserGroupList_UserGroup) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{93, 0}
}

func (x *UserGroupList_UserGroup) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *UserGroupList_UserGroup) GetState() *wrapperspb.Int32Value {
	if x != nil {
		return x.State
	}
	return nil
}
