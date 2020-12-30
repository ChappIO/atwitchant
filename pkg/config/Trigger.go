package config

import (
	"atwitchant/pkg/twitch"
	"regexp"
)

type Regexp struct {
	regexp.Regexp
}

func (r *Regexp) UnmarshalText(text []byte) error {
	rr, err := regexp.Compile(string(text))
	if err != nil {
		return err
	}
	*r = Regexp{Regexp: *rr}
	return nil
}

func (r *Regexp) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func MustCompile(text string) *Regexp {
	return &Regexp{Regexp: *regexp.MustCompile(text)}
}

type Match struct {
	Comment string  `json:"_comment,omitempty"`
	Message *Regexp `json:"msg"`
}

type Trigger struct {
	Comment string `json:"_comment,omitempty"`
	Match   Match  `json:"match"`
	Action  string `json:"action"`
}

func (t *Trigger) Check(message *twitch.ChatMessage) int {
	matchPosition := -1
	if t.Match.Message != nil {
		match := t.Match.Message.FindStringIndex(message.Body)
		if len(match) > 1 {
			return match[0]
		}
	}
	return matchPosition
}
