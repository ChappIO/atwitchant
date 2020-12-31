package twitch

import (
	"context"
	"strings"
)

type ChatMessage struct {
	Tags     map[string]string
	Command  string
	Sender   string
	SenderID string
	Channel  string
	Body     string
	Context  context.Context
}

const CommandPrivMsg = "PRIVMSG"
const CommandPing = "PING"

func messageFromLine(line string) (out ChatMessage) {
	remaining := line

	// Tags begin with the @
	out.Tags = make(map[string]string)
	if strings.HasPrefix(remaining, "@") {
		// find the tags terminator
		endOfTags := strings.Index(remaining, " :")
		tagsLine := remaining[1:endOfTags]
		remaining = remaining[endOfTags+1:]
		tagParts := strings.Split(tagsLine, ";")
		for _, part := range tagParts {
			equalsSign := strings.Index(part, "=")
			key := part[0:equalsSign]
			value := part[equalsSign+1:]
			out.Tags[key] = value
		}
	}

	if strings.HasPrefix(remaining, ":") {
		// get sender info
		nextSpace := strings.Index(remaining, " ")
		out.Sender = remaining[1:nextSpace]
		remaining = remaining[nextSpace+1:]
	}

	// get command
	nextSpace := strings.Index(remaining, " ")
	out.Command = remaining[0:nextSpace]
	remaining = remaining[nextSpace+1:]

	// get channel
	nextSpace = strings.Index(remaining, " ")
	if nextSpace == -1 {
		// there is no body
		out.Channel = remaining
	} else {
		// there is a body after the channel
		out.Channel = remaining[0:nextSpace]
		out.Body = strings.TrimPrefix(remaining[nextSpace+1:], ":")
	}

	// twitch-specific parsing
	if value, ok := out.Tags["user-id"]; ok {
		out.SenderID = value
	}
	return
}
