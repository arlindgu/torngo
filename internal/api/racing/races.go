package racing

import "torngo/internal/api"

type RacingRacesResponse struct {
	Races []struct {
		ID           int    `json:"id"`
		Title        string `json:"title"`
		TrackID      int    `json:"track_id"`
		CreatorID    int    `json:"creator_id"`
		Status       string `json:"status"`
		Laps         int    `json:"laps"`
		Participants struct {
			Minimum int `json:"minimum"`
			Maximum int `json:"maximum"`
			Current int `json:"current"`
		} `json:"participants"`
		Schedule struct {
			JoinFrom  int `json:"join_from"`
			JoinUntil int `json:"join_until"`
			Start     int `json:"start"`
			End       int `json:"end"`
		} `json:"schedule"`
		Requirements struct {
			CarClass         string `json:"car_class"`
			DriverClass      string `json:"driver_class"`
			CarItemID        int64  `json:"car_item_id"`
			RequiresStockCar bool   `json:"requires_stock_car"`
			RequiresPassword bool   `json:"requires_password"`
			JoinFee          int    `json:"join_fee"`
		} `json:"requirements"`
		IsOfficial bool `json:"is_official"`
	} `json:"races"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type RacesCategory string

const (
	RacesCategoryOfficial RacesCategory = "official"
	RacesCategoryCustom   RacesCategory = "custom"
)

type RacingRacesParams struct {
	Limit     api.ApiLimit
	Sort      api.ApiSort
	To        api.ApiTo
	From      api.ApiFrom
	Category  RacesCategory
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
