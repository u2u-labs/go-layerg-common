package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Authenticate against the server with a refresh token.
type SessionRefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Refresh token.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Extra information that will be bundled in the session token.
	Vars map[string]string `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SessionRefreshRequest) Reset() {
	*x = SessionRefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionRefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionRefreshRequest) ProtoMessage() {}

func (x *SessionRefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionRefreshRequest.ProtoReflect.Descriptor instead.
func (*SessionRefreshRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{13}
}

func (x *SessionRefreshRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SessionRefreshRequest) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}

// Log out a session, invalidate a refresh token, or log out all sessions/refresh tokens for a user.
type SessionLogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session token to log out.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Refresh token to invalidate.
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *SessionLogoutRequest) Reset() {
	*x = SessionLogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionLogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionLogoutRequest) ProtoMessage() {}

func (x *SessionLogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionLogoutRequest.ProtoReflect.Descriptor instead.
func (*SessionLogoutRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{14}
}

func (x *SessionLogoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SessionLogoutRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// A user's session used to authenticate messages.
type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// True if the corresponding account was just created, false otherwise.
	Created bool `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	// Authentication credentials.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// Refresh token that can be used for session token renewal.
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[81]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[81]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{81}
}

func (x *Session) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

func (x *Session) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Session) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}
