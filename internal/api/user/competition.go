package user

import (
	"fmt"
	"net/url"
	"path"
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
	api.BaseParams
}

func CreateCompetitionURL(t *UserCompetitionParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "competition")

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
