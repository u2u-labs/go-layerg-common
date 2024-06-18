package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Delete a leaderboard record.
type DeleteTournamentRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tournament ID to delete from.
	TournamentId string `protobuf:"bytes,1,opt,name=tournament_id,json=tournamentId,proto3" json:"tournament_id,omitempty"`
}

func (x *DeleteTournamentRecordRequest) Reset() {
	*x = DeleteTournamentRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[33]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTournamentRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTournamentRecordRequest) ProtoMessage() {}

func (x *DeleteTournamentRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[33]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTournamentRecordRequest.ProtoReflect.Descriptor instead.
func (*DeleteTournamentRecordRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{33}
}

func (x *DeleteTournamentRecordRequest) GetTournamentId() string {
	if x != nil {
		return x.TournamentId
	}
	return ""
}

// The request to join a tournament.
type JoinTournamentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the tournament to join. The tournament must already exist.
	TournamentId string `protobuf:"bytes,1,opt,name=tournament_id,json=tournamentId,proto3" json:"tournament_id,omitempty"`
}

func (x *JoinTournamentRequest) Reset() {
	*x = JoinTournamentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[48]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinTournamentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinTournamentRequest) ProtoMessage() {}

func (x *JoinTournamentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[48]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinTournamentRequest.ProtoReflect.Descriptor instead.
func (*JoinTournamentRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{48}
}

func (x *JoinTournamentRequest) GetTournamentId() string {
	if x != nil {
		return x.TournamentId
	}
	return ""
}

// List tournament records from a given tournament around the owner.
type ListTournamentRecordsAroundOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the tournament to list for.
	TournamentId string `protobuf:"bytes,1,opt,name=tournament_id,json=tournamentId,proto3" json:"tournament_id,omitempty"`
	// Max number of records to return. Between 1 and 100.
	Limit *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// The owner to retrieve records around.
	OwnerId string `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// Expiry in seconds (since epoch) to begin fetching records from.
	Expiry *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// A next or previous page cursor.
	Cursor string `protobuf:"bytes,5,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListTournamentRecordsAroundOwnerRequest) Reset() {
	*x = ListTournamentRecordsAroundOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[68]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTournamentRecordsAroundOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTournamentRecordsAroundOwnerRequest) ProtoMessage() {}

func (x *ListTournamentRecordsAroundOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[68]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTournamentRecordsAroundOwnerRequest.ProtoReflect.Descriptor instead.
func (*ListTournamentRecordsAroundOwnerRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{68}
}

func (x *ListTournamentRecordsAroundOwnerRequest) GetTournamentId() string {
	if x != nil {
		return x.TournamentId
	}
	return ""
}

func (x *ListTournamentRecordsAroundOwnerRequest) GetLimit() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListTournamentRecordsAroundOwnerRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *ListTournamentRecordsAroundOwnerRequest) GetExpiry() *wrapperspb.Int64Value {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *ListTournamentRecordsAroundOwnerRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// List tournament records from a given tournament.
type ListTournamentRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the tournament to list for.
	TournamentId string `protobuf:"bytes,1,opt,name=tournament_id,json=tournamentId,proto3" json:"tournament_id,omitempty"`
	// One or more owners to retrieve records for.
	OwnerIds []string `protobuf:"bytes,2,rep,name=owner_ids,json=ownerIds,proto3" json:"owner_ids,omitempty"`
	// Max number of records to return. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// A next or previous page cursor.
	Cursor string `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// Expiry in seconds (since epoch) to begin fetching records from.
	Expiry *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *ListTournamentRecordsRequest) Reset() {
	*x = ListTournamentRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[69]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTournamentRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTournamentRecordsRequest) ProtoMessage() {}

func (x *ListTournamentRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[69]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTournamentRecordsRequest.ProtoReflect.Descriptor instead.
func (*ListTournamentRecordsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{69}
}

func (x *ListTournamentRecordsRequest) GetTournamentId() string {
	if x != nil {
		return x.TournamentId
	}
	return ""
}

func (x *ListTournamentRecordsRequest) GetOwnerIds() []string {
	if x != nil {
		return x.OwnerIds
	}
	return nil
}

func (x *ListTournamentRecordsRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListTournamentRecordsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *ListTournamentRecordsRequest) GetExpiry() *wrapperspb.Int64Value {
	if x != nil {
		return x.Expiry
	}
	return nil
}

// List active/upcoming tournaments based on given filters.
type ListTournamentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The start of the categories to include. Defaults to 0.
	CategoryStart *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=category_start,json=categoryStart,proto3" json:"category_start,omitempty"`
	// The end of the categories to include. Defaults to 128.
	CategoryEnd *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=category_end,json=categoryEnd,proto3" json:"category_end,omitempty"`
	// The start time for tournaments. Defaults to epoch.
	StartTime *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The end time for tournaments. Defaults to +1 year from current Unix time.
	EndTime *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Max number of records to return. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,6,opt,name=limit,proto3" json:"limit,omitempty"`
	// A next page cursor for listings (optional).
	Cursor string `protobuf:"bytes,8,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListTournamentsRequest) Reset() {
	*x = ListTournamentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[70]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTournamentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTournamentsRequest) ProtoMessage() {}

func (x *ListTournamentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[70]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTournamentsRequest.ProtoReflect.Descriptor instead.
func (*ListTournamentsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{70}
}

func (x *ListTournamentsRequest) GetCategoryStart() *wrapperspb.UInt32Value {
	if x != nil {
		return x.CategoryStart
	}
	return nil
}

func (x *ListTournamentsRequest) GetCategoryEnd() *wrapperspb.UInt32Value {
	if x != nil {
		return x.CategoryEnd
	}
	return nil
}

func (x *ListTournamentsRequest) GetStartTime() *wrapperspb.UInt32Value {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ListTournamentsRequest) GetEndTime() *wrapperspb.UInt32Value {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *ListTournamentsRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListTournamentsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// A tournament on the server.
type Tournament struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the tournament.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The title for the tournament.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// The description of the tournament. May be blank.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The category of the tournament. e.g. "vip" could be category 1.
	Category uint32 `protobuf:"varint,4,opt,name=category,proto3" json:"category,omitempty"`
	// ASC (0) or DESC (1) sort mode of scores in the tournament.
	SortOrder uint32 `protobuf:"varint,5,opt,name=sort_order,json=sortOrder,proto3" json:"sort_order,omitempty"`
	// The current number of players in the tournament.
	Size uint32 `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	// The maximum number of players for the tournament.
	MaxSize uint32 `protobuf:"varint,7,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	// The maximum score updates allowed per player for the current tournament.
	MaxNumScore uint32 `protobuf:"varint,8,opt,name=max_num_score,json=maxNumScore,proto3" json:"max_num_score,omitempty"`
	// True if the tournament is active and can enter. A computed value.
	CanEnter bool `protobuf:"varint,9,opt,name=can_enter,json=canEnter,proto3" json:"can_enter,omitempty"`
	// The UNIX time when the tournament stops being active until next reset. A computed value.
	EndActive uint32 `protobuf:"varint,10,opt,name=end_active,json=endActive,proto3" json:"end_active,omitempty"`
	// The UNIX time when the tournament is next playable. A computed value.
	NextReset uint32 `protobuf:"varint,11,opt,name=next_reset,json=nextReset,proto3" json:"next_reset,omitempty"`
	// Additional information stored as a JSON object.
	Metadata string `protobuf:"bytes,12,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the tournament was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the tournament will start.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the tournament will be stopped.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Duration of the tournament in seconds.
	Duration uint32 `protobuf:"varint,16,opt,name=duration,proto3" json:"duration,omitempty"`
	// The UNIX time when the tournament start being active. A computed value.
	StartActive uint32 `protobuf:"varint,17,opt,name=start_active,json=startActive,proto3" json:"start_active,omitempty"`
	// The UNIX time when the tournament was last reset. A computed value.
	PrevReset uint32 `protobuf:"varint,18,opt,name=prev_reset,json=prevReset,proto3" json:"prev_reset,omitempty"`
	// Operator.
	Operator Operator `protobuf:"varint,19,opt,name=operator,proto3,enum=nakama.api.Operator" json:"operator,omitempty"`
	// Whether the leaderboard was created authoritatively or not.
	Authoritative bool `protobuf:"varint,20,opt,name=authoritative,proto3" json:"authoritative,omitempty"`
}

func (x *Tournament) Reset() {
	*x = Tournament{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[87]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tournament) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tournament) ProtoMessage() {}

func (x *Tournament) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[87]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tournament.ProtoReflect.Descriptor instead.
func (*Tournament) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{87}
}

func (x *Tournament) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tournament) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Tournament) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Tournament) GetCategory() uint32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *Tournament) GetSortOrder() uint32 {
	if x != nil {
		return x.SortOrder
	}
	return 0
}

func (x *Tournament) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Tournament) GetMaxSize() uint32 {
	if x != nil {
		return x.MaxSize
	}
	return 0
}

func (x *Tournament) GetMaxNumScore() uint32 {
	if x != nil {
		return x.MaxNumScore
	}
	return 0
}

func (x *Tournament) GetCanEnter() bool {
	if x != nil {
		return x.CanEnter
	}
	return false
}

func (x *Tournament) GetEndActive() uint32 {
	if x != nil {
		return x.EndActive
	}
	return 0
}

func (x *Tournament) GetNextReset() uint32 {
	if x != nil {
		return x.NextReset
	}
	return 0
}

func (x *Tournament) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Tournament) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Tournament) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Tournament) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Tournament) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Tournament) GetStartActive() uint32 {
	if x != nil {
		return x.StartActive
	}
	return 0
}

func (x *Tournament) GetPrevReset() uint32 {
	if x != nil {
		return x.PrevReset
	}
	return 0
}

func (x *Tournament) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_NO_OVERRIDE
}

func (x *Tournament) GetAuthoritative() bool {
	if x != nil {
		return x.Authoritative
	}
	return false
}

// A list of tournaments.
type TournamentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of tournaments returned.
	Tournaments []*Tournament `protobuf:"bytes,1,rep,name=tournaments,proto3" json:"tournaments,omitempty"`
	// A pagination cursor (optional).
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *TournamentList) Reset() {
	*x = TournamentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[88]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TournamentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TournamentList) ProtoMessage() {}

func (x *TournamentList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[88]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TournamentList.ProtoReflect.Descriptor instead.
func (*TournamentList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{88}
}

func (x *TournamentList) GetTournaments() []*Tournament {
	if x != nil {
		return x.Tournaments
	}
	return nil
}

func (x *TournamentList) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// A set of tournament records which may be part of a tournament records page or a batch of individual records.
type TournamentRecordList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of tournament records.
	Records []*LeaderboardRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	// A batched set of tournament records belonging to specified owners.
	OwnerRecords []*LeaderboardRecord `protobuf:"bytes,2,rep,name=owner_records,json=ownerRecords,proto3" json:"owner_records,omitempty"`
	// The cursor to send when retireving the next page (optional).
	NextCursor string `protobuf:"bytes,3,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	// The cursor to send when retrieving the previous page (optional).
	PrevCursor string `protobuf:"bytes,4,opt,name=prev_cursor,json=prevCursor,proto3" json:"prev_cursor,omitempty"`
	// The total number of ranks available.
	RankCount int64 `protobuf:"varint,5,opt,name=rank_count,json=rankCount,proto3" json:"rank_count,omitempty"`
}

func (x *TournamentRecordList) Reset() {
	*x = TournamentRecordList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[89]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TournamentRecordList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TournamentRecordList) ProtoMessage() {}

func (x *TournamentRecordList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[89]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TournamentRecordList.ProtoReflect.Descriptor instead.
func (*TournamentRecordList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{89}
}

func (x *TournamentRecordList) GetRecords() []*LeaderboardRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *TournamentRecordList) GetOwnerRecords() []*LeaderboardRecord {
	if x != nil {
		return x.OwnerRecords
	}
	return nil
}

func (x *TournamentRecordList) GetNextCursor() string {
	if x != nil {
		return x.NextCursor
	}
	return ""
}

func (x *TournamentRecordList) GetPrevCursor() string {
	if x != nil {
		return x.PrevCursor
	}
	return ""
}

func (x *TournamentRecordList) GetRankCount() int64 {
	if x != nil {
		return x.RankCount
	}
	return 0
}

// A request to submit a score to a tournament.
type WriteTournamentRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tournament ID to write the record for.
	TournamentId string `protobuf:"bytes,1,opt,name=tournament_id,json=tournamentId,proto3" json:"tournament_id,omitempty"`
	// Record input.
	Record *WriteTournamentRecordRequest_TournamentRecordWrite `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *WriteTournamentRecordRequest) Reset() {
	*x = WriteTournamentRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[110]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteTournamentRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteTournamentRecordRequest) ProtoMessage() {}

func (x *WriteTournamentRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[110]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteTournamentRecordRequest.ProtoReflect.Descriptor instead.
func (*WriteTournamentRecordRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{110}
}

func (x *WriteTournamentRecordRequest) GetTournamentId() string {
	if x != nil {
		return x.TournamentId
	}
	return ""
}

func (x *WriteTournamentRecordRequest) GetRecord() *WriteTournamentRecordRequest_TournamentRecordWrite {
	if x != nil {
		return x.Record
	}
	return nil
}

// Record values to write.
type WriteTournamentRecordRequest_TournamentRecordWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The score value to submit.
	Score int64 `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	// An optional secondary value.
	Subscore int64 `protobuf:"varint,2,opt,name=subscore,proto3" json:"subscore,omitempty"`
	// A JSON object of additional properties (optional).
	Metadata string `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Operator override.
	Operator Operator `protobuf:"varint,4,opt,name=operator,proto3,enum=nakama.api.Operator" json:"operator,omitempty"`
}

func (x *WriteTournamentRecordRequest_TournamentRecordWrite) Reset() {
	*x = WriteTournamentRecordRequest_TournamentRecordWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[127]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteTournamentRecordRequest_TournamentRecordWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteTournamentRecordRequest_TournamentRecordWrite) ProtoMessage() {}

func (x *WriteTournamentRecordRequest_TournamentRecordWrite) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[127]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteTournamentRecordRequest_TournamentRecordWrite.ProtoReflect.Descriptor instead.
func (*WriteTournamentRecordRequest_TournamentRecordWrite) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{110, 0}
}

func (x *WriteTournamentRecordRequest_TournamentRecordWrite) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *WriteTournamentRecordRequest_TournamentRecordWrite) GetSubscore() int64 {
	if x != nil {
		return x.Subscore
	}
	return 0
}

func (x *WriteTournamentRecordRequest_TournamentRecordWrite) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *WriteTournamentRecordRequest_TournamentRecordWrite) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_NO_OVERRIDE
}
