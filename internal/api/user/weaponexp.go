package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserWeaponExpResponse struct {
	Weaponexp []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Exp  int    `json:"exp"`
	} `json:"weaponexp"`
}

type UserWeaponExpParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserWeaponExpParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetWeaponExp(session *api.Session, params *UserWeaponExpParams) (*UserWeaponExpResponse, error) {
	var resp UserWeaponExpResponse
	if err := session.Get("user/weaponexp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
