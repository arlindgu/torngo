package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserAttacksResponse struct {
	Attacks []struct {
		ID       int    `json:"id"`
		Code     string `json:"code"`
		Started  int    `json:"started"`
		Ended    int    `json:"ended"`
		Attacker struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Level   int    `json:"level"`
			Faction struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"faction"`
		} `json:"attacker"`
		Defender struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Level   int    `json:"level"`
			Faction struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"faction"`
		} `json:"defender"`
		Result              string  `json:"result"`
		RespectGain         float64 `json:"respect_gain"`
		RespectLoss         float64 `json:"respect_loss"`
		Chain               int     `json:"chain"`
		IsInterrupted       bool    `json:"is_interrupted"`
		IsStealthed         bool    `json:"is_stealthed"`
		IsRaid              bool    `json:"is_raid"`
		IsRankedWar         bool    `json:"is_ranked_war"`
		FinishingHitEffects []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"finishing_hit_effects"`
		Modifiers struct {
			FairFight   float64 `json:"fair_fight"`
			War         float64 `json:"war"`
			Retaliation float64 `json:"retaliation"`
			Group       float64 `json:"group"`
			Overseas    float64 `json:"overseas"`
			Chain       float64 `json:"chain"`
			Warlord     float64 `json:"warlord"`
		} `json:"modifiers"`
	} `json:"attacks"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserAttacksFilter string

const (
	UserAttacksIncoming UserAttacksFilter = "incoming"
	UserAttacksOutgoing UserAttacksFilter = "outgoing"
	UserAttacksIDFilter UserAttacksFilter = "idFilter"
)

type UserAttacksParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
	Filter    UserAttacksFilter
	Limit     api.ApiLimit
	From      api.ApiFrom
	To        api.ApiTo
	Sort      api.ApiSort
}

func (p *UserAttacksParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "filter", string(p.Filter))
	api.SetIfNotZero(q, "limit", int(p.Limit))
	api.SetIfNotZero(q, "from", int(p.From))
	api.SetIfNotZero(q, "to", int(p.To))
	api.SetIfNotZero(q, "sort", string(p.Sort))
	return q
}

func GetAttacks(session *api.Session, params *UserAttacksParams) (*UserAttacksResponse, error) {
	var resp UserAttacksResponse
	if err := session.Get("user/attacks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
