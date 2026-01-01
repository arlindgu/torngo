package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserBarsResponse struct {
	Bars struct {
		Energy struct {
			Current   int `json:"current"`
			Maximum   int `json:"maximum"`
			Increment int `json:"increment"`
			Interval  int `json:"interval"`
			TickTime  int `json:"tick_time"`
			FullTime  int `json:"full_time"`
		} `json:"energy"`
		Nerve struct {
			Current   int `json:"current"`
			Maximum   int `json:"maximum"`
			Increment int `json:"increment"`
			Interval  int `json:"interval"`
			TickTime  int `json:"tick_time"`
			FullTime  int `json:"full_time"`
		} `json:"nerve"`
		Happy struct {
			Current   int `json:"current"`
			Maximum   int `json:"maximum"`
			Increment int `json:"increment"`
			Interval  int `json:"interval"`
			TickTime  int `json:"tick_time"`
			FullTime  int `json:"full_time"`
		} `json:"happy"`
		Life struct {
			Current   int `json:"current"`
			Maximum   int `json:"maximum"`
			Increment int `json:"increment"`
			Interval  int `json:"interval"`
			TickTime  int `json:"tick_time"`
			FullTime  int `json:"full_time"`
		} `json:"life"`
		Chain struct {
			ID       int `json:"id"`
			Current  int `json:"current"`
			Max      int `json:"max"`
			Timeout  int `json:"timeout"`
			Modifier int `json:"modifier"`
			Cooldown int `json:"cooldown"`
			Start    int `json:"start"`
			End      int `json:"end"`
		} `json:"chain"`
	} `json:"bars"`
}

type UserBarsParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserBarsParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetBars(session *api.Session, params *UserBarsParams) (*UserBarsResponse, error) {
	var resp UserBarsResponse
	if err := session.Get("user/bars", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
