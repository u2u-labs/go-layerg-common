package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Obtain a new authentication token using a refresh token.
type AccountRefresh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Refresh token.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Extra information that will be bundled in the session token.
	Vars map[string]string `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AccountRefresh) Reset() {
	*x = AccountRefresh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountRefresh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRefresh) ProtoMessage() {}

func (x *AccountRefresh) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRefresh.ProtoReflect.Descriptor instead.
func (*AccountRefresh) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *AccountRefresh) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AccountRefresh) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}
