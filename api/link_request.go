package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Link Facebook to the current user's account.
type LinkFacebookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Facebook account details.
	Account *AccountFacebook `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Import Facebook friends for the user.
	Sync *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=sync,proto3" json:"sync,omitempty"`
}

func (x *LinkFacebookRequest) Reset() {
	*x = LinkFacebookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[55]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkFacebookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkFacebookRequest) ProtoMessage() {}

func (x *LinkFacebookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[55]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkFacebookRequest.ProtoReflect.Descriptor instead.
func (*LinkFacebookRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{55}
}

func (x *LinkFacebookRequest) GetAccount() *AccountFacebook {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *LinkFacebookRequest) GetSync() *wrapperspb.BoolValue {
	if x != nil {
		return x.Sync
	}
	return nil
}

// Link Steam to the current user's account.
type LinkSteamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Facebook account details.
	Account *AccountSteam `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Import Steam friends for the user.
	Sync *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=sync,proto3" json:"sync,omitempty"`
}

func (x *LinkSteamRequest) Reset() {
	*x = LinkSteamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[56]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkSteamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkSteamRequest) ProtoMessage() {}

func (x *LinkSteamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[56]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkSteamRequest.ProtoReflect.Descriptor instead.
func (*LinkSteamRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{56}
}

func (x *LinkSteamRequest) GetAccount() *AccountSteam {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *LinkSteamRequest) GetSync() *wrapperspb.BoolValue {
	if x != nil {
		return x.Sync
	}
	return nil
}
