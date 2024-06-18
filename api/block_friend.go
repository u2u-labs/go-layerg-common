package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Block one or more friends for the current user.
type BlockFriendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account id of a user.
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	// The account username of a user.
	Usernames []string `protobuf:"bytes,2,rep,name=usernames,proto3" json:"usernames,omitempty"`
}

func (x *BlockFriendsRequest) Reset() {
	*x = BlockFriendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockFriendsRequest) ProtoMessage() {}

func (x *BlockFriendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockFriendsRequest.ProtoReflect.Descriptor instead.
func (*BlockFriendsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{25}
}

func (x *BlockFriendsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *BlockFriendsRequest) GetUsernames() []string {
	if x != nil {
		return x.Usernames
	}
	return nil
}
