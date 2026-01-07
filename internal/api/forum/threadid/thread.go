package threadid

import "torngo/internal/api"

type ForumThreadIdThreadResponse struct {
	Thread struct {
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
		FirstPostTime int    `json:"first_post_time"`
		LastPostTime  int    `json:"last_post_time"`
		HasPoll       bool   `json:"has_poll"`
		IsLocked      bool   `json:"is_locked"`
		IsSticky      bool   `json:"is_sticky"`
		Content       string `json:"content"`
		ContentRaw    string `json:"content_raw"`
		Poll          struct {
			Question     string `json:"question"`
			AnswersCount int    `json:"answers_count"`
			Answers      []struct {
				Answer string `json:"answer"`
				Votes  int    `json:"votes"`
			} `json:"answers"`
		} `json:"poll"`
	} `json:"thread"`
}

type ThreadIdThreadParams struct {
	ThreadId  int32
	Timestamp api.ApiTimestamp
	Comment   api.ApiComment
}
