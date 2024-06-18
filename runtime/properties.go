package runtime

type Properties struct {
	Default  map[string]string `json:"default,omitempty"`
	Custom   map[string]string `json:"custom,omitempty"`
	Computed map[string]string `json:"computed,omitempty"`
}

type PropertiesUpdate struct {
	Default   map[string]string `json:"default,omitempty"`
	Custom    map[string]string `json:"custom,omitempty"`
	Recompute *bool             `json:"recompute,omitempty"`
}
