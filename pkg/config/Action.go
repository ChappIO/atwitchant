package config

type SendMessageAction struct {
	Template string
}

type Action struct {
	Comment     string             `json:"_comment,omitempty"`
	SendMessage *SendMessageAction `json:"send_message,omitempty"`
}
