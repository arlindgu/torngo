package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserTravelResponse struct {
	Travel struct {
		Destination string `json:"destination"`
		Method      string `json:"method"`
		DepartedAt  int    `json:"departed_at"`
		ArrivalAt   int    `json:"arrival_at"`
		TimeLeft    int    `json:"time_left"`
	} `json:"travel"`
}

type UserTravelParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserTravelParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetTravel(session *api.Session, params *UserTravelParams) (*UserTravelResponse, error) {
	var resp UserTravelResponse
	if err := session.Get("user/travel", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
