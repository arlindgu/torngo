package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserNewMessagesResponse struct {
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
}

type UserNewMessagesParams struct {
	api.BaseParams
}

func CreateNewMessagesURL(t *UserNewMessagesParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "newmessages")

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
