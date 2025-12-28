package user

import (
	"fmt"
	"net/url"
	"path"
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
	api.BaseParams
}

func CreateEquipmentURL(t *UserEquipmentParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "equipment")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
