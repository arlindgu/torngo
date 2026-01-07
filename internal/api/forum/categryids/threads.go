package categryids

import "torngo/internal/api"

type ForumCategoryIdsThreadsResponse struct {
	Threads []struct {
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
	} `json:"threads"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type ForumCategoryIdsThreadsParams struct {
	CategoryIds []int32
	Limit       api.ApiLimit
	Sort        api.ApiSort
	From        api.ApiFrom
	To          api.ApiTo
	Timestamp   api.ApiTimestamp
	Comment     api.ApiComment
}
