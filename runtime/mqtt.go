package runtime

import (
	"context"
	"time"
)

// MQTTMessage represents a message received from MQTT
type MQTTMessage struct {
	Topic    string
	Payload  []byte
	QoS      byte
	Retained bool
}

// MQTTSubscription represents an MQTT subscription configuration
type MQTTSubscription struct {
	Topic   string
	QoS     byte
	Handler func(context.Context, *MQTTMessage) error
}

// MQTTConfig represents the configuration for MQTT client
type MQTTConfig struct {
	BrokerURL      string
	ClientID       string
	Username       string
	Password       string
	CleanSession   bool
	KeepAlive      time.Duration
	ConnectTimeout time.Duration
}

// MQTTModule defines the interface for MQTT operations
type MQTTModule interface {
	// Connect establishes a connection to the MQTT broker
	Connect(ctx context.Context, config MQTTConfig) error

	// Disconnect closes the connection to the MQTT broker
	Disconnect(ctx context.Context) error

	// Subscribe subscribes to one or more topics
	Subscribe(ctx context.Context, subscriptions []MQTTSubscription) error

	// Unsubscribe unsubscribes from one or more topics
	Unsubscribe(ctx context.Context, topics []string) error

	// Publish publishes a message to a topic
	Publish(ctx context.Context, topic string, payload []byte, qos byte, retained bool) error

	// IsConnected returns whether the client is currently connected
	IsConnected() bool
}

// MQTTModuleProvider is a function type that creates a new MQTT module instance
type MQTTModuleProvider func() MQTTModule
