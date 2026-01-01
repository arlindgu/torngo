package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserCooldownsResponse struct {
	Cooldowns struct {
		Drug    int `json:"drug"`
		Medical int `json:"medical"`
		Booster int `json:"booster"`
	} `json:"cooldowns"`
}

type UserCooldownsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserCooldownsParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetCooldowns(session *api.Session, params *UserCooldownsParams) (*UserCooldownsResponse, error) {
	var resp UserCooldownsResponse
	if err := session.Get("user/cooldowns", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
