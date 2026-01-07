package forum

import "torngo/internal/api"

type ForumLookupResponse struct {
	Selections []string `json:"selections"`
}

type ForumLookupParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
