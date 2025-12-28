package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserVirusResponse struct {
	Virus struct {
		Item struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"item"`
		Until int `json:"until"`
	} `json:"virus"`
}

type UserVirusParams struct {
	api.BaseParams
}

func CreateVirusURL(t *UserVirusParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "virus")

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
