package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Send a Google token to the server. Used with authenticate/link/unlink.
type AccountGoogle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The OAuth token received from Google to access their profile API.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Extra information that will be bundled in the session token.
	Vars map[string]string `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AccountGoogle) Reset() {
	*x = AccountGoogle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountGoogle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountGoogle) ProtoMessage() {}

func (x *AccountGoogle) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountGoogle.ProtoReflect.Descriptor instead.
func (*AccountGoogle) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *AccountGoogle) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AccountGoogle) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}
