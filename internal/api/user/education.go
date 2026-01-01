package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserEducationResponse struct {
	Education struct {
		Complete []int `json:"complete"`
		Current  struct {
			ID    int `json:"id"`
			Until int `json:"until"`
		} `json:"current"`
	} `json:"education"`
}

type UserEducationParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserEducationParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetEducation(session *api.Session, params *UserEducationParams) (*UserEducationResponse, error) {
	var resp UserEducationResponse
	if err := session.Get("user/education", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
