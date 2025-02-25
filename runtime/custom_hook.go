package runtime

import "context"

type WebhookConfig struct {
	Path        string                 `json:"path"`
	Method      string                 `json:"method"`
	Description string                 `json:"description"`
	Params      map[string]interface{} `json:"params"`
}

type WebhookHandler func(ctx context.Context, payload map[string]interface{}) (interface{}, error)
