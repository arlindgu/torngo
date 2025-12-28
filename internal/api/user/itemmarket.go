package user

import (
	"fmt"
	"net/url"
	"path"
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
	Offset int32
	api.BaseParams
}

func CreateItemmarketURL(t *UserItemmarketParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "itemmarket")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Offset != 0 {
		q.Set("offset", fmt.Sprintf("%d", t.Offset))
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
