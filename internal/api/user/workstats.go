package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserWorkstatsResponse struct {
	Workstats struct {
		Endurance    int `json:"endurance"`
		Intelligence int `json:"intelligence"`
		ManualLabor  int `json:"manual_labor"`
		Total        int `json:"total"`
	} `json:"workstats"`
}

type UserWorkstatsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserWorkstatsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetWorkstats(session *api.Session, params *UserWorkstatsParams) (*UserWorkstatsResponse, error) {
	var resp UserWorkstatsResponse
	if err := session.Get("user/workstats", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
