package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserForumThreadsResponse struct {
	ForumThreads []struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		ForumID int    `json:"forum_id"`
		Posts   int    `json:"posts"`
		Rating  int    `json:"rating"`
		Views   int    `json:"views"`
		Author  struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Karma    int    `json:"karma"`
		} `json:"author"`
		LastPoster struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Karma    int    `json:"karma"`
		} `json:"last_poster"`
		FirstPostTime int  `json:"first_post_time"`
		LastPostTime  int  `json:"last_post_time"`
		HasPoll       bool `json:"has_poll"`
		IsLocked      bool `json:"is_locked"`
		IsSticky      bool `json:"is_sticky"`
		NewPosts      int  `json:"new_posts"`
	} `json:"forumThreads"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserForumThreadsSort string

const (
	ForumThreadsAscending  UserForumThreadsSort = "ASC"
	ForumThreadsDescending UserForumThreadsSort = "DESC"
)

type UserForumThreadsParams struct {
	api.BaseParams
	Limit int32
	api.RangeParams
	Sort UserForumThreadsSort
}

func CreateForumThreadsURL(t *UserForumThreadsParams) (string, error) {
	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "forumthreads")

	q := u.Query()
	q.Set("key", t.APIKey)
	if t.Limit != 0 {
		q.Set("limit", fmt.Sprintf("%d", t.Limit))
	}

	if err := api.SetRangeParams(q, t.From, t.To); err != nil {
		return "", err
	}

	if t.Sort != "" {
		q.Set("sort", string(t.Sort))
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
