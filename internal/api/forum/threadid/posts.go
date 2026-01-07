package threadid

import "torngo/internal/api"

type ForumThreadIdPostsResponse struct {
	Posts []struct {
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
	} `json:"posts"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type ThreadIdPostsParams struct {
	ThreadId  int32
	Striptags api.ApiStriptags
	Offset    api.ApiOffset
	Sort      api.ApiSort
	From      api.ApiFrom
	To        api.ApiTo
	Timestamp api.ApiTimestamp
	Comment   api.ApiComment
}
