package rtapi

import (
	"github.com/u2u-labs/go-layerg-common/api"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// A collection of zero or more notifications.
type Notifications struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Collection of notifications.
	Notifications []*api.Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *Notifications) Reset() {
	*x = Notifications{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notifications) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notifications) ProtoMessage() {}

func (x *Notifications) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notifications.ProtoReflect.Descriptor instead.
func (*Notifications) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{21}
}

func (x *Notifications) GetNotifications() []*api.Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}
