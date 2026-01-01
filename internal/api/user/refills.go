package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserRefillsResponse struct {
	Refills struct {
		Energy       bool `json:"energy"`
		Nerve        bool `json:"nerve"`
		Token        bool `json:"token"`
		SpecialCount int  `json:"special_count"`
	} `json:"refills"`
}

type UserRefillsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserRefillsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetRefills(session *api.Session, params *UserRefillsParams) (*UserRefillsResponse, error) {
	var resp UserRefillsResponse
	if err := session.Get("user/refills", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
