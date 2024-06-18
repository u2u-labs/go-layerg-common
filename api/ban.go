package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Ban users from a group.
type BanGroupUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group to ban users from.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The users to ban.
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *BanGroupUsersRequest) Reset() {
	*x = BanGroupUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BanGroupUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BanGroupUsersRequest) ProtoMessage() {}

func (x *BanGroupUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BanGroupUsersRequest.ProtoReflect.Descriptor instead.
func (*BanGroupUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{24}
}

func (x *BanGroupUsersRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *BanGroupUsersRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}
