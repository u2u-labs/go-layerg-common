package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Send a Facebook Instant Game token to the server. Used with authenticate/link/unlink.
type AccountFacebookInstantGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The OAuth token received from a Facebook Instant Game that may be decoded with the Application Secret (must be available with the nakama configuration)
	SignedPlayerInfo string `protobuf:"bytes,1,opt,name=signed_player_info,json=signedPlayerInfo,proto3" json:"signed_player_info,omitempty"`
	// Extra information that will be bundled in the session token.
	Vars map[string]string `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AccountFacebookInstantGame) Reset() {
	*x = AccountFacebookInstantGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountFacebookInstantGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountFacebookInstantGame) ProtoMessage() {}

func (x *AccountFacebookInstantGame) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountFacebookInstantGame.ProtoReflect.Descriptor instead.
func (*AccountFacebookInstantGame) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *AccountFacebookInstantGame) GetSignedPlayerInfo() string {
	if x != nil {
		return x.SignedPlayerInfo
	}
	return ""
}

func (x *AccountFacebookInstantGame) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}
