package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Execute an Lua function on the server.
type Rpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the function.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The payload of the function which must be a JSON object.
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// The authentication key used when executed as a non-client HTTP request.
	HttpKey string `protobuf:"bytes,3,opt,name=http_key,json=httpKey,proto3" json:"http_key,omitempty"`
}

func (x *Rpc) Reset() {
	*x = Rpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[80]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rpc) ProtoMessage() {}

func (x *Rpc) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[80]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rpc.ProtoReflect.Descriptor instead.
func (*Rpc) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{80}
}

func (x *Rpc) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Rpc) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *Rpc) GetHttpKey() string {
	if x != nil {
		return x.HttpKey
	}
	return ""
}
