package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserSkillsResponse struct {
	Skills []struct {
		Slug  string  `json:"slug"`
		Name  string  `json:"name"`
		Level float64 `json:"level"`
	} `json:"skills"`
}

type UserSkillsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserSkillsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetSkills(session *api.Session, params *UserSkillsParams) (*UserSkillsResponse, error) {
	var resp UserSkillsResponse
	if err := session.Get("user/skills", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
