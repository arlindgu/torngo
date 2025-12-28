package user

import (
	"fmt"
	"net/url"
	"path"
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
	Filter UserRevivesFullFilter
	Limit  int
	Sort   api.SortParams
	api.RangeParams
	Striptags bool
	api.BaseParams
}

func CreateRevivesFullURL(t *UserRevivesFullParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "revives")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Filter != "" {
		q.Set("filter", string(t.Filter))
	}

	api.SetLimitParams(q, t.Limit, 20)

	if err := api.SetRangeParams(q, t.From, t.To); err != nil {
		return "", err
	}

	if t.Sort != "" {
		q.Set("sort", string(t.Sort))
	}

	if t.Striptags {
		q.Set("striptags", "true")
	}

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
