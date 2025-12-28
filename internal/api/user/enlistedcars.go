package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserEnlistedCarsResponse struct {
	Enlistedcars []struct {
		CarItemID    int64  `json:"car_item_id"`
		CarItemName  string `json:"car_item_name"`
		TopSpeed     int    `json:"top_speed"`
		Acceleration int    `json:"acceleration"`
		Braking      int    `json:"braking"`
		Dirt         int    `json:"dirt"`
		Handling     int    `json:"handling"`
		Safety       int    `json:"safety"`
		Tarmac       int    `json:"tarmac"`
		Class        string `json:"class"`
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Worth        int64  `json:"worth"`
		PointsSpent  int    `json:"points_spent"`
		RacesEntered int    `json:"races_entered"`
		RacesWon     int    `json:"races_won"`
		IsRemoved    bool   `json:"is_removed"`
		Parts        []int  `json:"parts"`
	} `json:"enlistedcars"`
}

type UserEnlistedCarsParams struct {
	api.BaseParams
}

func CreateEnlistedCarsURL(t *UserEnlistedCarsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "enlistedcars")

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
