package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserMessagesResponse struct {
	Messages []struct {
		ID     int64 `json:"id"`
		Sender struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"sender"`
		Timestamp int    `json:"timestamp"`
		Topic     string `json:"topic"`
		Type      string `json:"type"`
		Seen      bool   `json:"seen"`
		Read      bool   `json:"read"`
	} `json:"messages"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserMessagesParams struct {
	Limit int
	api.RangeParams
	Sort api.SortParams
	api.BaseParams
}

func CreateMessagesURL(t *UserMessagesParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}
	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "messages")

	q := u.Query()
	q.Set("key", t.APIKey)

	api.SetLimitParams(q, t.Limit, 20)

	if err := api.SetRangeParams(q, t.From, t.To); err != nil {
		return "", err
	}

	if t.Sort != "" {
		q.Set("sort", string(t.Sort))
	}

	return u.String(), nil
}
