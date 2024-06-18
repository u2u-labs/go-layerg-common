package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Import Facebook friends into the current user's account.
type ImportFacebookFriendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Facebook account details.
	Account *AccountFacebook `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Reset the current user's friends list.
	Reset_ *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=reset,proto3" json:"reset,omitempty"`
}

func (x *ImportFacebookFriendsRequest) Reset() {
	*x = ImportFacebookFriendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[45]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportFacebookFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportFacebookFriendsRequest) ProtoMessage() {}

func (x *ImportFacebookFriendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[45]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportFacebookFriendsRequest.ProtoReflect.Descriptor instead.
func (*ImportFacebookFriendsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{45}
}

func (x *ImportFacebookFriendsRequest) GetAccount() *AccountFacebook {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *ImportFacebookFriendsRequest) GetReset_() *wrapperspb.BoolValue {
	if x != nil {
		return x.Reset_
	}
	return nil
}

// Import Facebook friends into the current user's account.
type ImportSteamFriendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Facebook account details.
	Account *AccountSteam `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Reset the current user's friends list.
	Reset_ *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=reset,proto3" json:"reset,omitempty"`
}

func (x *ImportSteamFriendsRequest) Reset() {
	*x = ImportSteamFriendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[46]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportSteamFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSteamFriendsRequest) ProtoMessage() {}

func (x *ImportSteamFriendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[46]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportSteamFriendsRequest.ProtoReflect.Descriptor instead.
func (*ImportSteamFriendsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{46}
}

func (x *ImportSteamFriendsRequest) GetAccount() *AccountSteam {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *ImportSteamFriendsRequest) GetReset_() *wrapperspb.BoolValue {
	if x != nil {
		return x.Reset_
	}
	return nil
}
