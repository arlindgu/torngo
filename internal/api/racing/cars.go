package racing

import "torngo/internal/api"

type RacingCarsResponse struct {
	Cars []struct {
		CarItemID    int64  `json:"car_item_id"`
		CarItemName  string `json:"car_item_name"`
		TopSpeed     int    `json:"top_speed"`
		Acceleration int    `json:"acceleration"`
		Braking      int    `json:"braking"`
		Dirt         int    `json:"dirt"`
		Handling     int    `json:"handling"`
		Safety       int    `json:"safety"`
		Tarmac       int    `json:"tarmac"`
		Class        string `json:"class"`
	} `json:"cars"`
}

type RacingCarsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
