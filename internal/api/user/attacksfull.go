package user

import (
	"fmt"
	"net/url"
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

type UserAttacksFullParams struct {
	Filter    UserAttacksFullFilter
	Limit     api.ApiLimit
	Sort      api.ApiSort
	From      api.ApiFrom
	To        api.ApiTo
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserAttacksFullParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "filter", string(p.Filter))
	api.SetIfNotZero(q, "limit", p.Limit)
	api.SetIfNotZero(q, "sort", string(p.Sort))
	api.SetIfNotZero(q, "from", p.From)
	api.SetIfNotZero(q, "to", p.To)
	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetAttacksFull(session *api.Session, params *UserAttacksFullParams) (*UserAttacksFullResponse, error) {
	var resp UserAttacksFullResponse
	if err := session.Get("user/attacksfull", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
