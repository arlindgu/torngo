package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserEventsResponse struct {
	Events []struct {
		ID        string `json:"id"`
		Timestamp int    `json:"timestamp"`
		Event     string `json:"event"`
	} `json:"events"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserEventsParams struct {
	Striptags api.ApiStriptags
	Limit     api.ApiLimit
	From      api.ApiFrom
	To        api.ApiTo
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserEventsParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "striptags", fmt.Sprintf("%v", p.Striptags))
	api.SetIfNotZero(q, "limit", p.Limit)
	api.SetIfNotZero(q, "from", p.From)
	api.SetIfNotZero(q, "to", p.To)
	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetEvents(session *api.Session, params *UserEventsParams) (*UserEventsResponse, error) {
	var resp UserEventsResponse
	if err := session.Get("user/events", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
