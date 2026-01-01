package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserBattlestatsResponse struct {
	Battlestats struct {
		Strength struct {
			Value     int64 `json:"value"`
			Modifier  int   `json:"modifier"`
			Modifiers []struct {
				Effect string `json:"effect"`
				Value  int    `json:"value"`
				Type   string `json:"type"`
			} `json:"modifiers"`
		} `json:"strength"`
		Defense struct {
			Value     int64 `json:"value"`
			Modifier  int   `json:"modifier"`
			Modifiers []struct {
				Effect string `json:"effect"`
				Value  int    `json:"value"`
				Type   string `json:"type"`
			} `json:"modifiers"`
		} `json:"defense"`
		Speed struct {
			Value     int64 `json:"value"`
			Modifier  int   `json:"modifier"`
			Modifiers []struct {
				Effect string `json:"effect"`
				Value  int    `json:"value"`
				Type   string `json:"type"`
			} `json:"modifiers"`
		} `json:"speed"`
		Dexterity struct {
			Value     int64 `json:"value"`
			Modifier  int   `json:"modifier"`
			Modifiers []struct {
				Effect string `json:"effect"`
				Value  int    `json:"value"`
				Type   string `json:"type"`
			} `json:"modifiers"`
		} `json:"dexterity"`
		Total int64 `json:"total"`
	} `json:"battlestats"`
}

type UserBattlestatsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserBattlestatsParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetBattlestats(session *api.Session, params *UserBattlestatsParams) (*UserBattlestatsResponse, error) {
	var resp UserBattlestatsResponse
	if err := session.Get("user/battlestats", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
