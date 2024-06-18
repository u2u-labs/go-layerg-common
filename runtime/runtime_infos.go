package runtime

import "time"

/*
LayerG fleet manager definitions.
*/
type InstanceInfo struct {
	// A platform-specific unique instance identifier. Identifiers may be recycled for
	// future use, but the underlying Fleet Manager platform is expected to ensure
	// uniqueness at least among concurrently running instances.
	Id string `json:"id"`
	// Connection information in a platform-specific format, usually "address:port"
	ConnectionInfo *ConnectionInfo `json:"connection_info"`
	// When this instance was first created.
	CreateTime time.Time `json:"create_time"`
	// Number of active player sessions on the server
	PlayerCount int `json:"player_count"`
	// Status
	Status string `json:"status"`
	// Application-specific data for use in indexing and listings.
	Metadata map[string]any `json:"metadata"`
}

type ConnectionInfo struct {
	IpAddress string `json:"ip_address"`
	DnsName   string `json:"dns_name"`
	Port      int    `json:"port"`
}

type JoinInfo struct {
	InstanceInfo *InstanceInfo  `json:"instance_info"`
	SessionInfo  []*SessionInfo `json:"session_info"`
}

type SessionInfo struct {
	UserId    string `json:"user_id"`
	SessionId string `json:"session_id"`
}
