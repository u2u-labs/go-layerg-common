package rtapi

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// A user session associated to a stream, usually through a list operation or a join/leave event.
type UserPresence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user this presence belongs to.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// A unique session ID identifying the particular connection, because the user may have many.
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// The username for display purposes.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// Whether this presence generates persistent data/messages, if applicable for the stream type.
	Persistence bool `protobuf:"varint,4,opt,name=persistence,proto3" json:"persistence,omitempty"`
	// A user-set status message for this stream, if applicable.
	Status *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UserPresence) Reset() {
	*x = UserPresence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[49]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPresence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPresence) ProtoMessage() {}

func (x *UserPresence) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[49]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPresence.ProtoReflect.Descriptor instead.
func (*UserPresence) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{49}
}

func (x *UserPresence) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserPresence) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *UserPresence) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserPresence) GetPersistence() bool {
	if x != nil {
		return x.Persistence
	}
	return false
}

func (x *UserPresence) GetStatus() *wrapperspb.StringValue {
	if x != nil {
		return x.Status
	}
	return nil
}
