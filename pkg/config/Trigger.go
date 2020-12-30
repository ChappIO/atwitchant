package config

import "regexp"

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
