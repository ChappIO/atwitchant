package twitch

import (
	"encoding/json"
)

type UserData struct {
	Id              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Description     string `json:"description"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type getUserResponse struct {
	Data []UserData `json:"data"`
}

func (t *Integration) httpGet(url string, target interface{}) error {
	data, err := t.cache.Get(t.Token, url)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, target); err != nil {
		return err
	}
	return nil
}

func (t *Integration) GetUser() (UserData, error) {
	response := getUserResponse{}
	err := t.httpGet("https://api.twitch.tv/helix/users", &response)
	if err != nil {
		return UserData{}, err
	}
	return response.Data[0], nil
}

func (t *Integration) GetUserById(id string) (UserData, error) {
	response := getUserResponse{}
	err := t.httpGet("https://api.twitch.tv/helix/users?id="+id, &response)
	if err != nil {
		return UserData{}, err
	}
	return response.Data[0], nil
}
