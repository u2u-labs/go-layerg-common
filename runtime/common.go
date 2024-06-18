package runtime

const (
	// All available environmental variables made available to the runtime environment.
	// This is useful to store API keys and other secrets which may be different between servers run in production and in development.
	//   envs := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	// This can always be safely cast into a `map[string]string`.
	RUNTIME_CTX_ENV = "env"

	// The mode associated with the execution context. It's one of these values:
	//  "event", "run_once", "rpc", "before", "after", "match", "matchmaker", "leaderboard_reset", "tournament_reset", "tournament_end".
	RUNTIME_CTX_MODE = "execution_mode"

	// The node ID where the current runtime context is executing.
	RUNTIME_CTX_NODE = "node"

	// Server version.
	RUNTIME_CTX_VERSION = "version"

	// Http headers. Only applicable to HTTP RPC requests.
	RUNTIME_CTX_HEADERS = "headers"

	// Query params that was passed through from HTTP request.
	RUNTIME_CTX_QUERY_PARAMS = "query_params"

	// The user ID associated with the execution context.
	RUNTIME_CTX_USER_ID = "user_id"

	// The username associated with the execution context.
	RUNTIME_CTX_USERNAME = "username"

	// Variables stored in the user's session token.
	RUNTIME_CTX_VARS = "vars"

	// The user session expiry in seconds associated with the execution context.
	RUNTIME_CTX_USER_SESSION_EXP = "user_session_exp"

	// The user session associated with the execution context.
	RUNTIME_CTX_SESSION_ID = "session_id"

	// The user session's lang value, if one is set.
	RUNTIME_CTX_LANG = "lang"

	// The IP address of the client making the request.
	RUNTIME_CTX_CLIENT_IP = "client_ip"

	// The port number of the client making the request.
	RUNTIME_CTX_CLIENT_PORT = "client_port"

	// The match ID that is currently being executed. Only applicable to server authoritative multiplayer.
	RUNTIME_CTX_MATCH_ID = "match_id"

	// The node ID that the match is being executed on. Only applicable to server authoritative multiplayer.
	RUNTIME_CTX_MATCH_NODE = "match_node"

	// Labels associated with the match. Only applicable to server authoritative multiplayer.
	RUNTIME_CTX_MATCH_LABEL = "match_label"

	// Tick rate defined for this match. Only applicable to server authoritative multiplayer.
	RUNTIME_CTX_MATCH_TICK_RATE = "match_tick_rate"
)

const (
	// Storage permission for public read, any user can read the object.
	STORAGE_PERMISSION_PUBLIC_READ = 2

	// Storage permission for owner read, only the user who owns it may access.
	STORAGE_PERMISSION_OWNER_READ = 1

	// Storage permission for no read. The object is only readable by server runtime.
	STORAGE_PERMISSION_NO_READ = 0

	// Storage permission for owner write, only the user who owns it may write.
	STORAGE_PERMISSION_OWNER_WRITE = 1

	// Storage permission for no write. The object is only writable by server runtime.
	STORAGE_PERMISSION_NO_WRITE = 0
)

type ChannelType int

const (
	Room ChannelType = iota + 1
	DirectMessage
	Group
)
