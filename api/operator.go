package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Operator that can be used to override the one set in the leaderboard.
type Operator int32

const (
	// Do not override the leaderboard operator.
	Operator_NO_OVERRIDE Operator = 0
	// Override the leaderboard operator with BEST.
	Operator_BEST Operator = 1
	// Override the leaderboard operator with SET.
	Operator_SET Operator = 2
	// Override the leaderboard operator with INCREMENT.
	Operator_INCREMENT Operator = 3
	// Override the leaderboard operator with DECREMENT.
	Operator_DECREMENT Operator = 4
)

// Enum value maps for Operator.
var (
	Operator_name = map[int32]string{
		0: "NO_OVERRIDE",
		1: "BEST",
		2: "SET",
		3: "INCREMENT",
		4: "DECREMENT",
	}
	Operator_value = map[string]int32{
		"NO_OVERRIDE": 0,
		"BEST":        1,
		"SET":         2,
		"INCREMENT":   3,
		"DECREMENT":   4,
	}
)

func (x Operator) Enum() *Operator {
	p := new(Operator)
	*p = x
	return p
}

func (x Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[2].Descriptor()
}

func (Operator) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[2]
}

func (x Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operator.Descriptor instead.
func (Operator) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}
