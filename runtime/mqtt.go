package runtime

import (
	"context"
)

type MQTTConfig struct {
	Topic       string                 `json:"topic"`
	QoS         byte                   `json:"qos"`
	Description string                 `json:"description"`
	Params      map[string]interface{} `json:"params"`
	BrokerURL   string                 `json:"broker_url"`
	ClientID    string                 `json:"client_id"`
}

type MQTTHandler func(ctx context.Context, message MQTTMessage) error

type MQTTMessage struct {
	Topic   string
	Payload []byte
}
