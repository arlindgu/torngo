package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserJobResponse struct {
	Job struct {
		Type     string `json:"type"`
		Name     string `json:"name"`
		Position string `json:"position"`
	} `json:"job"`
}

type UserJobParams struct {
	api.BaseParams
}

func CreateJobURL(t *UserJobParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "job")

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
