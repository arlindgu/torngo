package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserNewMessagesResponse struct {
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
}

type UserNewMessagesParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserNewMessagesParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetNewMessages(session *api.Session, params *UserNewMessagesParams) (*UserNewMessagesResponse, error) {
	var resp UserNewMessagesResponse
	if err := session.Get("user/newmessages", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
