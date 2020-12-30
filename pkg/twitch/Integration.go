package twitch

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
)

const clientId = "tkro9r2rqee1s95hhyecyfq979lky7"
const twitchLoginFile = "twitch.json"

type Integration struct {
	Token string   `json:"token"`
	User  UserData `json:"user_details"`
	chat  *Chat
}

func (t *Integration) Connect() error {
	if usr, err := t.GetUser(); err != nil {
		return nil
	} else {
		t.User = usr
	}
	t.chat = &Chat{
		api: t,
	}
	if err := t.chat.Reconnect(); err != nil {
		return err
	}
	return nil
}

func LoadTwitch() Integration {
	file, err := os.Open(twitchLoginFile)
	if os.IsNotExist(err) {
		log.Println("no twitch session was found")
		return Integration{}
	}
	if err != nil {
		panic(err)
	}
	defer file.Close()
	out := Integration{}

	err = json.NewDecoder(file).Decode(&out)
	if err != nil {
		panic(err)
	}
	if out.Token != "" {
		out.Connect()
	}
	return out
}

func (t *Integration) Save() {
	file, err := os.Create(twitchLoginFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(t); err != nil {
		panic(err)
	}
}

func AuthorizeUrl(loginUri string) string {
	redirectUri := url.URL{
		Scheme: "https",
		Host:   "id.twitch.tv",
		Path:   "/oauth2/authorize",
		RawQuery: url.Values{
			"client_id":     []string{clientId},
			"redirect_uri":  []string{loginUri},
			"response_type": []string{"token"},
			"scope":         []string{"chat:read chat:edit channel:moderate whispers:read whispers:edit channel_editor"},
		}.Encode(),
	}
	return redirectUri.String()
}
