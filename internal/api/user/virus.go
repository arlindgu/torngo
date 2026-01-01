package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserVirusResponse struct {
	Virus struct {
		Item struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"item"`
		Until int `json:"until"`
	} `json:"virus"`
}

type UserVirusParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserVirusParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetVirus(session *api.Session, params *UserVirusParams) (*UserVirusResponse, error) {
	var resp UserVirusResponse
	if err := session.Get("user/virus", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
