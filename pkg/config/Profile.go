package config

type Profile struct {
	Triggers map[string]Trigger `json:"triggers"`
	Actions  map[string]Action  `json:"actions"`
}
