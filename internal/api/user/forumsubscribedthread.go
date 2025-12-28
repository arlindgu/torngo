package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserForumSubscribedThreadsResponse struct {
	ForumSubscribedThreads []struct {
		ID      int `json:"id"`
		ForumID int `json:"forum_id"`
		Author  struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Karma    int    `json:"karma"`
		} `json:"author"`
		Title string `json:"title"`
		Posts struct {
			New   int `json:"new"`
			Total int `json:"total"`
		} `json:"posts"`
	} `json:"forumSubscribedThreads"`
}

type UserForumSubscribedThreadsParams struct {
	api.BaseParams
}

func CreateForumSubscribedThreadsURL(t *UserForumSubscribedThreadsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "forumsubscribedthreads")

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
