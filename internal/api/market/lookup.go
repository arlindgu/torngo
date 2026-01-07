package market

import "torngo/internal/api"

type MarketLookupResponse struct {
	Selections []string `json:"selections"`
}

type MarketLookupParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
