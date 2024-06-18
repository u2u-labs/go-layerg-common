package runtime

type PresenceReason uint8

const (
	PresenceReasonUnknown PresenceReason = iota
	PresenceReasonJoin
	PresenceReasonUpdate
	PresenceReasonLeave
	PresenceReasonDisconnect
)

type PresenceMeta interface {
	GetHidden() bool
	GetPersistence() bool
	GetUsername() string
	GetStatus() string
	GetReason() PresenceReason
}

type Presence interface {
	PresenceMeta
	GetUserId() string
	GetSessionId() string
	GetNodeId() string
}
