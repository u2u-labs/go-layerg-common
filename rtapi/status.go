package rtapi

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// A snapshot of statuses for some set of users.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User statuses.
	Presences []*UserPresence `protobuf:"bytes,1,rep,name=presences,proto3" json:"presences,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[41]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[41]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{41}
}

func (x *Status) GetPresences() []*UserPresence {
	if x != nil {
		return x.Presences
	}
	return nil
}

// Start receiving status updates for some set of users.
type StatusFollow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User IDs to follow.
	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	// Usernames to follow.
	Usernames []string `protobuf:"bytes,2,rep,name=usernames,proto3" json:"usernames,omitempty"`
}

func (x *StatusFollow) Reset() {
	*x = StatusFollow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[42]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusFollow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusFollow) ProtoMessage() {}

func (x *StatusFollow) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[42]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusFollow.ProtoReflect.Descriptor instead.
func (*StatusFollow) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{42}
}

func (x *StatusFollow) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *StatusFollow) GetUsernames() []string {
	if x != nil {
		return x.Usernames
	}
	return nil
}

// A batch of status updates for a given user.
type StatusPresenceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// New statuses for the user.
	Joins []*UserPresence `protobuf:"bytes,2,rep,name=joins,proto3" json:"joins,omitempty"`
	// Previous statuses for the user.
	Leaves []*UserPresence `protobuf:"bytes,3,rep,name=leaves,proto3" json:"leaves,omitempty"`
}

func (x *StatusPresenceEvent) Reset() {
	*x = StatusPresenceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[43]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusPresenceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusPresenceEvent) ProtoMessage() {}

func (x *StatusPresenceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[43]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusPresenceEvent.ProtoReflect.Descriptor instead.
func (*StatusPresenceEvent) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{43}
}

func (x *StatusPresenceEvent) GetJoins() []*UserPresence {
	if x != nil {
		return x.Joins
	}
	return nil
}

func (x *StatusPresenceEvent) GetLeaves() []*UserPresence {
	if x != nil {
		return x.Leaves
	}
	return nil
}

// Stop receiving status updates for some set of users.
type StatusUnfollow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Users to unfollow.
	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *StatusUnfollow) Reset() {
	*x = StatusUnfollow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[44]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusUnfollow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusUnfollow) ProtoMessage() {}

func (x *StatusUnfollow) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[44]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusUnfollow.ProtoReflect.Descriptor instead.
func (*StatusUnfollow) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{44}
}

func (x *StatusUnfollow) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// Set the user's own status.
type StatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status string to set, if not present the user will appear offline.
	Status *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusUpdate) Reset() {
	*x = StatusUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[45]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusUpdate) ProtoMessage() {}

func (x *StatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[45]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusUpdate.ProtoReflect.Descriptor instead.
func (*StatusUpdate) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{45}
}

func (x *StatusUpdate) GetStatus() *wrapperspb.StringValue {
	if x != nil {
		return x.Status
	}
	return nil
}
