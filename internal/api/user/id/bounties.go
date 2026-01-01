package id

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserIdBountiesResponse struct {
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
	Id        api.ApiId
	DiscordId api.ApiDiscordId
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserBountiesParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetBounties(session *api.Session, params *UserBountiesParams) (*UserIdBountiesResponse, error) {
	var resp UserIdBountiesResponse

	var paramsid = ""

	if params.Id != 0 {
		paramsid = fmt.Sprintf("%d", params.Id)
	} else if params.DiscordId != 0 {
		paramsid = fmt.Sprintf("discord/%d", params.DiscordId)
	}

	if err := session.Get(fmt.Sprintf("user/%s/bounties", paramsid), params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
