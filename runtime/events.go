package runtime

type Events struct {
	Events []*Event
}

type Event struct {
	Name      string            `json:"name,omitempty"`
	Id        string            `json:"id,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	Value     string            `json:"value,omitempty"`
	Timestamp int64             `json:"-"`
}

type LiveEventList struct {
	LiveEvents []*LiveEvent `json:"live_events,omitempty"`
}

type LiveEvent struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	Value              string `json:"value,omitempty"`
	ActiveStartTimeSec int64  `json:"active_start_time_sec,string,omitempty"`
	ActiveEndTimeSec   int64  `json:"active_end_time_sec,string,omitempty"`
}
