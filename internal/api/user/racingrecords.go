package user

import (
	"fmt"
	"net/url"
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
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserRacingRecordsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetRacingRecords(session *api.Session, params *UserRacingRecordsParams) (*UserRacingRecordsResponse, error) {
	var resp UserRacingRecordsResponse
	if err := session.Get("user/racingrecords", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
