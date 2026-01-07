package propertytypeid

import "torngo/internal/api"

type MarketPropertyIdPropertiesResponse struct {
	Properties struct {
		Listings []struct {
			Happy         int      `json:"happy"`
			Cost          int64    `json:"cost"`
			MarketPrice   int64    `json:"market_price"`
			Upkeep        int      `json:"upkeep"`
			Modifications []string `json:"modifications"`
		} `json:"listings"`
		Property struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"property"`
	} `json:"properties"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type MarketPropertyIdPropertiesParams struct {
	PropertyTypeId int32
	Offset         api.ApiOffset
	Limit          api.ApiLimit
	Sort           api.ApiSort
	Timestamp      api.ApiTimestamp
	Comment        api.ApiComment
}
