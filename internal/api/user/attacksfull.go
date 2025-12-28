package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserAttacksFullResponse struct {
	Attacks []struct {
		ID       int    `json:"id"`
		Code     string `json:"code"`
		Started  int    `json:"started"`
		Ended    int    `json:"ended"`
		Attacker struct {
			ID        int `json:"id"`
			FactionID int `json:"faction_id"`
		} `json:"attacker"`
		Defender struct {
			ID        int `json:"id"`
			FactionID int `json:"faction_id"`
		} `json:"defender"`
		Result      string  `json:"result"`
		RespectGain float64 `json:"respect_gain"`
		RespectLoss float64 `json:"respect_loss"`
	} `json:"attacks"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserAttacksFullFilter string

const (
	AttacksFullIncoming UserAttacksFullFilter = "incoming"
	AttacksFullOutgoing UserAttacksFullFilter = "outgoing"
	AttacksFullIDFilter UserAttacksFullFilter = "idFilter"
)

type UserAttacksFullSort string

const (
	AttacksFullAscending  UserAttacksFullSort = "ASC"
	AttacksFullDescending UserAttacksFullSort = "DESC"
)

type UserAttacksFullParams struct {
	Filter UserAttacksFullFilter
	Limit  int32
	Sort   UserAttacksFullSort
	api.RangeParams
	api.BaseParams
}

func CreateAttacksFullURL(t *UserAttacksFullParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "attacksfull")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Filter != "" {
		q.Set("filter", string(t.Filter))
	}
	if t.Limit != 0 {
		q.Set("limit", fmt.Sprintf("%d", t.Limit))
	}

	if t.From > t.To {
		return "", fmt.Errorf("From parameter can't be greater than To parameter")
	} else {
		if t.To != 0 {
			q.Set("to", fmt.Sprintf("%d", t.To))
		}
		if t.From != 0 {
			q.Set("from", fmt.Sprintf("%d", t.From))
		}
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
