package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserRefillsResponse struct {
	Refills struct {
		Energy       bool `json:"energy"`
		Nerve        bool `json:"nerve"`
		Token        bool `json:"token"`
		SpecialCount int  `json:"special_count"`
	} `json:"refills"`
}

type UserRefillsParams struct {
	api.BaseParams
}

func CreateRefillsURL(t *UserRefillsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "refills")

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
