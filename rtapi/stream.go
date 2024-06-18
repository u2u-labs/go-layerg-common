package rtapi

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// Represents identifying information for a stream.
type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mode identifies the type of stream.
	Mode int32 `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	// Subject is the primary identifier, if any.
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// Subcontext is a secondary identifier, if any.
	Subcontext string `protobuf:"bytes,3,opt,name=subcontext,proto3" json:"subcontext,omitempty"`
	// The label is an arbitrary identifying string, if the stream has one.
	Label string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[46]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[46]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{46}
}

func (x *Stream) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *Stream) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Stream) GetSubcontext() string {
	if x != nil {
		return x.Subcontext
	}
	return ""
}

func (x *Stream) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// A data message delivered over a stream.
type StreamData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stream this data message relates to.
	Stream *Stream `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	// The sender, if any.
	Sender *UserPresence `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	// Arbitrary contents of the data message.
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// True if this data was delivered reliably, false otherwise.
	Reliable bool `protobuf:"varint,4,opt,name=reliable,proto3" json:"reliable,omitempty"`
}

func (x *StreamData) Reset() {
	*x = StreamData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[47]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamData) ProtoMessage() {}

func (x *StreamData) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[47]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamData.ProtoReflect.Descriptor instead.
func (*StreamData) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{47}
}

func (x *StreamData) GetStream() *Stream {
	if x != nil {
		return x.Stream
	}
	return nil
}

func (x *StreamData) GetSender() *UserPresence {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *StreamData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *StreamData) GetReliable() bool {
	if x != nil {
		return x.Reliable
	}
	return false
}

// A set of joins and leaves on a particular stream.
type StreamPresenceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stream this event relates to.
	Stream *Stream `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	// Presences joining the stream as part of this event, if any.
	Joins []*UserPresence `protobuf:"bytes,2,rep,name=joins,proto3" json:"joins,omitempty"`
	// Presences leaving the stream as part of this event, if any.
	Leaves []*UserPresence `protobuf:"bytes,3,rep,name=leaves,proto3" json:"leaves,omitempty"`
}

func (x *StreamPresenceEvent) Reset() {
	*x = StreamPresenceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[48]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamPresenceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamPresenceEvent) ProtoMessage() {}

func (x *StreamPresenceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[48]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamPresenceEvent.ProtoReflect.Descriptor instead.
func (*StreamPresenceEvent) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{48}
}

func (x *StreamPresenceEvent) GetStream() *Stream {
	if x != nil {
		return x.Stream
	}
	return nil
}

func (x *StreamPresenceEvent) GetJoins() []*UserPresence {
	if x != nil {
		return x.Joins
	}
	return nil
}

func (x *StreamPresenceEvent) GetLeaves() []*UserPresence {
	if x != nil {
		return x.Leaves
	}
	return nil
}
