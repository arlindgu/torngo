package user

import (
	"fmt"
	"net/url"
	"path"
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
	Striptags bool
	Limit     int32
	api.RangeParams
	api.BaseParams
}

func CreateForumPostsURL(t *UserForumPostsParams) (string, error) {
	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "battlestats")

	q := u.Query()
	if t.Striptags {
		q.Set("striptags", fmt.Sprintf("%t", t.Striptags))
	}

	api.SetLimitParams(q, int(t.Limit), 20)

	if err := api.SetRangeParams(q, t.From, t.To); err != nil {
		return "", err
	}

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}

	q.Set("key", t.APIKey)

	u.RawQuery = q.Encode()
	return u.String(), nil
}
