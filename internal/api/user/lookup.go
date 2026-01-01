package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserLookupResponse struct {
	Selections []string `json:"selections"`
}

type UserLookupParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserLookupParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetLookup(session *api.Session, params *UserLookupParams) (*UserLookupResponse, error) {
	var resp UserLookupResponse
	if err := session.Get("user/lookup", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
