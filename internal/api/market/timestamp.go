package market

import "torngo/internal/api"

type MarketTimestampResponse struct {
	Timestamp int `json:"timestamp"`
}

type MarketTimestampParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
