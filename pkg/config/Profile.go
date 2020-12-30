package config

type Profile struct {
	Triggers map[string]Trigger `json:"triggers"`
}
