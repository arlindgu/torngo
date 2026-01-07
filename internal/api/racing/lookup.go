package racing

import "torngo/internal/api"

type RacingLookupResponse struct {
	Selections []string `json:"selections"`
}

type RacingLookupParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
