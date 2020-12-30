package config

type SendMessageAction struct {
	Template string
}

type Action struct {
	SendMessage *SendMessageAction `json:"send_message,omitempty"`
}
