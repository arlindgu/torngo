package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserMedalsResponse struct {
	Medals []struct {
		ID        int `json:"id"`
		Timestamp int `json:"timestamp"`
	} `json:"medals"`
}

type UserMedalsParams struct {
	api.BaseParams
}

func CreateMedalsURL(t *UserMedalsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "medals")

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
