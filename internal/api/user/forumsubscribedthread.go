package user

import (
	"fmt"
	"net/url"
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
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserForumSubscribedThreadsParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetForumSubscribedThreads(session *api.Session, params *UserForumSubscribedThreadsParams) (*UserForumSubscribedThreadsResponse, error) {
	var resp UserForumSubscribedThreadsResponse
	if err := session.Get("user/battlestats", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
