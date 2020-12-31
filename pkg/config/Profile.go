package config

type Profile struct {
	Comment  string             `json:"_comment,omitempty"`
	Triggers map[string]Trigger `json:"triggers"`
	Actions  map[string]Action  `json:"actions"`
}
