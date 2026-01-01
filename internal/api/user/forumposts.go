package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserForumPostsResponse struct {
	ForumPosts []struct {
		ID       int `json:"id"`
		ThreadID int `json:"thread_id"`
		Author   struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Karma    int    `json:"karma"`
		} `json:"author"`
		IsLegacy     bool   `json:"is_legacy"`
		IsTopic      bool   `json:"is_topic"`
		IsEdited     bool   `json:"is_edited"`
		IsPinned     bool   `json:"is_pinned"`
		CreatedTime  int    `json:"created_time"`
		EditedBy     int    `json:"edited_by"`
		HasQuote     bool   `json:"has_quote"`
		QuotedPostID int    `json:"quoted_post_id"`
		Content      string `json:"content"`
		Likes        int    `json:"likes"`
		Dislikes     int    `json:"dislikes"`
	} `json:"forumPosts"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserForumPostsParams struct {
	Striptags api.ApiStriptags
	Limit     api.ApiLimit
	From      api.ApiFrom
	To        api.ApiTo
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserForumPostsParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "striptags", fmt.Sprintf("%v", p.Striptags))
	api.SetIfNotZero(q, "limit", p.Limit)
	api.SetIfNotZero(q, "from", p.From)
	api.SetIfNotZero(q, "to", p.To)
	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetForumPosts(session *api.Session, params *UserForumPostsParams) (*UserForumPostsResponse, error) {
	var resp UserForumPostsResponse
	if err := session.Get("user/forumposts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
