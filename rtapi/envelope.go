package rtapi

import (
	"github.com/u2u-labs/go-layerg-common/api"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// An envelope for a realtime message.
type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	// Types that are assignable to Message:
	//
	//	*Envelope_Channel
	//	*Envelope_ChannelJoin
	//	*Envelope_ChannelLeave
	//	*Envelope_ChannelMessage
	//	*Envelope_ChannelMessageAck
	//	*Envelope_ChannelMessageSend
	//	*Envelope_ChannelMessageUpdate
	//	*Envelope_ChannelMessageRemove
	//	*Envelope_ChannelPresenceEvent
	//	*Envelope_Error
	//	*Envelope_Match
	//	*Envelope_MatchCreate
	//	*Envelope_MatchData
	//	*Envelope_MatchDataSend
	//	*Envelope_MatchJoin
	//	*Envelope_MatchLeave
	//	*Envelope_MatchPresenceEvent
	//	*Envelope_MatchmakerAdd
	//	*Envelope_MatchmakerMatched
	//	*Envelope_MatchmakerRemove
	//	*Envelope_MatchmakerTicket
	//	*Envelope_Notifications
	//	*Envelope_Rpc
	//	*Envelope_Status
	//	*Envelope_StatusFollow
	//	*Envelope_StatusPresenceEvent
	//	*Envelope_StatusUnfollow
	//	*Envelope_StatusUpdate
	//	*Envelope_StreamData
	//	*Envelope_StreamPresenceEvent
	//	*Envelope_Ping
	//	*Envelope_Pong
	//	*Envelope_Party
	//	*Envelope_PartyCreate
	//	*Envelope_PartyJoin
	//	*Envelope_PartyLeave
	//	*Envelope_PartyPromote
	//	*Envelope_PartyLeader
	//	*Envelope_PartyAccept
	//	*Envelope_PartyRemove
	//	*Envelope_PartyClose
	//	*Envelope_PartyJoinRequestList
	//	*Envelope_PartyJoinRequest
	//	*Envelope_PartyMatchmakerAdd
	//	*Envelope_PartyMatchmakerRemove
	//	*Envelope_PartyMatchmakerTicket
	//	*Envelope_PartyData
	//	*Envelope_PartyDataSend
	//	*Envelope_PartyPresenceEvent
	Message isEnvelope_Message `protobuf_oneof:"message"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realtime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_realtime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_realtime_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (m *Envelope) GetMessage() isEnvelope_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *Envelope) GetChannel() *Channel {
	if x, ok := x.GetMessage().(*Envelope_Channel); ok {
		return x.Channel
	}
	return nil
}

func (x *Envelope) GetChannelJoin() *ChannelJoin {
	if x, ok := x.GetMessage().(*Envelope_ChannelJoin); ok {
		return x.ChannelJoin
	}
	return nil
}

func (x *Envelope) GetChannelLeave() *ChannelLeave {
	if x, ok := x.GetMessage().(*Envelope_ChannelLeave); ok {
		return x.ChannelLeave
	}
	return nil
}

func (x *Envelope) GetChannelMessage() *api.ChannelMessage {
	if x, ok := x.GetMessage().(*Envelope_ChannelMessage); ok {
		return x.ChannelMessage
	}
	return nil
}

func (x *Envelope) GetChannelMessageAck() *ChannelMessageAck {
	if x, ok := x.GetMessage().(*Envelope_ChannelMessageAck); ok {
		return x.ChannelMessageAck
	}
	return nil
}

func (x *Envelope) GetChannelMessageSend() *ChannelMessageSend {
	if x, ok := x.GetMessage().(*Envelope_ChannelMessageSend); ok {
		return x.ChannelMessageSend
	}
	return nil
}

func (x *Envelope) GetChannelMessageUpdate() *ChannelMessageUpdate {
	if x, ok := x.GetMessage().(*Envelope_ChannelMessageUpdate); ok {
		return x.ChannelMessageUpdate
	}
	return nil
}

func (x *Envelope) GetChannelMessageRemove() *ChannelMessageRemove {
	if x, ok := x.GetMessage().(*Envelope_ChannelMessageRemove); ok {
		return x.ChannelMessageRemove
	}
	return nil
}

func (x *Envelope) GetChannelPresenceEvent() *ChannelPresenceEvent {
	if x, ok := x.GetMessage().(*Envelope_ChannelPresenceEvent); ok {
		return x.ChannelPresenceEvent
	}
	return nil
}

func (x *Envelope) GetError() *Error {
	if x, ok := x.GetMessage().(*Envelope_Error); ok {
		return x.Error
	}
	return nil
}

func (x *Envelope) GetMatch() *Match {
	if x, ok := x.GetMessage().(*Envelope_Match); ok {
		return x.Match
	}
	return nil
}

func (x *Envelope) GetMatchCreate() *MatchCreate {
	if x, ok := x.GetMessage().(*Envelope_MatchCreate); ok {
		return x.MatchCreate
	}
	return nil
}

func (x *Envelope) GetMatchData() *MatchData {
	if x, ok := x.GetMessage().(*Envelope_MatchData); ok {
		return x.MatchData
	}
	return nil
}

func (x *Envelope) GetMatchDataSend() *MatchDataSend {
	if x, ok := x.GetMessage().(*Envelope_MatchDataSend); ok {
		return x.MatchDataSend
	}
	return nil
}

func (x *Envelope) GetMatchJoin() *MatchJoin {
	if x, ok := x.GetMessage().(*Envelope_MatchJoin); ok {
		return x.MatchJoin
	}
	return nil
}

func (x *Envelope) GetMatchLeave() *MatchLeave {
	if x, ok := x.GetMessage().(*Envelope_MatchLeave); ok {
		return x.MatchLeave
	}
	return nil
}

func (x *Envelope) GetMatchPresenceEvent() *MatchPresenceEvent {
	if x, ok := x.GetMessage().(*Envelope_MatchPresenceEvent); ok {
		return x.MatchPresenceEvent
	}
	return nil
}

func (x *Envelope) GetMatchmakerAdd() *MatchmakerAdd {
	if x, ok := x.GetMessage().(*Envelope_MatchmakerAdd); ok {
		return x.MatchmakerAdd
	}
	return nil
}

func (x *Envelope) GetMatchmakerMatched() *MatchmakerMatched {
	if x, ok := x.GetMessage().(*Envelope_MatchmakerMatched); ok {
		return x.MatchmakerMatched
	}
	return nil
}

func (x *Envelope) GetMatchmakerRemove() *MatchmakerRemove {
	if x, ok := x.GetMessage().(*Envelope_MatchmakerRemove); ok {
		return x.MatchmakerRemove
	}
	return nil
}

func (x *Envelope) GetMatchmakerTicket() *MatchmakerTicket {
	if x, ok := x.GetMessage().(*Envelope_MatchmakerTicket); ok {
		return x.MatchmakerTicket
	}
	return nil
}

func (x *Envelope) GetNotifications() *Notifications {
	if x, ok := x.GetMessage().(*Envelope_Notifications); ok {
		return x.Notifications
	}
	return nil
}

func (x *Envelope) GetRpc() *api.Rpc {
	if x, ok := x.GetMessage().(*Envelope_Rpc); ok {
		return x.Rpc
	}
	return nil
}

func (x *Envelope) GetStatus() *Status {
	if x, ok := x.GetMessage().(*Envelope_Status); ok {
		return x.Status
	}
	return nil
}

func (x *Envelope) GetStatusFollow() *StatusFollow {
	if x, ok := x.GetMessage().(*Envelope_StatusFollow); ok {
		return x.StatusFollow
	}
	return nil
}

func (x *Envelope) GetStatusPresenceEvent() *StatusPresenceEvent {
	if x, ok := x.GetMessage().(*Envelope_StatusPresenceEvent); ok {
		return x.StatusPresenceEvent
	}
	return nil
}

func (x *Envelope) GetStatusUnfollow() *StatusUnfollow {
	if x, ok := x.GetMessage().(*Envelope_StatusUnfollow); ok {
		return x.StatusUnfollow
	}
	return nil
}

func (x *Envelope) GetStatusUpdate() *StatusUpdate {
	if x, ok := x.GetMessage().(*Envelope_StatusUpdate); ok {
		return x.StatusUpdate
	}
	return nil
}

func (x *Envelope) GetStreamData() *StreamData {
	if x, ok := x.GetMessage().(*Envelope_StreamData); ok {
		return x.StreamData
	}
	return nil
}

func (x *Envelope) GetStreamPresenceEvent() *StreamPresenceEvent {
	if x, ok := x.GetMessage().(*Envelope_StreamPresenceEvent); ok {
		return x.StreamPresenceEvent
	}
	return nil
}

func (x *Envelope) GetPing() *Ping {
	if x, ok := x.GetMessage().(*Envelope_Ping); ok {
		return x.Ping
	}
	return nil
}

func (x *Envelope) GetPong() *Pong {
	if x, ok := x.GetMessage().(*Envelope_Pong); ok {
		return x.Pong
	}
	return nil
}

func (x *Envelope) GetParty() *Party {
	if x, ok := x.GetMessage().(*Envelope_Party); ok {
		return x.Party
	}
	return nil
}

func (x *Envelope) GetPartyCreate() *PartyCreate {
	if x, ok := x.GetMessage().(*Envelope_PartyCreate); ok {
		return x.PartyCreate
	}
	return nil
}

func (x *Envelope) GetPartyJoin() *PartyJoin {
	if x, ok := x.GetMessage().(*Envelope_PartyJoin); ok {
		return x.PartyJoin
	}
	return nil
}

func (x *Envelope) GetPartyLeave() *PartyLeave {
	if x, ok := x.GetMessage().(*Envelope_PartyLeave); ok {
		return x.PartyLeave
	}
	return nil
}

func (x *Envelope) GetPartyPromote() *PartyPromote {
	if x, ok := x.GetMessage().(*Envelope_PartyPromote); ok {
		return x.PartyPromote
	}
	return nil
}

func (x *Envelope) GetPartyLeader() *PartyLeader {
	if x, ok := x.GetMessage().(*Envelope_PartyLeader); ok {
		return x.PartyLeader
	}
	return nil
}

func (x *Envelope) GetPartyAccept() *PartyAccept {
	if x, ok := x.GetMessage().(*Envelope_PartyAccept); ok {
		return x.PartyAccept
	}
	return nil
}

func (x *Envelope) GetPartyRemove() *PartyRemove {
	if x, ok := x.GetMessage().(*Envelope_PartyRemove); ok {
		return x.PartyRemove
	}
	return nil
}

func (x *Envelope) GetPartyClose() *PartyClose {
	if x, ok := x.GetMessage().(*Envelope_PartyClose); ok {
		return x.PartyClose
	}
	return nil
}

func (x *Envelope) GetPartyJoinRequestList() *PartyJoinRequestList {
	if x, ok := x.GetMessage().(*Envelope_PartyJoinRequestList); ok {
		return x.PartyJoinRequestList
	}
	return nil
}

func (x *Envelope) GetPartyJoinRequest() *PartyJoinRequest {
	if x, ok := x.GetMessage().(*Envelope_PartyJoinRequest); ok {
		return x.PartyJoinRequest
	}
	return nil
}

func (x *Envelope) GetPartyMatchmakerAdd() *PartyMatchmakerAdd {
	if x, ok := x.GetMessage().(*Envelope_PartyMatchmakerAdd); ok {
		return x.PartyMatchmakerAdd
	}
	return nil
}

func (x *Envelope) GetPartyMatchmakerRemove() *PartyMatchmakerRemove {
	if x, ok := x.GetMessage().(*Envelope_PartyMatchmakerRemove); ok {
		return x.PartyMatchmakerRemove
	}
	return nil
}

func (x *Envelope) GetPartyMatchmakerTicket() *PartyMatchmakerTicket {
	if x, ok := x.GetMessage().(*Envelope_PartyMatchmakerTicket); ok {
		return x.PartyMatchmakerTicket
	}
	return nil
}

func (x *Envelope) GetPartyData() *PartyData {
	if x, ok := x.GetMessage().(*Envelope_PartyData); ok {
		return x.PartyData
	}
	return nil
}

func (x *Envelope) GetPartyDataSend() *PartyDataSend {
	if x, ok := x.GetMessage().(*Envelope_PartyDataSend); ok {
		return x.PartyDataSend
	}
	return nil
}

func (x *Envelope) GetPartyPresenceEvent() *PartyPresenceEvent {
	if x, ok := x.GetMessage().(*Envelope_PartyPresenceEvent); ok {
		return x.PartyPresenceEvent
	}
	return nil
}

type isEnvelope_Message interface {
	isEnvelope_Message()
}

type Envelope_Channel struct {
	// A response from a channel join operation.
	Channel *Channel `protobuf:"bytes,2,opt,name=channel,proto3,oneof"`
}

type Envelope_ChannelJoin struct {
	// Join a realtime chat channel.
	ChannelJoin *ChannelJoin `protobuf:"bytes,3,opt,name=channel_join,json=channelJoin,proto3,oneof"`
}

type Envelope_ChannelLeave struct {
	// Leave a realtime chat channel.
	ChannelLeave *ChannelLeave `protobuf:"bytes,4,opt,name=channel_leave,json=channelLeave,proto3,oneof"`
}

type Envelope_ChannelMessage struct {
	// An incoming message on a realtime chat channel.
	ChannelMessage *api.ChannelMessage `protobuf:"bytes,5,opt,name=channel_message,json=channelMessage,proto3,oneof"`
}

type Envelope_ChannelMessageAck struct {
	// An acknowledgement received in response to sending a message on a chat channel.
	ChannelMessageAck *ChannelMessageAck `protobuf:"bytes,6,opt,name=channel_message_ack,json=channelMessageAck,proto3,oneof"`
}

type Envelope_ChannelMessageSend struct {
	// Send a message to a realtime chat channel.
	ChannelMessageSend *ChannelMessageSend `protobuf:"bytes,7,opt,name=channel_message_send,json=channelMessageSend,proto3,oneof"`
}

type Envelope_ChannelMessageUpdate struct {
	// Update a message previously sent to a realtime chat channel.
	ChannelMessageUpdate *ChannelMessageUpdate `protobuf:"bytes,8,opt,name=channel_message_update,json=channelMessageUpdate,proto3,oneof"`
}

type Envelope_ChannelMessageRemove struct {
	// Remove a message previously sent to a realtime chat channel.
	ChannelMessageRemove *ChannelMessageRemove `protobuf:"bytes,9,opt,name=channel_message_remove,json=channelMessageRemove,proto3,oneof"`
}

type Envelope_ChannelPresenceEvent struct {
	// Presence update for a particular realtime chat channel.
	ChannelPresenceEvent *ChannelPresenceEvent `protobuf:"bytes,10,opt,name=channel_presence_event,json=channelPresenceEvent,proto3,oneof"`
}

type Envelope_Error struct {
	// Describes an error which occurred on the server.
	Error *Error `protobuf:"bytes,11,opt,name=error,proto3,oneof"`
}

type Envelope_Match struct {
	// Incoming information about a realtime match.
	Match *Match `protobuf:"bytes,12,opt,name=match,proto3,oneof"`
}

type Envelope_MatchCreate struct {
	// A client to server request to create a realtime match.
	MatchCreate *MatchCreate `protobuf:"bytes,13,opt,name=match_create,json=matchCreate,proto3,oneof"`
}

type Envelope_MatchData struct {
	// Incoming realtime match data delivered from the server.
	MatchData *MatchData `protobuf:"bytes,14,opt,name=match_data,json=matchData,proto3,oneof"`
}

type Envelope_MatchDataSend struct {
	// A client to server request to send data to a realtime match.
	MatchDataSend *MatchDataSend `protobuf:"bytes,15,opt,name=match_data_send,json=matchDataSend,proto3,oneof"`
}

type Envelope_MatchJoin struct {
	// A client to server request to join a realtime match.
	MatchJoin *MatchJoin `protobuf:"bytes,16,opt,name=match_join,json=matchJoin,proto3,oneof"`
}

type Envelope_MatchLeave struct {
	// A client to server request to leave a realtime match.
	MatchLeave *MatchLeave `protobuf:"bytes,17,opt,name=match_leave,json=matchLeave,proto3,oneof"`
}

type Envelope_MatchPresenceEvent struct {
	// Presence update for a particular realtime match.
	MatchPresenceEvent *MatchPresenceEvent `protobuf:"bytes,18,opt,name=match_presence_event,json=matchPresenceEvent,proto3,oneof"`
}

type Envelope_MatchmakerAdd struct {
	// Submit a new matchmaking process request.
	MatchmakerAdd *MatchmakerAdd `protobuf:"bytes,19,opt,name=matchmaker_add,json=matchmakerAdd,proto3,oneof"`
}

type Envelope_MatchmakerMatched struct {
	// A successful matchmaking result.
	MatchmakerMatched *MatchmakerMatched `protobuf:"bytes,20,opt,name=matchmaker_matched,json=matchmakerMatched,proto3,oneof"`
}

type Envelope_MatchmakerRemove struct {
	// Cancel a matchmaking process using a ticket.
	MatchmakerRemove *MatchmakerRemove `protobuf:"bytes,21,opt,name=matchmaker_remove,json=matchmakerRemove,proto3,oneof"`
}

type Envelope_MatchmakerTicket struct {
	// A response from starting a new matchmaking process.
	MatchmakerTicket *MatchmakerTicket `protobuf:"bytes,22,opt,name=matchmaker_ticket,json=matchmakerTicket,proto3,oneof"`
}

type Envelope_Notifications struct {
	// Notifications send by the server.
	Notifications *Notifications `protobuf:"bytes,23,opt,name=notifications,proto3,oneof"`
}

type Envelope_Rpc struct {
	// RPC call or response.
	Rpc *api.Rpc `protobuf:"bytes,24,opt,name=rpc,proto3,oneof"`
}

type Envelope_Status struct {
	// An incoming status snapshot for some set of users.
	Status *Status `protobuf:"bytes,25,opt,name=status,proto3,oneof"`
}

type Envelope_StatusFollow struct {
	// Start following some set of users to receive their status updates.
	StatusFollow *StatusFollow `protobuf:"bytes,26,opt,name=status_follow,json=statusFollow,proto3,oneof"`
}

type Envelope_StatusPresenceEvent struct {
	// An incoming status update.
	StatusPresenceEvent *StatusPresenceEvent `protobuf:"bytes,27,opt,name=status_presence_event,json=statusPresenceEvent,proto3,oneof"`
}

type Envelope_StatusUnfollow struct {
	// Stop following some set of users to no longer receive their status updates.
	StatusUnfollow *StatusUnfollow `protobuf:"bytes,28,opt,name=status_unfollow,json=statusUnfollow,proto3,oneof"`
}

type Envelope_StatusUpdate struct {
	// Set the user's own status.
	StatusUpdate *StatusUpdate `protobuf:"bytes,29,opt,name=status_update,json=statusUpdate,proto3,oneof"`
}

type Envelope_StreamData struct {
	// A data message delivered over a stream.
	StreamData *StreamData `protobuf:"bytes,30,opt,name=stream_data,json=streamData,proto3,oneof"`
}

type Envelope_StreamPresenceEvent struct {
	// Presence update for a particular stream.
	StreamPresenceEvent *StreamPresenceEvent `protobuf:"bytes,31,opt,name=stream_presence_event,json=streamPresenceEvent,proto3,oneof"`
}

type Envelope_Ping struct {
	// Application-level heartbeat and connection check.
	Ping *Ping `protobuf:"bytes,32,opt,name=ping,proto3,oneof"`
}

type Envelope_Pong struct {
	// Application-level heartbeat and connection check response.
	Pong *Pong `protobuf:"bytes,33,opt,name=pong,proto3,oneof"`
}

type Envelope_Party struct {
	// Incoming information about a party.
	Party *Party `protobuf:"bytes,34,opt,name=party,proto3,oneof"`
}

type Envelope_PartyCreate struct {
	// Create a party.
	PartyCreate *PartyCreate `protobuf:"bytes,35,opt,name=party_create,json=partyCreate,proto3,oneof"`
}

type Envelope_PartyJoin struct {
	// Join a party, or request to join if the party is not open.
	PartyJoin *PartyJoin `protobuf:"bytes,36,opt,name=party_join,json=partyJoin,proto3,oneof"`
}

type Envelope_PartyLeave struct {
	// Leave a party.
	PartyLeave *PartyLeave `protobuf:"bytes,37,opt,name=party_leave,json=partyLeave,proto3,oneof"`
}

type Envelope_PartyPromote struct {
	// Promote a new party leader.
	PartyPromote *PartyPromote `protobuf:"bytes,38,opt,name=party_promote,json=partyPromote,proto3,oneof"`
}

type Envelope_PartyLeader struct {
	// Announcement of a new party leader.
	PartyLeader *PartyLeader `protobuf:"bytes,39,opt,name=party_leader,json=partyLeader,proto3,oneof"`
}

type Envelope_PartyAccept struct {
	// Accept a request to join.
	PartyAccept *PartyAccept `protobuf:"bytes,40,opt,name=party_accept,json=partyAccept,proto3,oneof"`
}

type Envelope_PartyRemove struct {
	// Kick a party member, or decline a request to join.
	PartyRemove *PartyRemove `protobuf:"bytes,41,opt,name=party_remove,json=partyRemove,proto3,oneof"`
}

type Envelope_PartyClose struct {
	// End a party, kicking all party members and closing it.
	PartyClose *PartyClose `protobuf:"bytes,42,opt,name=party_close,json=partyClose,proto3,oneof"`
}

type Envelope_PartyJoinRequestList struct {
	// Request a list of pending join requests for a party.
	PartyJoinRequestList *PartyJoinRequestList `protobuf:"bytes,43,opt,name=party_join_request_list,json=partyJoinRequestList,proto3,oneof"`
}

type Envelope_PartyJoinRequest struct {
	// Incoming notification for one or more new presences attempting to join the party.
	PartyJoinRequest *PartyJoinRequest `protobuf:"bytes,44,opt,name=party_join_request,json=partyJoinRequest,proto3,oneof"`
}

type Envelope_PartyMatchmakerAdd struct {
	// Begin matchmaking as a party.
	PartyMatchmakerAdd *PartyMatchmakerAdd `protobuf:"bytes,45,opt,name=party_matchmaker_add,json=partyMatchmakerAdd,proto3,oneof"`
}

type Envelope_PartyMatchmakerRemove struct {
	// Cancel a party matchmaking process using a ticket.
	PartyMatchmakerRemove *PartyMatchmakerRemove `protobuf:"bytes,46,opt,name=party_matchmaker_remove,json=partyMatchmakerRemove,proto3,oneof"`
}

type Envelope_PartyMatchmakerTicket struct {
	// A response from starting a new party matchmaking process.
	PartyMatchmakerTicket *PartyMatchmakerTicket `protobuf:"bytes,47,opt,name=party_matchmaker_ticket,json=partyMatchmakerTicket,proto3,oneof"`
}

type Envelope_PartyData struct {
	// Incoming party data delivered from the server.
	PartyData *PartyData `protobuf:"bytes,48,opt,name=party_data,json=partyData,proto3,oneof"`
}

type Envelope_PartyDataSend struct {
	// A client to server request to send data to a party.
	PartyDataSend *PartyDataSend `protobuf:"bytes,49,opt,name=party_data_send,json=partyDataSend,proto3,oneof"`
}

type Envelope_PartyPresenceEvent struct {
	// Presence update for a particular party.
	PartyPresenceEvent *PartyPresenceEvent `protobuf:"bytes,50,opt,name=party_presence_event,json=partyPresenceEvent,proto3,oneof"`
}

func (*Envelope_Channel) isEnvelope_Message() {}

func (*Envelope_ChannelJoin) isEnvelope_Message() {}

func (*Envelope_ChannelLeave) isEnvelope_Message() {}

func (*Envelope_ChannelMessage) isEnvelope_Message() {}

func (*Envelope_ChannelMessageAck) isEnvelope_Message() {}

func (*Envelope_ChannelMessageSend) isEnvelope_Message() {}

func (*Envelope_ChannelMessageUpdate) isEnvelope_Message() {}

func (*Envelope_ChannelMessageRemove) isEnvelope_Message() {}

func (*Envelope_ChannelPresenceEvent) isEnvelope_Message() {}

func (*Envelope_Error) isEnvelope_Message() {}

func (*Envelope_Match) isEnvelope_Message() {}

func (*Envelope_MatchCreate) isEnvelope_Message() {}

func (*Envelope_MatchData) isEnvelope_Message() {}

func (*Envelope_MatchDataSend) isEnvelope_Message() {}

func (*Envelope_MatchJoin) isEnvelope_Message() {}

func (*Envelope_MatchLeave) isEnvelope_Message() {}

func (*Envelope_MatchPresenceEvent) isEnvelope_Message() {}

func (*Envelope_MatchmakerAdd) isEnvelope_Message() {}

func (*Envelope_MatchmakerMatched) isEnvelope_Message() {}

func (*Envelope_MatchmakerRemove) isEnvelope_Message() {}

func (*Envelope_MatchmakerTicket) isEnvelope_Message() {}

func (*Envelope_Notifications) isEnvelope_Message() {}

func (*Envelope_Rpc) isEnvelope_Message() {}

func (*Envelope_Status) isEnvelope_Message() {}

func (*Envelope_StatusFollow) isEnvelope_Message() {}

func (*Envelope_StatusPresenceEvent) isEnvelope_Message() {}

func (*Envelope_StatusUnfollow) isEnvelope_Message() {}

func (*Envelope_StatusUpdate) isEnvelope_Message() {}

func (*Envelope_StreamData) isEnvelope_Message() {}

func (*Envelope_StreamPresenceEvent) isEnvelope_Message() {}

func (*Envelope_Ping) isEnvelope_Message() {}

func (*Envelope_Pong) isEnvelope_Message() {}

func (*Envelope_Party) isEnvelope_Message() {}

func (*Envelope_PartyCreate) isEnvelope_Message() {}

func (*Envelope_PartyJoin) isEnvelope_Message() {}

func (*Envelope_PartyLeave) isEnvelope_Message() {}

func (*Envelope_PartyPromote) isEnvelope_Message() {}

func (*Envelope_PartyLeader) isEnvelope_Message() {}

func (*Envelope_PartyAccept) isEnvelope_Message() {}

func (*Envelope_PartyRemove) isEnvelope_Message() {}

func (*Envelope_PartyClose) isEnvelope_Message() {}

func (*Envelope_PartyJoinRequestList) isEnvelope_Message() {}

func (*Envelope_PartyJoinRequest) isEnvelope_Message() {}

func (*Envelope_PartyMatchmakerAdd) isEnvelope_Message() {}

func (*Envelope_PartyMatchmakerRemove) isEnvelope_Message() {}

func (*Envelope_PartyMatchmakerTicket) isEnvelope_Message() {}

func (*Envelope_PartyData) isEnvelope_Message() {}

func (*Envelope_PartyDataSend) isEnvelope_Message() {}

func (*Envelope_PartyPresenceEvent) isEnvelope_Message() {}
