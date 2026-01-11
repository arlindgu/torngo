package property

import "torngo/internal/api"

type PropertyTimestampResponse struct {
	Timestamp int `json:"timestamp"`
}

type PropertyTimestampParams struct {
	Timestamp api.ApiTimestamp
	Comment   api.ApiComment
}
