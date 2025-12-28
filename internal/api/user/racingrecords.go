package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserRacingRecordsResponse struct {
	Racingrecords []struct {
		Track struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"track"`
		Records []struct {
			CarID   int64  `json:"car_id"`
			CarName string `json:"car_name"`
			LapTime int    `json:"lap_time"`
		} `json:"records"`
	} `json:"racingrecords"`
}

type UserRacingRecordsParams struct {
	api.BaseParams
}

func CreateRacingRecordsURL(t *UserRacingRecordsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "racingrecords")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
