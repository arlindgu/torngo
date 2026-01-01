package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserNewEventsResponse struct {
	Events []struct {
		ID        string `json:"id"`
		Timestamp int    `json:"timestamp"`
		Event     string `json:"event"`
	} `json:"events"`
}

type UserNewEventsParams struct {
	Striptags api.ApiStriptags
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserNewEventsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "striptags", fmt.Sprintf("%v", p.Striptags))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetNewEvents(session *api.Session, params *UserNewEventsParams) (*UserNewEventsResponse, error) {
	var resp UserNewEventsResponse
	if err := session.Get("user/newevents", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
