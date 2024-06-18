package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Environment where a purchase/subscription took place,
type StoreEnvironment int32

const (
	// Unknown environment.
	StoreEnvironment_UNKNOWN StoreEnvironment = 0
	// Sandbox/test environment.
	StoreEnvironment_SANDBOX StoreEnvironment = 1
	// Production environment.
	StoreEnvironment_PRODUCTION StoreEnvironment = 2
)

// Enum value maps for StoreEnvironment.
var (
	StoreEnvironment_name = map[int32]string{
		0: "UNKNOWN",
		1: "SANDBOX",
		2: "PRODUCTION",
	}
	StoreEnvironment_value = map[string]int32{
		"UNKNOWN":    0,
		"SANDBOX":    1,
		"PRODUCTION": 2,
	}
)

func (x StoreEnvironment) Enum() *StoreEnvironment {
	p := new(StoreEnvironment)
	*p = x
	return p
}

func (x StoreEnvironment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StoreEnvironment) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[1].Descriptor()
}

func (StoreEnvironment) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[1]
}

func (x StoreEnvironment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StoreEnvironment.Descriptor instead.
func (StoreEnvironment) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}
