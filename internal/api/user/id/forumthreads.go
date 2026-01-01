package id

type UserIdForumThreadsResponse struct {
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
