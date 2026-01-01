package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserEnlistedCarsResponse struct {
	Enlistedcars []struct {
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
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Worth        int64  `json:"worth"`
		PointsSpent  int    `json:"points_spent"`
		RacesEntered int    `json:"races_entered"`
		RacesWon     int    `json:"races_won"`
		IsRemoved    bool   `json:"is_removed"`
		Parts        []int  `json:"parts"`
	} `json:"enlistedcars"`
}

type UserEnlistedCarsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserEnlistedCarsParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetEnlistedCars(session *api.Session, params *UserEnlistedCarsParams) (*UserEnlistedCarsResponse, error) {
	var resp UserEnlistedCarsResponse
	if err := session.Get("user/enlistedcars", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
