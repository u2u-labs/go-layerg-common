package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Authenticate against the server with a device ID.
type AuthenticateDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The device account details.
	Account *AccountDevice `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Register the account if the user does not already exist.
	Create *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=create,proto3" json:"create,omitempty"`
	// Set the username on the account at register. Must be unique.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AuthenticateDeviceRequest) Reset() {
	*x = AuthenticateDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateDeviceRequest) ProtoMessage() {}

func (x *AuthenticateDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateDeviceRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateDeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{17}
}

func (x *AuthenticateDeviceRequest) GetAccount() *AccountDevice {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *AuthenticateDeviceRequest) GetCreate() *wrapperspb.BoolValue {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *AuthenticateDeviceRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}
