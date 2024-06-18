package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Authenticate against the server with Facebook.
type AuthenticateFacebookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Facebook account details.
	Account *AccountFacebook `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Register the account if the user does not already exist.
	Create *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=create,proto3" json:"create,omitempty"`
	// Set the username on the account at register. Must be unique.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// Import Facebook friends for the user.
	Sync *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=sync,proto3" json:"sync,omitempty"`
}

func (x *AuthenticateFacebookRequest) Reset() {
	*x = AuthenticateFacebookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateFacebookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateFacebookRequest) ProtoMessage() {}

func (x *AuthenticateFacebookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateFacebookRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateFacebookRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{19}
}

func (x *AuthenticateFacebookRequest) GetAccount() *AccountFacebook {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *AuthenticateFacebookRequest) GetCreate() *wrapperspb.BoolValue {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *AuthenticateFacebookRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthenticateFacebookRequest) GetSync() *wrapperspb.BoolValue {
	if x != nil {
		return x.Sync
	}
	return nil
}

// Authenticate against the server with Facebook Instant Game token.
type AuthenticateFacebookInstantGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Facebook Instant Game account details.
	Account *AccountFacebookInstantGame `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Register the account if the user does not already exist.
	Create *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=create,proto3" json:"create,omitempty"`
	// Set the username on the account at register. Must be unique.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AuthenticateFacebookInstantGameRequest) Reset() {
	*x = AuthenticateFacebookInstantGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateFacebookInstantGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateFacebookInstantGameRequest) ProtoMessage() {}

func (x *AuthenticateFacebookInstantGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateFacebookInstantGameRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateFacebookInstantGameRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{20}
}

func (x *AuthenticateFacebookInstantGameRequest) GetAccount() *AccountFacebookInstantGame {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *AuthenticateFacebookInstantGameRequest) GetCreate() *wrapperspb.BoolValue {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *AuthenticateFacebookInstantGameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}
