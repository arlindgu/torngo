package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserHonorsResponse struct {
	Honors []struct {
		ID        int `json:"id"`
		Timestamp int `json:"timestamp"`
	} `json:"honors"`
}

type UserHonorsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserHonorsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetHonors(session *api.Session, params *UserHonorsParams) (*UserHonorsResponse, error) {
	var resp UserHonorsResponse
	if err := session.Get("user/honors", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
