package user

import (
	"fmt"
	"net/url"
	"path"
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
	AttacksIncoming UserAttacksFilter = "incoming"
	AttacksOutgoing UserAttacksFilter = "outgoing"
	AttacksIDFilter UserAttacksFilter = "idFilter"
)

type UserAttacksParams struct {
	api.BaseParams
	Filter UserAttacksFilter
	Limit  int
	Sort   api.SortParams
	api.RangeParams
}

func CreateAttacksURL(t *UserAttacksParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}
	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "attacks")

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

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
