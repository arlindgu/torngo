package user

import (
	"fmt"
	"net/url"
	"path"
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
	api.BaseParams
}

func CreatePropertyURL(t *UserPropertyParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "property")
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
