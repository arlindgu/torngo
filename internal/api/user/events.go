package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserEventsResponse struct {
	Events []struct {
		ID        string `json:"id"`
		Timestamp int    `json:"timestamp"`
		Event     string `json:"event"`
	} `json:"events"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserEventsParams struct {
	Striptags bool
	Limit     int32
	api.RangeParams
	api.BaseParams
}

func CreateEventsURL(t *UserEventsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "battlestats")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Striptags {
		q.Set("striptags", fmt.Sprintf("%t", t.Striptags))
	}

	if t.Limit != 0 {
		q.Set("limit", fmt.Sprintf("%d", t.Limit))
	} else if t.Limit == 0 {
		q.Set("limit", "20")
	}

	if t.From > t.To {
		return "", fmt.Errorf("From parameter can't be greater than To parameter")
	} else {
		if t.To != 0 {
			q.Set("to", fmt.Sprintf("%d", t.To))
		}
		if t.From != 0 {
			q.Set("from", fmt.Sprintf("%d", t.From))
		}
	}

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
