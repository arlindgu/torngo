package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserNewEventsResponse struct {
	Events []struct {
		ID        string `json:"id"`
		Timestamp int    `json:"timestamp"`
		Event     string `json:"event"`
	} `json:"events"`
}

type UserNewEventsParams struct {
	Striptags bool
	api.BaseParams
}

func CreateNewEventsURL(t *UserNewEventsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "newevents")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}
	if t.Striptags {
		q.Set("striptags", fmt.Sprintf("%t", t.Striptags))
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
