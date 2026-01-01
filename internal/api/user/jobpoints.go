package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserJobpointsResponse struct {
	Jobpoints struct {
		Jobs struct {
			Army      int `json:"army"`
			Casino    int `json:"casino"`
			Education int `json:"education"`
			Grocer    int `json:"grocer"`
			Law       int `json:"law"`
			Medical   int `json:"medical"`
		} `json:"jobs"`
		Companies []struct {
			Company struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"company"`
			Points int `json:"points"`
		} `json:"companies"`
	} `json:"jobpoints"`
}

type UserJobpointsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserJobpointsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetJobpoints(session *api.Session, params *UserJobpointsParams) (*UserJobpointsResponse, error) {
	var resp UserJobpointsResponse
	if err := session.Get("user/jobpoints", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
