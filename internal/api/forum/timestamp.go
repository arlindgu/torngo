package forum

import "torngo/internal/api"

type ForumTimestampResponse struct {
	Timestamp int `json:"timestamp"`
}

type ForumTimestampParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
