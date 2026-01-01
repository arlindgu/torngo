package id

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

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

func GetBasic(session *api.Session, params *UserBasicIdParams) (*UserIdBasicResponse, error) {
	var resp UserIdBasicResponse

	var paramsid = ""

	if params.Id != 0 {
		paramsid = fmt.Sprintf("%d", params.Id)
	} else if params.DiscordId != 0 {
		paramsid = fmt.Sprintf("discord/%d", params.DiscordId)
	}

	if err := session.Get(fmt.Sprintf("user/%s/basic", paramsid), params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
