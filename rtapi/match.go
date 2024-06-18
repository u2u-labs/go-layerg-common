package rtapi

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// A realtime match.
type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The match unique ID.
	MatchId string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	// True if it's an server-managed authoritative match, false otherwise.
	Authoritative bool `protobuf:"varint,2,opt,name=authoritative,proto3" json:"authoritative,omitempty"`
	// Match label, if any.
	Label *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	// The number of users currently in the match.
	Size int32 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	// The users currently in the match.
	Presences []*UserPresence `protobuf:"bytes,5,rep,name=presences,proto3" json:"presences,omitempty"`
	// A reference to the current user's presence in the match.
	Self *UserPresence `protobuf:"bytes,6,opt,name=self,proto3" json:"self,omitempty"`
}

func (x *Match) Reset() {
	*x = Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{10}
}

func (x *Match) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *Match) GetAuthoritative() bool {
	if x != nil {
		return x.Authoritative
	}
	return false
}

func (x *Match) GetLabel() *wrapperspb.StringValue {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *Match) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Match) GetPresences() []*UserPresence {
	if x != nil {
		return x.Presences
	}
	return nil
}

func (x *Match) GetSelf() *UserPresence {
	if x != nil {
		return x.Self
	}
	return nil
}

// Create a new realtime match.
type MatchCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional name to use when creating the match.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MatchCreate) Reset() {
	*x = MatchCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchCreate) ProtoMessage() {}

func (x *MatchCreate) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchCreate.ProtoReflect.Descriptor instead.
func (*MatchCreate) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{11}
}

func (x *MatchCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Realtime match data received from the server.
type MatchData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The match unique ID.
	MatchId string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	// A reference to the user presence that sent this data, if any.
	Presence *UserPresence `protobuf:"bytes,2,opt,name=presence,proto3" json:"presence,omitempty"`
	// Op code value.
	OpCode int64 `protobuf:"varint,3,opt,name=op_code,json=opCode,proto3" json:"op_code,omitempty"`
	// Data payload, if any.
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// True if this data was delivered reliably, false otherwise.
	Reliable bool `protobuf:"varint,5,opt,name=reliable,proto3" json:"reliable,omitempty"`
}

func (x *MatchData) Reset() {
	*x = MatchData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchData) ProtoMessage() {}

func (x *MatchData) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchData.ProtoReflect.Descriptor instead.
func (*MatchData) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{12}
}

func (x *MatchData) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *MatchData) GetPresence() *UserPresence {
	if x != nil {
		return x.Presence
	}
	return nil
}

func (x *MatchData) GetOpCode() int64 {
	if x != nil {
		return x.OpCode
	}
	return 0
}

func (x *MatchData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MatchData) GetReliable() bool {
	if x != nil {
		return x.Reliable
	}
	return false
}

// Send realtime match data to the server.
type MatchDataSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The match unique ID.
	MatchId string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	// Op code value.
	OpCode int64 `protobuf:"varint,2,opt,name=op_code,json=opCode,proto3" json:"op_code,omitempty"`
	// Data payload, if any.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// List of presences in the match to deliver to, if filtering is required. Otherwise deliver to everyone in the match.
	Presences []*UserPresence `protobuf:"bytes,4,rep,name=presences,proto3" json:"presences,omitempty"`
	// True if the data should be sent reliably, false otherwise.
	Reliable bool `protobuf:"varint,5,opt,name=reliable,proto3" json:"reliable,omitempty"`
}

func (x *MatchDataSend) Reset() {
	*x = MatchDataSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchDataSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchDataSend) ProtoMessage() {}

func (x *MatchDataSend) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchDataSend.ProtoReflect.Descriptor instead.
func (*MatchDataSend) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{13}
}

func (x *MatchDataSend) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *MatchDataSend) GetOpCode() int64 {
	if x != nil {
		return x.OpCode
	}
	return 0
}

func (x *MatchDataSend) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MatchDataSend) GetPresences() []*UserPresence {
	if x != nil {
		return x.Presences
	}
	return nil
}

func (x *MatchDataSend) GetReliable() bool {
	if x != nil {
		return x.Reliable
	}
	return false
}

// Join an existing realtime match.
type MatchJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Id:
	//
	//	*MatchJoin_MatchId
	//	*MatchJoin_Token
	Id isMatchJoin_Id `protobuf_oneof:"id"`
	// An optional set of key-value metadata pairs to be passed to the match handler, if any.
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MatchJoin) Reset() {
	*x = MatchJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchJoin) ProtoMessage() {}

func (x *MatchJoin) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchJoin.ProtoReflect.Descriptor instead.
func (*MatchJoin) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{14}
}

func (m *MatchJoin) GetId() isMatchJoin_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *MatchJoin) GetMatchId() string {
	if x, ok := x.GetId().(*MatchJoin_MatchId); ok {
		return x.MatchId
	}
	return ""
}

func (x *MatchJoin) GetToken() string {
	if x, ok := x.GetId().(*MatchJoin_Token); ok {
		return x.Token
	}
	return ""
}

func (x *MatchJoin) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type isMatchJoin_Id interface {
	isMatchJoin_Id()
}

type MatchJoin_MatchId struct {
	// The match unique ID.
	MatchId string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3,oneof"`
}

type MatchJoin_Token struct {
	// A matchmaking result token.
	Token string `protobuf:"bytes,2,opt,name=token,proto3,oneof"`
}

func (*MatchJoin_MatchId) isMatchJoin_Id() {}

func (*MatchJoin_Token) isMatchJoin_Id() {}

// Leave a realtime match.
type MatchLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The match unique ID.
	MatchId string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
}

func (x *MatchLeave) Reset() {
	*x = MatchLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchLeave) ProtoMessage() {}

func (x *MatchLeave) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchLeave.ProtoReflect.Descriptor instead.
func (*MatchLeave) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{15}
}

func (x *MatchLeave) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

// A set of joins and leaves on a particular realtime match.
type MatchPresenceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The match unique ID.
	MatchId string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	// User presences that have just joined the match.
	Joins []*UserPresence `protobuf:"bytes,2,rep,name=joins,proto3" json:"joins,omitempty"`
	// User presences that have just left the match.
	Leaves []*UserPresence `protobuf:"bytes,3,rep,name=leaves,proto3" json:"leaves,omitempty"`
}

func (x *MatchPresenceEvent) Reset() {
	*x = MatchPresenceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchPresenceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchPresenceEvent) ProtoMessage() {}

func (x *MatchPresenceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchPresenceEvent.ProtoReflect.Descriptor instead.
func (*MatchPresenceEvent) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{16}
}

func (x *MatchPresenceEvent) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *MatchPresenceEvent) GetJoins() []*UserPresence {
	if x != nil {
		return x.Joins
	}
	return nil
}

func (x *MatchPresenceEvent) GetLeaves() []*UserPresence {
	if x != nil {
		return x.Leaves
	}
	return nil
}

// Start a new matchmaking process.
type MatchmakerAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimum total user count to match together.
	MinCount int32 `protobuf:"varint,1,opt,name=min_count,json=minCount,proto3" json:"min_count,omitempty"`
	// Maximum total user count to match together.
	MaxCount int32 `protobuf:"varint,2,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"`
	// Filter query used to identify suitable users.
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// String properties.
	StringProperties map[string]string `protobuf:"bytes,4,rep,name=string_properties,json=stringProperties,proto3" json:"string_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Numeric properties.
	NumericProperties map[string]float64 `protobuf:"bytes,5,rep,name=numeric_properties,json=numericProperties,proto3" json:"numeric_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	// Optional multiple of the count that must be satisfied.
	CountMultiple *wrapperspb.Int32Value `protobuf:"bytes,6,opt,name=count_multiple,json=countMultiple,proto3" json:"count_multiple,omitempty"`
}

func (x *MatchmakerAdd) Reset() {
	*x = MatchmakerAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchmakerAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchmakerAdd) ProtoMessage() {}

func (x *MatchmakerAdd) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchmakerAdd.ProtoReflect.Descriptor instead.
func (*MatchmakerAdd) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{17}
}

func (x *MatchmakerAdd) GetMinCount() int32 {
	if x != nil {
		return x.MinCount
	}
	return 0
}

func (x *MatchmakerAdd) GetMaxCount() int32 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

func (x *MatchmakerAdd) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *MatchmakerAdd) GetStringProperties() map[string]string {
	if x != nil {
		return x.StringProperties
	}
	return nil
}

func (x *MatchmakerAdd) GetNumericProperties() map[string]float64 {
	if x != nil {
		return x.NumericProperties
	}
	return nil
}

func (x *MatchmakerAdd) GetCountMultiple() *wrapperspb.Int32Value {
	if x != nil {
		return x.CountMultiple
	}
	return nil
}

// A successful matchmaking result.
type MatchmakerMatched struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The matchmaking ticket that has completed.
	Ticket string `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	// The match token or match ID to join.
	//
	// Types that are assignable to Id:
	//
	//	*MatchmakerMatched_MatchId
	//	*MatchmakerMatched_Token
	Id isMatchmakerMatched_Id `protobuf_oneof:"id"`
	// The users that have been matched together, and information about their matchmaking data.
	Users []*MatchmakerMatched_MatchmakerUser `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	// A reference to the current user and their properties.
	Self *MatchmakerMatched_MatchmakerUser `protobuf:"bytes,5,opt,name=self,proto3" json:"self,omitempty"`
}

func (x *MatchmakerMatched) Reset() {
	*x = MatchmakerMatched{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchmakerMatched) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchmakerMatched) ProtoMessage() {}

func (x *MatchmakerMatched) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchmakerMatched.ProtoReflect.Descriptor instead.
func (*MatchmakerMatched) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{18}
}

func (x *MatchmakerMatched) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

func (m *MatchmakerMatched) GetId() isMatchmakerMatched_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *MatchmakerMatched) GetMatchId() string {
	if x, ok := x.GetId().(*MatchmakerMatched_MatchId); ok {
		return x.MatchId
	}
	return ""
}

func (x *MatchmakerMatched) GetToken() string {
	if x, ok := x.GetId().(*MatchmakerMatched_Token); ok {
		return x.Token
	}
	return ""
}

func (x *MatchmakerMatched) GetUsers() []*MatchmakerMatched_MatchmakerUser {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *MatchmakerMatched) GetSelf() *MatchmakerMatched_MatchmakerUser {
	if x != nil {
		return x.Self
	}
	return nil
}

type isMatchmakerMatched_Id interface {
	isMatchmakerMatched_Id()
}

type MatchmakerMatched_MatchId struct {
	// Match ID.
	MatchId string `protobuf:"bytes,2,opt,name=match_id,json=matchId,proto3,oneof"`
}

type MatchmakerMatched_Token struct {
	// Match join token.
	Token string `protobuf:"bytes,3,opt,name=token,proto3,oneof"`
}

func (*MatchmakerMatched_MatchId) isMatchmakerMatched_Id() {}

func (*MatchmakerMatched_Token) isMatchmakerMatched_Id() {}

// Cancel an existing ongoing matchmaking process.
type MatchmakerRemove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ticket to cancel.
	Ticket string `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *MatchmakerRemove) Reset() {
	*x = MatchmakerRemove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchmakerRemove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchmakerRemove) ProtoMessage() {}

func (x *MatchmakerRemove) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchmakerRemove.ProtoReflect.Descriptor instead.
func (*MatchmakerRemove) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{19}
}

func (x *MatchmakerRemove) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

// A ticket representing a new matchmaking process.
type MatchmakerTicket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ticket that can be used to cancel matchmaking.
	Ticket string `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *MatchmakerTicket) Reset() {
	*x = MatchmakerTicket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchmakerTicket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchmakerTicket) ProtoMessage() {}

func (x *MatchmakerTicket) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchmakerTicket.ProtoReflect.Descriptor instead.
func (*MatchmakerTicket) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{20}
}

func (x *MatchmakerTicket) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

type MatchmakerMatched_MatchmakerUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User info.
	Presence *UserPresence `protobuf:"bytes,1,opt,name=presence,proto3" json:"presence,omitempty"`
	// Party identifier, if this user was matched as a party member.
	PartyId string `protobuf:"bytes,2,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// String properties.
	StringProperties map[string]string `protobuf:"bytes,5,rep,name=string_properties,json=stringProperties,proto3" json:"string_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Numeric properties.
	NumericProperties map[string]float64 `protobuf:"bytes,6,rep,name=numeric_properties,json=numericProperties,proto3" json:"numeric_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *MatchmakerMatched_MatchmakerUser) Reset() {
	*x = MatchmakerMatched_MatchmakerUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[54]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchmakerMatched_MatchmakerUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchmakerMatched_MatchmakerUser) ProtoMessage() {}

func (x *MatchmakerMatched_MatchmakerUser) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[54]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchmakerMatched_MatchmakerUser.ProtoReflect.Descriptor instead.
func (*MatchmakerMatched_MatchmakerUser) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{18, 0}
}

func (x *MatchmakerMatched_MatchmakerUser) GetPresence() *UserPresence {
	if x != nil {
		return x.Presence
	}
	return nil
}

func (x *MatchmakerMatched_MatchmakerUser) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *MatchmakerMatched_MatchmakerUser) GetStringProperties() map[string]string {
	if x != nil {
		return x.StringProperties
	}
	return nil
}

func (x *MatchmakerMatched_MatchmakerUser) GetNumericProperties() map[string]float64 {
	if x != nil {
		return x.NumericProperties
	}
	return nil
}
