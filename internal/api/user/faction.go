package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserFactionResponse struct {
	Faction struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Tag           string `json:"tag"`
		TagImage      string `json:"tag_image"`
		Position      string `json:"position"`
		DaysInFaction int    `json:"days_in_faction"`
	} `json:"faction"`
}

type UserFactionParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserFactionParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetFaction(session *api.Session, params *UserFactionParams) (*UserFactionResponse, error) {
	var resp UserFactionResponse
	if err := session.Get("user/faction", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
