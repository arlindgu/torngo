package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserPropertyResponse struct {
	Property struct {
		ID    int64 `json:"id"`
		Owner struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"owner"`
		Property struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"property"`
		Happy  int `json:"happy"`
		Upkeep struct {
			Property int `json:"property"`
			Staff    int `json:"staff"`
		} `json:"upkeep"`
		MarketPrice   int64    `json:"market_price"`
		Modifications []string `json:"modifications"`
		Staff         []struct {
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"staff"`
		UsedBy []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"used_by"`
	} `json:"property"`
}

type UserPropertyParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserPropertyParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetProperty(session *api.Session, params *UserPropertyParams) (*UserPropertyResponse, error) {
	var resp UserPropertyResponse
	if err := session.Get("user/property", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
