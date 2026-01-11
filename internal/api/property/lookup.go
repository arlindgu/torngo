package property

import "torngo/internal/api"

type PropertyLookupResponse struct {
	Selections []string `json:"selections"`
}

type PropertyLookupParams struct {
	Timestamp api.ApiTimestamp
	Comment   api.ApiComment
}
