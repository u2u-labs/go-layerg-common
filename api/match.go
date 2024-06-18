package api

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// List realtime matches.
type ListMatchesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Limit the number of returned matches.
	Limit *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Authoritative or relayed matches.
	Authoritative *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=authoritative,proto3" json:"authoritative,omitempty"`
	// Label filter.
	Label *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	// Minimum user count.
	MinSize *wrapperspb.Int32Value `protobuf:"bytes,4,opt,name=min_size,json=minSize,proto3" json:"min_size,omitempty"`
	// Maximum user count.
	MaxSize *wrapperspb.Int32Value `protobuf:"bytes,5,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	// Arbitrary label query.
	Query *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *ListMatchesRequest) Reset() {
	*x = ListMatchesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[64]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMatchesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMatchesRequest) ProtoMessage() {}

func (x *ListMatchesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[64]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMatchesRequest.ProtoReflect.Descriptor instead.
func (*ListMatchesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{64}
}

func (x *ListMatchesRequest) GetLimit() *wrapperspb.Int32Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *ListMatchesRequest) GetAuthoritative() *wrapperspb.BoolValue {
	if x != nil {
		return x.Authoritative
	}
	return nil
}

func (x *ListMatchesRequest) GetLabel() *wrapperspb.StringValue {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *ListMatchesRequest) GetMinSize() *wrapperspb.Int32Value {
	if x != nil {
		return x.MinSize
	}
	return nil
}

func (x *ListMatchesRequest) GetMaxSize() *wrapperspb.Int32Value {
	if x != nil {
		return x.MaxSize
	}
	return nil
}

func (x *ListMatchesRequest) GetQuery() *wrapperspb.StringValue {
	if x != nil {
		return x.Query
	}
	return nil
}

// Represents a realtime match.
type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the match, can be used to join.
	MatchId string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	// True if it's an server-managed authoritative match, false otherwise.
	Authoritative bool `protobuf:"varint,2,opt,name=authoritative,proto3" json:"authoritative,omitempty"`
	// Match label, if any.
	Label *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	// Current number of users in the match.
	Size int32 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	// Tick Rate
	TickRate int32 `protobuf:"varint,5,opt,name=tick_rate,json=tickRate,proto3" json:"tick_rate,omitempty"`
	// Handler name
	HandlerName string `protobuf:"bytes,6,opt,name=handler_name,json=handlerName,proto3" json:"handler_name,omitempty"`
}

func (x *Match) Reset() {
	*x = Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[72]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[72]
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
	return file_api_proto_rawDescGZIP(), []int{72}
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

func (x *Match) GetTickRate() int32 {
	if x != nil {
		return x.TickRate
	}
	return 0
}

func (x *Match) GetHandlerName() string {
	if x != nil {
		return x.HandlerName
	}
	return ""
}

// A list of realtime matches.
type MatchList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A number of matches corresponding to a list operation.
	Matches []*Match `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
}

func (x *MatchList) Reset() {
	*x = MatchList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[73]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchList) ProtoMessage() {}

func (x *MatchList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[73]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchList.ProtoReflect.Descriptor instead.
func (*MatchList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{73}
}

func (x *MatchList) GetMatches() []*Match {
	if x != nil {
		return x.Matches
	}
	return nil
}
