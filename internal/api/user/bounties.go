package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserBountiesResponse struct {
	Bounties []struct {
		TargetID    int    `json:"target_id"`
		TargetName  string `json:"target_name"`
		TargetLevel int    `json:"target_level"`
		ListerID    int    `json:"lister_id"`
		ListerName  string `json:"lister_name"`
		Reward      int64  `json:"reward"`
		Reason      string `json:"reason"`
		Quantity    int    `json:"quantity"`
		IsAnonymous bool   `json:"is_anonymous"`
		ValidUntil  int    `json:"valid_until"`
	} `json:"bounties"`
}

type UserBountiesParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserBountiesParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetBounties(session *api.Session, params *UserBountiesParams) (*UserBountiesResponse, error) {
	var resp UserBountiesResponse
	if err := session.Get("user/bounties", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
