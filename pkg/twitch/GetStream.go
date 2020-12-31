package twitch

import "errors"

type StreamData struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	GameID       string `json:"game_id"`
	GameName     string `json:"game_name"`
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	Language     string `json:"language"`
	StartedAt    string `json:"started_at"`
	Type         string `json:"type"`
	ViewerCount  int    `json:"viewer_count"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type getStreamResponse struct {
	Data []StreamData `json:"data"`
}

func (t *Integration) GetStream() (StreamData, error) {
	response := getStreamResponse{}
	err := t.httpGet("https://api.twitch.tv/helix/streams&user_id="+t.User.Id, &response)
	if err != nil {
		return StreamData{}, err
	}
	if len(response.Data) > 0 {
		return response.Data[0], nil
	} else {
		return StreamData{}, errors.New("no active stream found")
	}
}
