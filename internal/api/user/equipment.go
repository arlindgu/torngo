package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserEquipmentResponse struct {
	Equipment []struct {
		UID   int64 `json:"uid"`
		Stats struct {
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
		Rarity  string `json:"rarity"`
		ID      int64  `json:"id"`
		Name    string `json:"name"`
		Type    string `json:"type"`
		SubType string `json:"sub_type"`
		Slot    int    `json:"slot"`
	} `json:"equipment"`
	Clothing []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		UID  int64  `json:"uid"`
		Type string `json:"type"`
	} `json:"clothing"`
}

type UserEquipmentParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserEquipmentParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetEquipment(session *api.Session, params *UserEquipmentParams) (*UserEquipmentResponse, error) {
	var resp UserEquipmentResponse
	if err := session.Get("user/equipment", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
