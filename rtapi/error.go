package rtapi

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// The selection of possible error codes.
type Error_Code int32

const (
	// An unexpected result from the server.
	Error_RUNTIME_EXCEPTION Error_Code = 0
	// The server received a message which is not recognised.
	Error_UNRECOGNIZED_PAYLOAD Error_Code = 1
	// A message was expected but contains no content.
	Error_MISSING_PAYLOAD Error_Code = 2
	// Fields in the message have an invalid format.
	Error_BAD_INPUT Error_Code = 3
	// The match id was not found.
	Error_MATCH_NOT_FOUND Error_Code = 4
	// The match join was rejected.
	Error_MATCH_JOIN_REJECTED Error_Code = 5
	// The runtime function does not exist on the server.
	Error_RUNTIME_FUNCTION_NOT_FOUND Error_Code = 6
	// The runtime function executed with an error.
	Error_RUNTIME_FUNCTION_EXCEPTION Error_Code = 7
)

// Enum value maps for Error_Code.
var (
	Error_Code_name = map[int32]string{
		0: "RUNTIME_EXCEPTION",
		1: "UNRECOGNIZED_PAYLOAD",
		2: "MISSING_PAYLOAD",
		3: "BAD_INPUT",
		4: "MATCH_NOT_FOUND",
		5: "MATCH_JOIN_REJECTED",
		6: "RUNTIME_FUNCTION_NOT_FOUND",
		7: "RUNTIME_FUNCTION_EXCEPTION",
	}
	Error_Code_value = map[string]int32{
		"RUNTIME_EXCEPTION":          0,
		"UNRECOGNIZED_PAYLOAD":       1,
		"MISSING_PAYLOAD":            2,
		"BAD_INPUT":                  3,
		"MATCH_NOT_FOUND":            4,
		"MATCH_JOIN_REJECTED":        5,
		"RUNTIME_FUNCTION_NOT_FOUND": 6,
		"RUNTIME_FUNCTION_EXCEPTION": 7,
	}
)

func (x Error_Code) Enum() *Error_Code {
	p := new(Error_Code)
	*p = x
	return p
}

func (x Error_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_realtime_proto_enumTypes[1].Descriptor()
}

func (Error_Code) Type() protoreflect.EnumType {
	return &file_realtime_proto_enumTypes[1]
}

func (x Error_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_Code.Descriptor instead.
func (Error_Code) EnumDescriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{9, 0}
}

// A logical error which may occur on the server.
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The error code which should be one of "Error.Code" enums.
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A message in English to help developers debug the response.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Additional error details which may be different for each response.
	Context map[string]string `protobuf:"bytes,3,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{9}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetContext() map[string]string {
	if x != nil {
		return x.Context
	}
	return nil
}
