package racing

import "torngo/internal/api"

type RacingTracksResponse struct {
	Tracks []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"tracks"`
}

type RacingTracksParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
