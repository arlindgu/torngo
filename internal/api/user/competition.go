package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserCompetitionResponse struct {
	Competition struct {
		Name            string `json:"name"`
		TreatsCollected int    `json:"treats_collected"`
		Basket          struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"basket"`
	} `json:"competition"`
}

type UserCompetitionParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserCompetitionParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetCompetition(session *api.Session, params *UserCompetitionParams) (*UserCompetitionResponse, error) {
	var resp UserCompetitionResponse
	if err := session.Get("user/competition", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
