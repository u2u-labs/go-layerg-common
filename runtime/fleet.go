package runtime

import "context"

type FmCreateStatus int

const (
	// Create successfully created a new game instance.
	CreateSuccess FmCreateStatus = iota
	// Create request could not find a suitable instance within the configured timeout.
	CreateTimeout
	// Create failed to create a new game instance.
	CreateError
)

type FmCallbackHandler interface {
	// Generate a new callback id.
	GenerateCallbackId() string
	// Set the callback indexed by the generated id.
	SetCallback(callbackId string, fn FmCreateCallbackFn)
	// Invoke a callback by callback Id.
	InvokeCallback(callbackId string, status FmCreateStatus, instanceInfo *InstanceInfo, sessionInfo []*SessionInfo, metadata map[string]any, err error)
}

type FleetUserLatencies struct {
	// User id
	UserId string
	// Latency experienced by the user contacting a server in a fleet instance region.
	LatencyInMilliseconds float32
	// Region associated to the experienced latency value.
	RegionIdentifier string
}

// FmCreateCallbackFn is the function that is invoked when Create asynchronously succeeds or fails (due to timeout or issues bringing up a new instance).
// The function params include all the information needed to inform a client with a realtime connection to the server of the status of the Create request,
// including the new instance connection information in case of success.
// If status != CreateSuccess, then instanceInfo, sessionInfo and metadata will be nil and err will contain an error message.
// If no userIds were provided to Create, then sessionInfo will be nil regardless of successful instance creation.
type FmCreateCallbackFn func(status FmCreateStatus, instanceInfo *InstanceInfo, sessionInfo []*SessionInfo, metadata map[string]any, err error)

type FleetManager interface {
	// Get retrieves the most up-to-date information about an instance currently running
	// in the Fleet Manager platform. An error is expected if the instance does not exist,
	// either because it never existed or it was otherwise removed at some point.
	Get(ctx context.Context, id string) (instance *InstanceInfo, err error)

	// List retrieves a set of instances, optionally filtered by a platform-specific query.
	// The limit and previous cursor inputs are used as part of pagination, if supported.
	List(ctx context.Context, query string, limit int, previousCursor string) (list []*InstanceInfo, nextCursor string, err error)

	// Create issues a request to the underlying Fleet Manager platform to create a new
	// instance and initialize it with the given metadata. The metadata is expected to be
	// application-specific and only relevant to the application itself, not the platform.
	// The instance creation happens asynchronously - the passed callback is invoked once the
	// creation process was either successful or failed.
	// If a list of userIds is optionally provided, the new instance (on successful creation) will reserve slots
	// for the respective clients to connect, and the callback will contain the required []*SessionInfo.
	// Latencies is optional and its support depends on the Fleet Manager provider.
	Create(ctx context.Context, maxPlayers int, userIds []string, latencies []FleetUserLatencies, metadata map[string]any, callback FmCreateCallbackFn) (err error)

	// Join reserves a number of player slots in the target instance. These slots are reserved for a minute, after which,
	// if clients do not connect to the instance to claim them, the returned SessionInfo will become invalid and the
	// player slots will become available to new player sessions.
	Join(ctx context.Context, id string, userIds []string, metadata map[string]string) (joinInfo *JoinInfo, err error)
}

type FleetManagerInitializer interface {
	FleetManager
	// Init function - it is called internally by RegisterFleetManager to expose LayergModule and FmCallbackHandler.
	// The implementation should keep references to nk and callbackHandler.
	Init(layerg LayerGModule, callbackHandler FmCallbackHandler) error
	Update(ctx context.Context, id string, playerCount int, metadata map[string]any) error
	Delete(ctx context.Context, id string) error
}
