package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserHofResponse struct {
	Hof struct {
		Attacks struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"attacks"`
		Busts struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"busts"`
		Defends struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"defends"`
		Networth struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"networth"`
		Offences struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"offences"`
		Revives struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"revives"`
		Level struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"level"`
		Rank struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"rank"`
		Awards struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"awards"`
		RacingSkill struct {
			Value float64 `json:"value"`
			Rank  int     `json:"rank"`
		} `json:"racing_skill"`
		RacingPoints struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"racing_points"`
		RacingWins struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"racing_wins"`
		TravelTime struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"travel_time"`
		WorkingStats struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"working_stats"`
		BattleStats struct {
			Value int64 `json:"value"`
			Rank  int   `json:"rank"`
		} `json:"battle_stats"`
	} `json:"hof"`
}

type UserHofParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserHofParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetHof(session *api.Session, params *UserHofParams) (*UserHofResponse, error) {
	var resp UserHofResponse
	if err := session.Get("user/hof", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
