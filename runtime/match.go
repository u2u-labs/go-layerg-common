package runtime

import (
	"context"
	"database/sql"
)

type MatchmakerEntry interface {
	GetPresence() Presence
	GetTicket() string
	GetProperties() map[string]interface{}
	GetPartyId() string
}

type MatchData interface {
	Presence
	GetOpCode() int64
	GetData() []byte
	GetReliable() bool
	GetReceiveTime() int64
}

type MatchDispatcher interface {
	BroadcastMessage(opCode int64, data []byte, presences []Presence, sender Presence, reliable bool) error
	BroadcastMessageDeferred(opCode int64, data []byte, presences []Presence, sender Presence, reliable bool) error
	MatchKick(presences []Presence) error
	MatchLabelUpdate(label string) error
}

type Match interface {
	MatchInit(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, params map[string]interface{}) (interface{}, int, string)
	MatchJoinAttempt(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, dispatcher MatchDispatcher, tick int64, state interface{}, presence Presence, metadata map[string]string) (interface{}, bool, string)
	MatchJoin(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, dispatcher MatchDispatcher, tick int64, state interface{}, presences []Presence) interface{}
	MatchLeave(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, dispatcher MatchDispatcher, tick int64, state interface{}, presences []Presence) interface{}
	MatchLoop(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, dispatcher MatchDispatcher, tick int64, state interface{}, messages []MatchData) interface{}
	MatchTerminate(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, dispatcher MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{}
	MatchSignal(ctx context.Context, logger Logger, db *sql.DB, layerg LayerGModule, dispatcher MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string)
}
