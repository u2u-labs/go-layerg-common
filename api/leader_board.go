package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Delete a leaderboard record.
type DeleteLeaderboardRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The leaderboard ID to delete from.
	LeaderboardId string `protobuf:"bytes,1,opt,name=leaderboard_id,json=leaderboardId,proto3" json:"leaderboard_id,omitempty"`
}

func (x *DeleteLeaderboardRecordRequest) Reset() {
	*x = DeleteLeaderboardRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[31]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLeaderboardRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLeaderboardRecordRequest) ProtoMessage() {}

func (x *DeleteLeaderboardRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[31]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLeaderboardRecordRequest.ProtoReflect.Descriptor instead.
func (*DeleteLeaderboardRecordRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{31}
}

func (x *DeleteLeaderboardRecordRequest) GetLeaderboardId() string {
	if x != nil {
		return x.LeaderboardId
	}
	return ""
}

// A leaderboard on the server.
type Leaderboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the leaderboard.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ASC(0) or DESC(1) sort mode of scores in the leaderboard.
	SortOrder uint32 `protobuf:"varint,2,opt,name=sort_order,json=sortOrder,proto3" json:"sort_order,omitempty"`
	// BEST, SET, INCREMENT or DECREMENT operator mode of the leaderboard.
	Operator Operator `protobuf:"varint,3,opt,name=operator,proto3,enum=nakama.api.Operator" json:"operator,omitempty"`
	// The UNIX time when the leaderboard was previously reset. A computed value.
	PrevReset uint32 `protobuf:"varint,4,opt,name=prev_reset,json=prevReset,proto3" json:"prev_reset,omitempty"`
	// The UNIX time when the leaderboard is next playable. A computed value.
	NextReset uint32 `protobuf:"varint,5,opt,name=next_reset,json=nextReset,proto3" json:"next_reset,omitempty"`
	// Additional information stored as a JSON object.
	Metadata string `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the leaderboard was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Whether the leaderboard was created authoritatively or not.
	Authoritative bool `protobuf:"varint,8,opt,name=authoritative,proto3" json:"authoritative,omitempty"`
}

func (x *Leaderboard) Reset() {
	*x = Leaderboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[50]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Leaderboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Leaderboard) ProtoMessage() {}

func (x *Leaderboard) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[50]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Leaderboard.ProtoReflect.Descriptor instead.
func (*Leaderboard) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{50}
}

func (x *Leaderboard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Leaderboard) GetSortOrder() uint32 {
	if x != nil {
		return x.SortOrder
	}
	return 0
}

func (x *Leaderboard) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_NO_OVERRIDE
}

func (x *Leaderboard) GetPrevReset() uint32 {
	if x != nil {
		return x.PrevReset
	}
	return 0
}

func (x *Leaderboard) GetNextReset() uint32 {
	if x != nil {
		return x.NextReset
	}
	return 0
}

func (x *Leaderboard) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Leaderboard) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Leaderboard) GetAuthoritative() bool {
	if x != nil {
		return x.Authoritative
	}
	return false
}

// A list of leaderboards
type LeaderboardList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of leaderboards returned.
	Leaderboards []*Leaderboard `protobuf:"bytes,1,rep,name=leaderboards,proto3" json:"leaderboards,omitempty"`
	// A pagination cursor (optional).
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *LeaderboardList) Reset() {
	*x = LeaderboardList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[51]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderboardList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderboardList) ProtoMessage() {}

func (x *LeaderboardList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[51]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderboardList.ProtoReflect.Descriptor instead.
func (*LeaderboardList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{51}
}

func (x *LeaderboardList) GetLeaderboards() []*Leaderboard {
	if x != nil {
		return x.Leaderboards
	}
	return nil
}

func (x *LeaderboardList) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// Represents a complete leaderboard record with all scores and associated metadata.
type LeaderboardRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the leaderboard this score belongs to.
	LeaderboardId string `protobuf:"bytes,1,opt,name=leaderboard_id,json=leaderboardId,proto3" json:"leaderboard_id,omitempty"`
	// The ID of the score owner, usually a user or group.
	OwnerId string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// The username of the score owner, if the owner is a user.
	Username *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// The score value.
	Score int64 `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	// An optional subscore value.
	Subscore int64 `protobuf:"varint,5,opt,name=subscore,proto3" json:"subscore,omitempty"`
	// The number of submissions to this score record.
	NumScore int32 `protobuf:"varint,6,opt,name=num_score,json=numScore,proto3" json:"num_score,omitempty"`
	// Metadata.
	Metadata string `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the leaderboard record was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the leaderboard record was updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the leaderboard record expires.
	ExpiryTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=expiry_time,json=expiryTime,proto3" json:"expiry_time,omitempty"`
	// The rank of this record.
	Rank int64 `protobuf:"varint,11,opt,name=rank,proto3" json:"rank,omitempty"`
	// The maximum number of score updates allowed by the owner.
	MaxNumScore uint32 `protobuf:"varint,12,opt,name=max_num_score,json=maxNumScore,proto3" json:"max_num_score,omitempty"`
}

func (x *LeaderboardRecord) Reset() {
	*x = LeaderboardRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[52]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderboardRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderboardRecord) ProtoMessage() {}

func (x *LeaderboardRecord) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[52]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderboardRecord.ProtoReflect.Descriptor instead.
func (*LeaderboardRecord) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{52}
}

func (x *LeaderboardRecord) GetLeaderboardId() string {
	if x != nil {
		return x.LeaderboardId
	}
	return ""
}

func (x *LeaderboardRecord) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *LeaderboardRecord) GetUsername() *wrapperspb.StringValue {
	if x != nil {
		return x.Username
	}
	return nil
}

func (x *LeaderboardRecord) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *LeaderboardRecord) GetSubscore() int64 {
	if x != nil {
		return x.Subscore
	}
	return 0
}

func (x *LeaderboardRecord) GetNumScore() int32 {
	if x != nil {
		return x.NumScore
	}
	return 0
}

func (x *LeaderboardRecord) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *LeaderboardRecord) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *LeaderboardRecord) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *LeaderboardRecord) GetExpiryTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiryTime
	}
	return nil
}

func (x *LeaderboardRecord) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *LeaderboardRecord) GetMaxNumScore() uint32 {
	if x != nil {
		return x.MaxNumScore
	}
	return 0
}

// A set of leaderboard records, may be part of a leaderboard records page or a batch of individual records.
type LeaderboardRecordList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of leaderboard records.
	Records []*LeaderboardRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	// A batched set of leaderboard records belonging to specified owners.
	OwnerRecords []*LeaderboardRecord `protobuf:"bytes,2,rep,name=owner_records,json=ownerRecords,proto3" json:"owner_records,omitempty"`
	// The cursor to send when retrieving the next page, if any.
	NextCursor string `protobuf:"bytes,3,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	// The cursor to send when retrieving the previous page, if any.
	PrevCursor string `protobuf:"bytes,4,opt,name=prev_cursor,json=prevCursor,proto3" json:"prev_cursor,omitempty"`
	// The total number of ranks available.
	RankCount int64 `protobuf:"varint,5,opt,name=rank_count,json=rankCount,proto3" json:"rank_count,omitempty"`
}

func (x *LeaderboardRecordList) Reset() {
	*x = LeaderboardRecordList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[53]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderboardRecordList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderboardRecordList) ProtoMessage() {}

func (x *LeaderboardRecordList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[53]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderboardRecordList.ProtoReflect.Descriptor instead.
func (*LeaderboardRecordList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{53}
}

func (x *LeaderboardRecordList) GetRecords() []*LeaderboardRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *LeaderboardRecordList) GetOwnerRecords() []*LeaderboardRecord {
	if x != nil {
		return x.OwnerRecords
	}
	return nil
}

func (x *LeaderboardRecordList) GetNextCursor() string {
	if x != nil {
		return x.NextCursor
	}
	return ""
}

func (x *LeaderboardRecordList) GetPrevCursor() string {
	if x != nil {
		return x.PrevCursor
	}
	return ""
}

func (x *LeaderboardRecordList) GetRankCount() int64 {
	if x != nil {
		return x.RankCount
	}
	return 0
}

// Leave a group.
type LeaveGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group ID to leave.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *LeaveGroupRequest) Reset() {
	*x = LeaveGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[54]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveGroupRequest) ProtoMessage() {}

func (x *LeaveGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[54]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveGroupRequest.ProtoReflect.Descriptor instead.
func (*LeaveGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{54}
}

func (x *LeaveGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// List leaerboard records from a given leaderboard around the owner.
type ListLeaderboardRecordsAroundOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the tournament to list for.
	LeaderboardId string `protobuf:"bytes,1,opt,name=leaderboard_id,json=leaderboardId,proto3" json:"leaderboard_id,omitempty"`
	// Max number of records to return. Between 1 and 100.
	Limit *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// The owner to retrieve records around.
	OwnerId string `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// Expiry in seconds (since epoch) to begin fetching records from.
	Expiry *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	// A next or previous page cursor.
	Cursor string `protobuf:"bytes,5,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ListLeaderboardRecordsAroundOwnerRequest) Reset() {
	*x = ListLeaderboardRecordsAroundOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[62]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLeaderboardRecordsAroundOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLeaderboardRecordsAroundOwnerRequest) ProtoMessage() {}

func (x *ListLeaderboardRecordsAroundOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[62]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLeaderboardRecordsAroundOwnerRequest.ProtoReflect.Descriptor instead.
func (*ListLeaderboardRecordsAroundOwnerRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{62}
}

func (x *ListLeaderboardRecordsAroundOwnerRequest) GetLeaderboardId() string {
	if x != nil {
		return x.LeaderboardId
	}
	return ""
}

func (x *ListLeaderboardRecordsAroundOwnerRequest) GetLimit() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListLeaderboardRecordsAroundOwnerRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *ListLeaderboardRecordsAroundOwnerRequest) GetExpiry() *wrapperspb.Int64Value {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *ListLeaderboardRecordsAroundOwnerRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

// List leaderboard records from a given leaderboard.
type ListLeaderboardRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the leaderboard to list for.
	LeaderboardId string `protobuf:"bytes,1,opt,name=leaderboard_id,json=leaderboardId,proto3" json:"leaderboard_id,omitempty"`
	// One or more owners to retrieve records for.
	OwnerIds []string `protobuf:"bytes,2,rep,name=owner_ids,json=ownerIds,proto3" json:"owner_ids,omitempty"`
	// Max number of records to return. Between 1 and 100.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// A next or previous page cursor.
	Cursor string `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// Expiry in seconds (since epoch) to begin fetching records from. Optional. 0 means from current time.
	Expiry *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *ListLeaderboardRecordsRequest) Reset() {
	*x = ListLeaderboardRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[63]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLeaderboardRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLeaderboardRecordsRequest) ProtoMessage() {}

func (x *ListLeaderboardRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[63]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLeaderboardRecordsRequest.ProtoReflect.Descriptor instead.
func (*ListLeaderboardRecordsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{63}
}

func (x *ListLeaderboardRecordsRequest) GetLeaderboardId() string {
	if x != nil {
		return x.LeaderboardId
	}
	return ""
}

func (x *ListLeaderboardRecordsRequest) GetOwnerIds() []string {
	if x != nil {
		return x.OwnerIds
	}
	return nil
}

func (x *ListLeaderboardRecordsRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListLeaderboardRecordsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *ListLeaderboardRecordsRequest) GetExpiry() *wrapperspb.Int64Value {
	if x != nil {
		return x.Expiry
	}
	return nil
}

// A request to submit a score to a leaderboard.
type WriteLeaderboardRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the leaderboard to write to.
	LeaderboardId string `protobuf:"bytes,1,opt,name=leaderboard_id,json=leaderboardId,proto3" json:"leaderboard_id,omitempty"`
	// Record input.
	Record *WriteLeaderboardRecordRequest_LeaderboardRecordWrite `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *WriteLeaderboardRecordRequest) Reset() {
	*x = WriteLeaderboardRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[107]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteLeaderboardRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteLeaderboardRecordRequest) ProtoMessage() {}

func (x *WriteLeaderboardRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[107]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteLeaderboardRecordRequest.ProtoReflect.Descriptor instead.
func (*WriteLeaderboardRecordRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{107}
}

func (x *WriteLeaderboardRecordRequest) GetLeaderboardId() string {
	if x != nil {
		return x.LeaderboardId
	}
	return ""
}

func (x *WriteLeaderboardRecordRequest) GetRecord() *WriteLeaderboardRecordRequest_LeaderboardRecordWrite {
	if x != nil {
		return x.Record
	}
	return nil
}

// Record values to write.
type WriteLeaderboardRecordRequest_LeaderboardRecordWrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The score value to submit.
	Score int64 `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	// An optional secondary value.
	Subscore int64 `protobuf:"varint,2,opt,name=subscore,proto3" json:"subscore,omitempty"`
	// Optional record metadata.
	Metadata string `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Operator override.
	Operator Operator `protobuf:"varint,4,opt,name=operator,proto3,enum=nakama.api.Operator" json:"operator,omitempty"`
}

func (x *WriteLeaderboardRecordRequest_LeaderboardRecordWrite) Reset() {
	*x = WriteLeaderboardRecordRequest_LeaderboardRecordWrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[126]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteLeaderboardRecordRequest_LeaderboardRecordWrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteLeaderboardRecordRequest_LeaderboardRecordWrite) ProtoMessage() {}

func (x *WriteLeaderboardRecordRequest_LeaderboardRecordWrite) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[126]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteLeaderboardRecordRequest_LeaderboardRecordWrite.ProtoReflect.Descriptor instead.
func (*WriteLeaderboardRecordRequest_LeaderboardRecordWrite) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{107, 0}
}

func (x *WriteLeaderboardRecordRequest_LeaderboardRecordWrite) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *WriteLeaderboardRecordRequest_LeaderboardRecordWrite) GetSubscore() int64 {
	if x != nil {
		return x.Subscore
	}
	return 0
}

func (x *WriteLeaderboardRecordRequest_LeaderboardRecordWrite) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *WriteLeaderboardRecordRequest_LeaderboardRecordWrite) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_NO_OVERRIDE
}
