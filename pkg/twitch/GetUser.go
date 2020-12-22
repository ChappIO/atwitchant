package twitch

import (
	"encoding/json"
	"net/http"
)

type userData struct {
	Id              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Description     string `json:"description"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type getUserResponse struct {
	Data []userData `json:"data"`
}

func (t *Integration) GetUser() userData {
	req, _ := http.NewRequest(
		"GET",
		"https://api.twitch.tv/helix/users",
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Set("Client-Id", clientId)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	response := getUserResponse{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		panic(err)
	}
	return response.Data[0]
}
