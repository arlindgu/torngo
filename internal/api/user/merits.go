package user

import (
	"fmt"
	"net/url"
	"path"
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
	api.BaseParams
}

func CreateMeritsURL(t *UserMeritsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "merits")

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
