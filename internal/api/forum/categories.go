package forum

import "torngo/internal/api"

type ForumCategoriesResponse struct {
	Categories []struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Acronym string `json:"acronym"`
		Threads int    `json:"threads"`
	} `json:"categories"`
}

type ForumCategoriesParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
