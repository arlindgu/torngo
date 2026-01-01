package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserJobranksResponse struct {
	Jobranks struct {
		Army      string `json:"army"`
		Grocer    string `json:"grocer"`
		Casino    string `json:"casino"`
		Medical   string `json:"medical"`
		Law       string `json:"law"`
		Education string `json:"education"`
	} `json:"jobranks"`
}

type UserJobranksParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserJobranksParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetJobranks(session *api.Session, params *UserJobranksParams) (*UserJobranksResponse, error) {
	var resp UserJobranksResponse
	if err := session.Get("user/jobranks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
