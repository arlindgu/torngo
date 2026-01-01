package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserJobResponse struct {
	Job struct {
		Type     string `json:"type"`
		Name     string `json:"name"`
		Position string `json:"position"`
	} `json:"job"`
}

type UserJobParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserJobParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetJob(session *api.Session, params *UserJobParams) (*UserJobResponse, error) {
	var resp UserJobResponse
	if err := session.Get("user/job", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
