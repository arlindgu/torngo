package id

import "torngo/internal/api"

type MarketIdBazaarResponse struct {
	Bazaar struct {
		Specialized []struct {
			ID              int    `json:"id"`
			Name            string `json:"name"`
			IsOpen          bool   `json:"is_open"`
			WeeklyCustomers int    `json:"weekly_customers"`
		} `json:"specialized"`
	} `json:"bazaar"`
}

type MarketIdBazaarParams struct {
	ItemId    int64
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
