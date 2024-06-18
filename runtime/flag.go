package runtime

type FlagList struct {
	Flags []*Flag `json:"flags,omitempty"`
}

type Flag struct {
	Name             string `json:"name,omitempty"`
	Value            string `json:"value,omitempty"`
	ConditionChanged bool   `json:"condition_changed,omitempty"`
}
