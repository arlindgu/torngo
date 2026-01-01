package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserRevivesFullResponse struct {
	Revives []struct {
		ID      int `json:"id"`
		Reviver struct {
			ID        int `json:"id"`
			FactionID int `json:"faction_id"`
		} `json:"reviver"`
		Target struct {
			ID             int    `json:"id"`
			FactionID      int    `json:"faction_id"`
			HospitalReason string `json:"hospital_reason"`
			EarlyDischarge bool   `json:"early_discharge"`
			LastAction     int    `json:"last_action"`
			OnlineStatus   string `json:"online_status"`
		} `json:"target"`
		SuccessChance float64 `json:"success_chance"`
		Result        string  `json:"result"`
		Timestamp     int     `json:"timestamp"`
	} `json:"revives"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserRevivesFullFilter string

const (
	RevivesFullFilterIncoming UserRevivesFullFilter = "incoming"
	RevivesFullFilterOutgoing UserRevivesFullFilter = "outgoing"
	RevivesFullFilterIDFilter UserRevivesFullFilter = "idFilter"
)

type UserRevivesFullParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
	To        api.ApiTo
	From      api.ApiFrom
	Limit     api.ApiLimit
	Filter    UserRevivesFullFilter
	Sort      api.ApiSort
	Striptags api.ApiStriptags
}

func (p *UserRevivesFullParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "filter", string(p.Filter))
	api.SetIfNotZero(q, "limit", int(p.Limit))
	api.SetIfNotZero(q, "sort", string(p.Sort))
	api.SetIfNotZero(q, "to", int(p.To))
	api.SetIfNotZero(q, "from", int(p.From))
	api.SetIfNotZero(q, "striptags", fmt.Sprintf("%v", p.Striptags))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetRevivesFull(session *api.Session, params *UserRevivesFullParams) (*UserRevivesFullResponse, error) {
	var resp UserRevivesFullResponse
	if err := session.Get("user/revives", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
