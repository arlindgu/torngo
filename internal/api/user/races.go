package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserRacesResponse struct {
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
		Results []struct {
			DriverID    int     `json:"driver_id"`
			Position    int     `json:"position"`
			CarID       int     `json:"car_id"`
			CarItemID   int64   `json:"car_item_id"`
			CarItemName string  `json:"car_item_name"`
			CarClass    string  `json:"car_class"`
			HasCrashed  bool    `json:"has_crashed"`
			BestLapTime float64 `json:"best_lap_time"`
			RaceTime    float64 `json:"race_time"`
			TimeEnded   int     `json:"time_ended"`
		} `json:"results"`
		IsOfficial bool    `json:"is_official"`
		SkillGain  float64 `json:"skill_gain"`
	} `json:"races"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserRacesCategory string

const (
	RacesCategoryOfficial UserRacesCategory = "official"
	RacesCategoryCustom   UserRacesCategory = "custom"
)

type UserRacesParams struct {
	Limit     api.ApiLimit
	Sort      api.ApiSort
	From      api.ApiFrom
	To        api.ApiTo
	Category  UserRacesCategory
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserRacesParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "limit", int(p.Limit))
	api.SetIfNotZero(q, "sort", string(p.Sort))
	api.SetIfNotZero(q, "from", int(p.From))
	api.SetIfNotZero(q, "to", int(p.To))
	api.SetIfNotZero(q, "category", string(p.Category))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetRaces(session *api.Session, params *UserRacesParams) (*UserRacesResponse, error) {
	var resp UserRacesResponse
	if err := session.Get("user/races", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
