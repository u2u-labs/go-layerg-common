package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Send an email with password to the server. Used with authenticate/link/unlink.
type AccountEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A valid RFC-5322 email address.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// A password for the user account.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // Ignored with unlink operations.
	// Extra information that will be bundled in the session token.
	Vars map[string]string `protobuf:"bytes,3,rep,name=vars,proto3" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AccountEmail) Reset() {
	*x = AccountEmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountEmail) ProtoMessage() {}

func (x *AccountEmail) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountEmail.ProtoReflect.Descriptor instead.
func (*AccountEmail) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *AccountEmail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AccountEmail) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AccountEmail) GetVars() map[string]string {
	if x != nil {
		return x.Vars
	}
	return nil
}
