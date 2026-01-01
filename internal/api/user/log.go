package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserLogResponse struct {
	Log []struct {
		ID        string `json:"id"`
		Timestamp int    `json:"timestamp"`
		Details   struct {
			ID       int    `json:"id"`
			Title    string `json:"title"`
			Category string `json:"category"`
		} `json:"details"`
		Data struct {
		} `json:"data"`
		Params struct {
		} `json:"params"`
	} `json:"log"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserLogParams struct {
	Log       []int
	Cat       int
	Target    int
	Limit     api.ApiLimit
	From      api.ApiFrom
	To        api.ApiTo
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserLogParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "log", fmt.Sprintf("%v", p.Log))
	api.SetIfNotZero(q, "cat", p.Cat)
	api.SetIfNotZero(q, "target", p.Target)
	api.SetIfNotZero(q, "limit", p.Limit)
	api.SetIfNotZero(q, "from", p.From)
	api.SetIfNotZero(q, "to", p.To)
	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetLog(session *api.Session, params *UserLogParams) (*UserLogResponse, error) {
	var resp UserLogResponse
	if err := session.Get("user/log", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
