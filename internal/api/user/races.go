package user

import (
	"fmt"
	"net/url"
	"path"
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
		IsOfficial bool `json:"is_official"`
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
	Limit int
	Sort  api.SortParams
	api.RangeParams
	Category UserRacesCategory
	api.BaseParams
}

func CreateRacesURL(t *UserRacesParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "races")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Category != "" {
		q.Set("category", string(t.Category))
	}

	api.SetLimitParams(q, t.Limit, 20)

	if err := api.SetRangeParams(q, t.From, t.To); err != nil {
		return "", err
	}

	if t.Sort != "" {
		q.Set("sort", string(t.Sort))
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
