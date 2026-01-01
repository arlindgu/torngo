package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserMessagesResponse struct {
	Messages []struct {
		ID     int64 `json:"id"`
		Sender struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"sender"`
		Timestamp int    `json:"timestamp"`
		Topic     string `json:"topic"`
		Type      string `json:"type"`
		Seen      bool   `json:"seen"`
		Read      bool   `json:"read"`
	} `json:"messages"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserMessagesParams struct {
	Limit     api.ApiLimit
	From      api.ApiFrom
	To        api.ApiTo
	Sort      api.ApiSort
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserMessagesParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "limit", p.Limit)
	api.SetIfNotZero(q, "from", p.From)
	api.SetIfNotZero(q, "to", p.To)
	api.SetIfNotZero(q, "sort", string(p.Sort))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetMessages(session *api.Session, params *UserMessagesParams) (*UserMessagesResponse, error) {
	var resp UserMessagesResponse
	if err := session.Get("user/messages", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
