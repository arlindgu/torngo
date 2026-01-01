package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserAmmoResponse struct {
	Ammo []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Types []struct {
			Name     string `json:"name"`
			Quantity int    `json:"quantity"`
			Equipped bool   `json:"equipped"`
		} `json:"types"`
	} `json:"ammo"`
}

type UserAmmoParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserAmmoParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetAmmo(session *api.Session, params *UserAmmoParams) (*UserAmmoResponse, error) {
	var resp UserAmmoResponse
	if err := session.Get("user/ammo", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
