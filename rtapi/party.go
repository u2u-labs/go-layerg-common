package rtapi

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Incoming information about a party.
type Party struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique party identifier.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// Open flag.
	Open bool `protobuf:"varint,2,opt,name=open,proto3" json:"open,omitempty"`
	// Maximum number of party members.
	MaxSize int32 `protobuf:"varint,3,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	// Self.
	Self *UserPresence `protobuf:"bytes,4,opt,name=self,proto3" json:"self,omitempty"`
	// Leader.
	Leader *UserPresence `protobuf:"bytes,5,opt,name=leader,proto3" json:"leader,omitempty"`
	// All current party members.
	Presences []*UserPresence `protobuf:"bytes,6,rep,name=presences,proto3" json:"presences,omitempty"`
}

func (x *Party) Reset() {
	*x = Party{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Party) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Party) ProtoMessage() {}

func (x *Party) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Party.ProtoReflect.Descriptor instead.
func (*Party) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{22}
}

func (x *Party) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *Party) GetOpen() bool {
	if x != nil {
		return x.Open
	}
	return false
}

func (x *Party) GetMaxSize() int32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *Party) GetSelf() *UserPresence {
	if x != nil {
		return x.Self
	}
	return nil
}

func (x *Party) GetLeader() *UserPresence {
	if x != nil {
		return x.Leader
	}
	return nil
}

func (x *Party) GetPresences() []*UserPresence {
	if x != nil {
		return x.Presences
	}
	return nil
}

// Create a party.
type PartyCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not the party will require join requests to be approved by the party leader.
	Open bool `protobuf:"varint,1,opt,name=open,proto3" json:"open,omitempty"`
	// Maximum number of party members.
	MaxSize int32 `protobuf:"varint,2,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
}

func (x *PartyCreate) Reset() {
	*x = PartyCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyCreate) ProtoMessage() {}

func (x *PartyCreate) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyCreate.ProtoReflect.Descriptor instead.
func (*PartyCreate) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{23}
}

func (x *PartyCreate) GetOpen() bool {
	if x != nil {
		return x.Open
	}
	return false
}

func (x *PartyCreate) GetMaxSize() int32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

// Join a party, or request to join if the party is not open.
type PartyJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID to join.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *PartyJoin) Reset() {
	*x = PartyJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyJoin) ProtoMessage() {}

func (x *PartyJoin) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyJoin.ProtoReflect.Descriptor instead.
func (*PartyJoin) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{24}
}

func (x *PartyJoin) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

// Leave a party.
type PartyLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID to leave.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *PartyLeave) Reset() {
	*x = PartyLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyLeave) ProtoMessage() {}

func (x *PartyLeave) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyLeave.ProtoReflect.Descriptor instead.
func (*PartyLeave) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{25}
}

func (x *PartyLeave) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

// Promote a new party leader.
type PartyPromote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID to promote a new leader for.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// The presence of an existing party member to promote as the new leader.
	Presence *UserPresence `protobuf:"bytes,2,opt,name=presence,proto3" json:"presence,omitempty"`
}

func (x *PartyPromote) Reset() {
	*x = PartyPromote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[26]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyPromote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyPromote) ProtoMessage() {}

func (x *PartyPromote) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[26]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyPromote.ProtoReflect.Descriptor instead.
func (*PartyPromote) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{26}
}

func (x *PartyPromote) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyPromote) GetPresence() *UserPresence {
	if x != nil {
		return x.Presence
	}
	return nil
}

// Announcement of a new party leader.
type PartyLeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID to announce the new leader for.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// The presence of the new party leader.
	Presence *UserPresence `protobuf:"bytes,2,opt,name=presence,proto3" json:"presence,omitempty"`
}

func (x *PartyLeader) Reset() {
	*x = PartyLeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[27]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyLeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyLeader) ProtoMessage() {}

func (x *PartyLeader) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[27]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyLeader.ProtoReflect.Descriptor instead.
func (*PartyLeader) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{27}
}

func (x *PartyLeader) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyLeader) GetPresence() *UserPresence {
	if x != nil {
		return x.Presence
	}
	return nil
}

// Accept a request to join.
type PartyAccept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID to accept a join request for.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// The presence to accept as a party member.
	Presence *UserPresence `protobuf:"bytes,2,opt,name=presence,proto3" json:"presence,omitempty"`
}

func (x *PartyAccept) Reset() {
	*x = PartyAccept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[28]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyAccept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyAccept) ProtoMessage() {}

func (x *PartyAccept) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[28]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyAccept.ProtoReflect.Descriptor instead.
func (*PartyAccept) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{28}
}

func (x *PartyAccept) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyAccept) GetPresence() *UserPresence {
	if x != nil {
		return x.Presence
	}
	return nil
}

// Kick a party member, or decline a request to join.
type PartyRemove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID to remove/reject from.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// The presence to remove or reject.
	Presence *UserPresence `protobuf:"bytes,2,opt,name=presence,proto3" json:"presence,omitempty"`
}

func (x *PartyRemove) Reset() {
	*x = PartyRemove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[29]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyRemove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyRemove) ProtoMessage() {}

func (x *PartyRemove) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[29]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyRemove.ProtoReflect.Descriptor instead.
func (*PartyRemove) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{29}
}

func (x *PartyRemove) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyRemove) GetPresence() *UserPresence {
	if x != nil {
		return x.Presence
	}
	return nil
}

// End a party, kicking all party members and closing it.
type PartyClose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID to close.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *PartyClose) Reset() {
	*x = PartyClose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[30]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyClose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyClose) ProtoMessage() {}

func (x *PartyClose) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[30]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyClose.ProtoReflect.Descriptor instead.
func (*PartyClose) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{30}
}

func (x *PartyClose) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

// Request a list of pending join requests for a party.
type PartyJoinRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID to get a list of join requests for.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *PartyJoinRequestList) Reset() {
	*x = PartyJoinRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[31]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyJoinRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyJoinRequestList) ProtoMessage() {}

func (x *PartyJoinRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[31]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyJoinRequestList.ProtoReflect.Descriptor instead.
func (*PartyJoinRequestList) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{31}
}

func (x *PartyJoinRequestList) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

// Incoming notification for one or more new presences attempting to join the party.
type PartyJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID these presences are attempting to join.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// Presences attempting to join.
	Presences []*UserPresence `protobuf:"bytes,2,rep,name=presences,proto3" json:"presences,omitempty"`
}

func (x *PartyJoinRequest) Reset() {
	*x = PartyJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[32]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyJoinRequest) ProtoMessage() {}

func (x *PartyJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[32]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyJoinRequest.ProtoReflect.Descriptor instead.
func (*PartyJoinRequest) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{32}
}

func (x *PartyJoinRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyJoinRequest) GetPresences() []*UserPresence {
	if x != nil {
		return x.Presences
	}
	return nil
}

// Begin matchmaking as a party.
type PartyMatchmakerAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// Minimum total user count to match together.
	MinCount int32 `protobuf:"varint,2,opt,name=min_count,json=minCount,proto3" json:"min_count,omitempty"`
	// Maximum total user count to match together.
	MaxCount int32 `protobuf:"varint,3,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"`
	// Filter query used to identify suitable users.
	Query string `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	// String properties.
	StringProperties map[string]string `protobuf:"bytes,5,rep,name=string_properties,json=stringProperties,proto3" json:"string_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Numeric properties.
	NumericProperties map[string]float64 `protobuf:"bytes,6,rep,name=numeric_properties,json=numericProperties,proto3" json:"numeric_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	// Optional multiple of the count that must be satisfied.
	CountMultiple *wrapperspb.Int32Value `protobuf:"bytes,7,opt,name=count_multiple,json=countMultiple,proto3" json:"count_multiple,omitempty"`
}

func (x *PartyMatchmakerAdd) Reset() {
	*x = PartyMatchmakerAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[33]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyMatchmakerAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyMatchmakerAdd) ProtoMessage() {}

func (x *PartyMatchmakerAdd) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[33]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyMatchmakerAdd.ProtoReflect.Descriptor instead.
func (*PartyMatchmakerAdd) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{33}
}

func (x *PartyMatchmakerAdd) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyMatchmakerAdd) GetMinCount() int32 {
	if x != nil {
		return x.MinCount
	}
	return 0
}

func (x *PartyMatchmakerAdd) GetMaxCount() int32 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

func (x *PartyMatchmakerAdd) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *PartyMatchmakerAdd) GetStringProperties() map[string]string {
	if x != nil {
		return x.StringProperties
	}
	return nil
}

func (x *PartyMatchmakerAdd) GetNumericProperties() map[string]float64 {
	if x != nil {
		return x.NumericProperties
	}
	return nil
}

func (x *PartyMatchmakerAdd) GetCountMultiple() *wrapperspb.Int32Value {
	if x != nil {
		return x.CountMultiple
	}
	return nil
}

// Cancel a party matchmaking process using a ticket.
type PartyMatchmakerRemove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// The ticket to cancel.
	Ticket string `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *PartyMatchmakerRemove) Reset() {
	*x = PartyMatchmakerRemove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[34]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyMatchmakerRemove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyMatchmakerRemove) ProtoMessage() {}

func (x *PartyMatchmakerRemove) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[34]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyMatchmakerRemove.ProtoReflect.Descriptor instead.
func (*PartyMatchmakerRemove) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{34}
}

func (x *PartyMatchmakerRemove) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyMatchmakerRemove) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

// A response from starting a new party matchmaking process.
type PartyMatchmakerTicket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// The ticket that can be used to cancel matchmaking.
	Ticket string `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *PartyMatchmakerTicket) Reset() {
	*x = PartyMatchmakerTicket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[35]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyMatchmakerTicket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyMatchmakerTicket) ProtoMessage() {}

func (x *PartyMatchmakerTicket) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[35]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyMatchmakerTicket.ProtoReflect.Descriptor instead.
func (*PartyMatchmakerTicket) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{35}
}

func (x *PartyMatchmakerTicket) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyMatchmakerTicket) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

// Incoming party data delivered from the server.
type PartyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The party ID.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// A reference to the user presence that sent this data, if any.
	Presence *UserPresence `protobuf:"bytes,2,opt,name=presence,proto3" json:"presence,omitempty"`
	// Op code value.
	OpCode int64 `protobuf:"varint,3,opt,name=op_code,json=opCode,proto3" json:"op_code,omitempty"`
	// Data payload, if any.
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PartyData) Reset() {
	*x = PartyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[36]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyData) ProtoMessage() {}

func (x *PartyData) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[36]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyData.ProtoReflect.Descriptor instead.
func (*PartyData) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{36}
}

func (x *PartyData) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyData) GetPresence() *UserPresence {
	if x != nil {
		return x.Presence
	}
	return nil
}

func (x *PartyData) GetOpCode() int64 {
	if x != nil {
		return x.OpCode
	}
	return 0
}

func (x *PartyData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Send data to a party.
type PartyDataSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Party ID to send to.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// Op code value.
	OpCode int64 `protobuf:"varint,2,opt,name=op_code,json=opCode,proto3" json:"op_code,omitempty"`
	// Data payload, if any.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PartyDataSend) Reset() {
	*x = PartyDataSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[37]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyDataSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyDataSend) ProtoMessage() {}

func (x *PartyDataSend) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[37]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyDataSend.ProtoReflect.Descriptor instead.
func (*PartyDataSend) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{37}
}

func (x *PartyDataSend) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyDataSend) GetOpCode() int64 {
	if x != nil {
		return x.OpCode
	}
	return 0
}

func (x *PartyDataSend) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Presence update for a particular party.
type PartyPresenceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The party ID.
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// User presences that have just joined the party.
	Joins []*UserPresence `protobuf:"bytes,2,rep,name=joins,proto3" json:"joins,omitempty"`
	// User presences that have just left the party.
	Leaves []*UserPresence `protobuf:"bytes,3,rep,name=leaves,proto3" json:"leaves,omitempty"`
}

func (x *PartyPresenceEvent) Reset() {
	*x = PartyPresenceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[38]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyPresenceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyPresenceEvent) ProtoMessage() {}

func (x *PartyPresenceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[38]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyPresenceEvent.ProtoReflect.Descriptor instead.
func (*PartyPresenceEvent) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{38}
}

func (x *PartyPresenceEvent) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyPresenceEvent) GetJoins() []*UserPresence {
	if x != nil {
		return x.Joins
	}
	return nil
}

func (x *PartyPresenceEvent) GetLeaves() []*UserPresence {
	if x != nil {
		return x.Leaves
	}
	return nil
}
