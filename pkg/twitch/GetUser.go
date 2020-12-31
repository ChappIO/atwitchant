package twitch

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
