package racing

import "torngo/internal/api"

type RacingTimestampResponse struct {
	Timestamp int `json:"timestamp"`
}

type RacingTimestampParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
