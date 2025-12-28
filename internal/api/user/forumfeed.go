package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserForumFeedResponse struct {
	ForumFeed []struct {
		ThreadID int `json:"thread_id"`
		PostID   int `json:"post_id"`
		User     struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Karma    int    `json:"karma"`
		} `json:"user"`
		Title     string `json:"title"`
		Text      string `json:"text"`
		Timestamp int    `json:"timestamp"`
		IsSeen    bool   `json:"is_seen"`
		Type      int    `json:"type"`
	} `json:"forumFeed"`
}

type UserForumFeedParams struct {
	api.BaseParams
}

func CreateForumFeedURL(t *UserForumFeedParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "ammo")

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
