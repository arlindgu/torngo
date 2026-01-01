package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserMeritsResponse struct {
	Merits struct {
		Upgrades []struct {
			ID    int `json:"id"`
			Level int `json:"level"`
		} `json:"upgrades"`
		Available int `json:"available"`
		Used      int `json:"used"`
		Medals    int `json:"medals"`
		Honors    int `json:"honors"`
	} `json:"merits"`
}

type UserMeritsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserMeritsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetMerits(session *api.Session, params *UserMeritsParams) (*UserMeritsResponse, error) {
	var resp UserMeritsResponse
	if err := session.Get("user/merits", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
