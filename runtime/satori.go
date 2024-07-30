package runtime

import "context"

/*
Satori runtime integration definitions.
*/
type Satori interface {
	Authenticate(ctx context.Context, id string, ipAddress ...string) error
	PropertiesGet(ctx context.Context, id string) (*Properties, error)
	PropertiesUpdate(ctx context.Context, id string, properties *PropertiesUpdate) error
	EventsPublish(ctx context.Context, id string, events []*Event) error
	ExperimentsList(ctx context.Context, id string, names ...string) (*ExperimentList, error)
	FlagsList(ctx context.Context, id string, names ...string) (*FlagList, error)
	LiveEventsList(ctx context.Context, id string, names ...string) (*LiveEventList, error)
}
