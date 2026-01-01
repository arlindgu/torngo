package user

import (
	"fmt"
	"net/url"
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
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserForumFeedParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetForumFeed(session *api.Session, params *UserForumFeedParams) (*UserForumFeedResponse, error) {
	var resp UserForumFeedResponse
	if err := session.Get("user/forumfeed", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
