package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserTimestampResponse struct {
	Timestamp int `json:"timestamp"`
}

type UserTimestampParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserTimestampParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetTimestamp(session *api.Session, params *UserTimestampParams) (*UserTimestampResponse, error) {
	var resp UserTimestampResponse
	if err := session.Get("user/timestamp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
