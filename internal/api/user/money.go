package user

import (
	"fmt"
	"net/url"
	"path"
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
	api.BaseParams
}

func CreateMoneyURL(t *UserMoneyParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "money")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
