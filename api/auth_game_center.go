package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Authenticate against the server with Apple's Game Center.
type AuthenticateGameCenterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Game Center account details.
	Account *AccountGameCenter `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Register the account if the user does not already exist.
	Create *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=create,proto3" json:"create,omitempty"`
	// Set the username on the account at register. Must be unique.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AuthenticateGameCenterRequest) Reset() {
	*x = AuthenticateGameCenterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateGameCenterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateGameCenterRequest) ProtoMessage() {}

func (x *AuthenticateGameCenterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateGameCenterRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateGameCenterRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{21}
}

func (x *AuthenticateGameCenterRequest) GetAccount() *AccountGameCenter {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *AuthenticateGameCenterRequest) GetCreate() *wrapperspb.BoolValue {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *AuthenticateGameCenterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}
