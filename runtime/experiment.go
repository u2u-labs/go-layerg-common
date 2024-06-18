package runtime

type ExperimentList struct {
	Experiments []*Experiment `json:"experiments,omitempty"`
}

type Experiment struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
