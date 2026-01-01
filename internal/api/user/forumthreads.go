package user

import (
	"fmt"
	"net/url"
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
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
	Limit     api.ApiLimit
	From      api.ApiFrom
	To        api.ApiTo
	Sort      api.ApiSort
}

func (p *UserForumThreadsParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "limit", p.Limit)
	api.SetIfNotZero(q, "from", p.From)
	api.SetIfNotZero(q, "to", p.To)
	api.SetIfNotZero(q, "sort", string(p.Sort))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetForumThreads(session *api.Session, params *UserForumThreadsParams) (*UserForumThreadsResponse, error) {
	var resp UserForumThreadsResponse
	if err := session.Get("user/forumthreads", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
