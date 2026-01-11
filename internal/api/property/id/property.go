package id

import "torngo/internal/api"

type PropertyIdPropertyResponse struct {
	Property struct {
		ID    int64 `json:"id"`
		Owner struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"owner"`
		Property struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"property"`
		Happy  int `json:"happy"`
		Upkeep struct {
			Property int `json:"property"`
			Staff    int `json:"staff"`
		} `json:"upkeep"`
		MarketPrice   int64    `json:"market_price"`
		Modifications []string `json:"modifications"`
		Staff         []struct {
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"staff"`
		UsedBy []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"used_by"`
		Status   string `json:"status"`
		RentedBy struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"rented_by"`
	} `json:"property"`
}

type PropertyIdPropertyParams struct {
	PropertyId int64
	Timestamp  api.ApiTimestamp
	Comment    api.ApiComment
}
