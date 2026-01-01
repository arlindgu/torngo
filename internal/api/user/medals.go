package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserMedalsResponse struct {
	Medals []struct {
		ID        int `json:"id"`
		Timestamp int `json:"timestamp"`
	} `json:"medals"`
}

type UserMedalsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserMedalsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetMedals(session *api.Session, params *UserMedalsParams) (*UserMedalsResponse, error) {
	var resp UserMedalsResponse
	if err := session.Get("user/medals", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
