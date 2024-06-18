package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Authenticate against the server with Steam.
type AuthenticateSteamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Steam account details.
	Account *AccountSteam `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Register the account if the user does not already exist.
	Create *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=create,proto3" json:"create,omitempty"`
	// Set the username on the account at register. Must be unique.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// Import Steam friends for the user.
	Sync *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=sync,proto3" json:"sync,omitempty"`
}

func (x *AuthenticateSteamRequest) Reset() {
	*x = AuthenticateSteamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateSteamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateSteamRequest) ProtoMessage() {}

func (x *AuthenticateSteamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateSteamRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateSteamRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{23}
}

func (x *AuthenticateSteamRequest) GetAccount() *AccountSteam {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *AuthenticateSteamRequest) GetCreate() *wrapperspb.BoolValue {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *AuthenticateSteamRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthenticateSteamRequest) GetSync() *wrapperspb.BoolValue {
	if x != nil {
		return x.Sync
	}
	return nil
}
