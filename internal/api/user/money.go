package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserMoneyResponse struct {
	Money struct {
		Points     int64 `json:"points"`
		Wallet     int64 `json:"wallet"`
		Company    int64 `json:"company"`
		Vault      int64 `json:"vault"`
		CaymanBank int64 `json:"cayman_bank"`
		CityBank   struct {
			Amount       int64   `json:"amount"`
			Profit       int64   `json:"profit"`
			Duration     int     `json:"duration"`
			InterestRate float64 `json:"interest_rate"`
			Until        int64   `json:"until"`
			InvestedAt   int     `json:"invested_at"`
		} `json:"city_bank"`
		Faction struct {
			Money  int64 `json:"money"`
			Points int64 `json:"points"`
		} `json:"faction"`
		DailyNetworth int64 `json:"daily_networth"`
	} `json:"money"`
}

type UserMoneyParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserMoneyParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetMoney(session *api.Session, params *UserMoneyParams) (*UserMoneyResponse, error) {
	var resp UserMoneyResponse
	if err := session.Get("user/money", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
