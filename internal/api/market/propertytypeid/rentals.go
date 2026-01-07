package propertytypeid

import "torngo/internal/api"

type MarketPropertyIdRentalsResponse struct {
	Rentals struct {
		Listings []struct {
			Happy         int      `json:"happy"`
			Cost          int64    `json:"cost"`
			CostPerDay    int64    `json:"cost_per_day"`
			RentalPeriod  int      `json:"rental_period"`
			MarketPrice   int64    `json:"market_price"`
			Upkeep        int      `json:"upkeep"`
			Modifications []string `json:"modifications"`
		} `json:"listings"`
		Property struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"property"`
	} `json:"rentals"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type MarketPropertyIdRentalsParams struct {
	PropertyTypeId int32
	Offset         api.ApiOffset
	Limit          api.ApiLimit
	Sort           api.ApiSort
	Timestamp      api.ApiTimestamp
	Comment        api.ApiComment
}
