package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserItemmarketResponse struct {
	Itemmarket []struct {
		ID           int64 `json:"id"`
		Price        int64 `json:"price"`
		AveragePrice int64 `json:"average_price"`
		Amount       int   `json:"amount"`
		IsAnonymous  bool  `json:"is_anonymous"`
		Available    int   `json:"available"`
		Item         struct {
			ID     int64  `json:"id"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			Rarity string `json:"rarity"`
			UID    int64  `json:"uid"`
			Stats  struct {
				Damage   float64 `json:"damage"`
				Accuracy float64 `json:"accuracy"`
				Armor    float64 `json:"armor"`
				Quality  float64 `json:"quality"`
			} `json:"stats"`
			Bonuses []struct {
				ID          int    `json:"id"`
				Title       string `json:"title"`
				Description string `json:"description"`
				Value       int    `json:"value"`
			} `json:"bonuses"`
		} `json:"item"`
	} `json:"itemmarket"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserItemmarketParams struct {
	Offset    api.ApiOffset
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserItemmarketParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "offset", p.Offset)
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetItemmarket(session *api.Session, params *UserItemmarketParams) (*UserItemmarketResponse, error) {
	var resp UserItemmarketResponse
	if err := session.Get("user/itemmarket", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
