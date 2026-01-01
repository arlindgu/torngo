package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserBasicResponse struct {
	Profile struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Level  int    `json:"level"`
		Gender string `json:"gender"`
		Status struct {
			Description    string `json:"description"`
			Details        string `json:"details"`
			State          string `json:"state"`
			Color          string `json:"color"`
			Until          int    `json:"until"`
			PlaneImageType string `json:"plane_image_type"`
		} `json:"status"`
	} `json:"profile"`
}

type UserBasicParams struct {
	Striptags api.ApiStriptags
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserBasicParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "striptags", fmt.Sprintf("%t", p.Striptags))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	return q
}

func GetBasic(session *api.Session, params *UserBasicParams) (*UserBasicResponse, error) {
	var resp UserBasicResponse
	if err := session.Get("user/basic", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UserIdBasicResponse struct {
	Profile struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Level  int    `json:"level"`
		Gender string `json:"gender"`
		Status struct {
			Description    string `json:"description"`
			Details        string `json:"details"`
			State          string `json:"state"`
			Color          string `json:"color"`
			Until          int    `json:"until"`
			PlaneImageType string `json:"plane_image_type"`
		} `json:"status"`
	} `json:"profile"`
}

type UserBasicIdParams struct {
	Id        api.ApiId
	DiscordId api.ApiDiscordId
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserBasicIdParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetBasicId(session *api.Session, params *UserBasicIdParams) (*UserIdBasicResponse, error) {
	var resp UserIdBasicResponse
	if err := session.Get(fmt.Sprintf("user/%d/basic", params.Id), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
