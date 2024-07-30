package runtime

import (
	"errors"
	"fmt"
)

/*
Error is used to indicate a failure in code. The message and code are returned to the client.
If an Error is used as response for a HTTP/gRPC request, then the server tries to use the error value as the gRPC error code. This will in turn translate to HTTP status codes.

For more information, please have a look at the following:

	https://github.com/grpc/grpc-go/blob/master/codes/codes.go
	https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
	https://golang.org/pkg/net/http/
*/
type Error struct {
	Message string
	Code    int
}

// Error returns the encapsulated error message.
func (e *Error) Error() string {
	return e.Message
}

/*
NewError returns a new error. The message and code are sent directly to the client. The code field is also optionally translated to gRPC/HTTP code.

	runtime.NewError("Server unavailable", 14) // 14 = Unavailable = 503 HTTP status code
*/
func NewError(message string, code int) *Error {
	return &Error{Message: message, Code: code}
}

func (e *WalletNegativeError) Error() string {
	return fmt.Sprintf("wallet update rejected negative value at path '%v'", e.Path)
}

var (
	ErrStorageRejectedVersion    = errors.New("storage write rejected - version check failed")
	ErrStorageRejectedPermission = errors.New("storage write rejected - permission denied")

	ErrChannelIDInvalid     = errors.New("invalid channel id")
	ErrChannelCursorInvalid = errors.New("invalid channel cursor")
	ErrChannelGroupNotFound = errors.New("group not found")

	ErrInvalidChannelTarget = errors.New("invalid channel target")
	ErrInvalidChannelType   = errors.New("invalid channel type")

	ErrFriendInvalidCursor = errors.New("friend cursor invalid")

	ErrLeaderboardNotFound = errors.New("leaderboard not found")

	ErrTournamentNotFound                = errors.New("tournament not found")
	ErrTournamentAuthoritative           = errors.New("tournament only allows authoritative submissions")
	ErrTournamentMaxSizeReached          = errors.New("tournament max size reached")
	ErrTournamentOutsideDuration         = errors.New("tournament outside of duration")
	ErrTournamentWriteMaxNumScoreReached = errors.New("max number score count reached")
	ErrTournamentWriteJoinRequired       = errors.New("required to join before writing tournament record")

	ErrMatchmakerQueryInvalid     = errors.New("matchmaker query invalid")
	ErrMatchmakerDuplicateSession = errors.New("matchmaker duplicate session")
	ErrMatchmakerIndex            = errors.New("matchmaker index error")
	ErrMatchmakerDelete           = errors.New("matchmaker delete error")
	ErrMatchmakerNotAvailable     = errors.New("matchmaker not available")
	ErrMatchmakerTooManyTickets   = errors.New("matchmaker too many tickets")
	ErrMatchmakerTicketNotFound   = errors.New("matchmaker ticket not found")

	ErrPartyClosed                   = errors.New("party closed")
	ErrPartyFull                     = errors.New("party full")
	ErrPartyJoinRequestDuplicate     = errors.New("party join request duplicate")
	ErrPartyJoinRequestAlreadyMember = errors.New("party join request already member")
	ErrPartyJoinRequestsFull         = errors.New("party join requests full")
	ErrPartyNotLeader                = errors.New("party leader only")
	ErrPartyNotMember                = errors.New("party member not found")
	ErrPartyNotRequest               = errors.New("party join request not found")
	ErrPartyAcceptRequest            = errors.New("party could not accept request")
	ErrPartyRemove                   = errors.New("party could not remove")
	ErrPartyRemoveSelf               = errors.New("party cannot remove self")

	ErrGracePeriodExpired = errors.New("grace period expired")

	ErrGroupNameInUse         = errors.New("group name in use")
	ErrGroupPermissionDenied  = errors.New("group permission denied")
	ErrGroupNoUpdateOps       = errors.New("no group updates")
	ErrGroupNotUpdated        = errors.New("group not updated")
	ErrGroupNotFound          = errors.New("group not found")
	ErrGroupFull              = errors.New("group is full")
	ErrGroupUserNotFound      = errors.New("user not found")
	ErrGroupLastSuperadmin    = errors.New("user is last group superadmin")
	ErrGroupUserInvalidCursor = errors.New("group user cursor invalid")
	ErrUserGroupInvalidCursor = errors.New("user group cursor invalid")
	ErrGroupCreatorInvalid    = errors.New("group creator user ID not valid")

	ErrWalletLedgerInvalidCursor = errors.New("wallet ledger cursor invalid")

	ErrCannotEncodeParams    = errors.New("error creating match: cannot encode params")
	ErrCannotDecodeParams    = errors.New("error creating match: cannot decode params")
	ErrMatchIdInvalid        = errors.New("match id invalid")
	ErrMatchNotFound         = errors.New("match not found")
	ErrMatchBusy             = errors.New("match busy")
	ErrMatchStateFailed      = errors.New("match did not return state")
	ErrMatchLabelTooLong     = errors.New("match label too long, must be 0-2048 bytes")
	ErrDeferredBroadcastFull = errors.New("too many deferred message broadcasts per tick")

	ErrSatoriConfigurationInvalid = errors.New("satori configuration is invalid")
)
