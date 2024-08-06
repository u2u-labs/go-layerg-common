package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Authenticate against the server with Metamask.
type AuthenticateMetamaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Metamask account details.
	Account *AccountMetamask `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Register the account if the user does not already exist.
	Create *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=create,proto3" json:"create,omitempty"`
	// Set the username on the account at register. Must be unique.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AuthenticateMetamaskRequest) Reset() {
	*x = AuthenticateMetamaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateMetamaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateMetamaskRequest) ProtoMessage() {}

func (x *AuthenticateMetamaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateMetamaskRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateMetamaskRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{15}
}

func (x *AuthenticateMetamaskRequest) GetAccount() *AccountMetamask {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *AuthenticateMetamaskRequest) GetCreate() *wrapperspb.BoolValue {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *AuthenticateMetamaskRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}
