package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserRevivesResponse struct {
	Revives []struct {
		ID      int `json:"id"`
		Reviver struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Faction struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"faction"`
			Skill float64 `json:"skill"`
		} `json:"reviver"`
		Target struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Faction struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"faction"`
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

type UserRevivesFilter string

const (
	RevivesFilterIncoming UserRevivesFilter = "incoming"
	RevivesFilterOutgoing UserRevivesFilter = "outgoing"
	RevivesFilterIDFilter UserRevivesFilter = "idFilter"
)

type UserRevivesParams struct {
	Filter    UserRevivesFilter
	Limit     api.ApiLimit
	Sort      api.ApiSort
	To        api.ApiTo
	From      api.ApiFrom
	Striptags api.ApiStriptags
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserRevivesParams) EncodeParams() url.Values {
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

func GetRevives(session *api.Session, params *UserRevivesParams) (*UserRevivesResponse, error) {
	var resp UserRevivesResponse
	if err := session.Get("user/revives", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
