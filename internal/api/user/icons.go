package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserIconsResponse struct {
	Icons []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Until       int    `json:"until"`
	} `json:"icons"`
}

type UserIconsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserIconsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetIcons(session *api.Session, params *UserIconsParams) (*UserIconsResponse, error) {
	var resp UserIconsResponse
	if err := session.Get("user/icons", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
