package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// Validation Provider,
type StoreProvider int32

const (
	// Apple App Store
	StoreProvider_APPLE_APP_STORE StoreProvider = 0
	// Google Play Store
	StoreProvider_GOOGLE_PLAY_STORE StoreProvider = 1
	// Huawei App Gallery
	StoreProvider_HUAWEI_APP_GALLERY StoreProvider = 2
	// Facebook Instant Store
	StoreProvider_FACEBOOK_INSTANT_STORE StoreProvider = 3
)

// Enum value maps for StoreProvider.
var (
	StoreProvider_name = map[int32]string{
		0: "APPLE_APP_STORE",
		1: "GOOGLE_PLAY_STORE",
		2: "HUAWEI_APP_GALLERY",
		3: "FACEBOOK_INSTANT_STORE",
	}
	StoreProvider_value = map[string]int32{
		"APPLE_APP_STORE":        0,
		"GOOGLE_PLAY_STORE":      1,
		"HUAWEI_APP_GALLERY":     2,
		"FACEBOOK_INSTANT_STORE": 3,
	}
)

func (x StoreProvider) Enum() *StoreProvider {
	p := new(StoreProvider)
	*p = x
	return p
}

func (x StoreProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StoreProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (StoreProvider) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x StoreProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StoreProvider.Descriptor instead.
func (StoreProvider) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}
