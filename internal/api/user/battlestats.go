package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserBattlestatsResponse struct {
	Battlestats struct {
		Strength struct {
			Value     int64 `json:"value"`
			Modifier  int   `json:"modifier"`
			Modifiers []struct {
				Effect string `json:"effect"`
				Value  int    `json:"value"`
				Type   string `json:"type"`
			} `json:"modifiers"`
		} `json:"strength"`
		Defense struct {
			Value     int64 `json:"value"`
			Modifier  int   `json:"modifier"`
			Modifiers []struct {
				Effect string `json:"effect"`
				Value  int    `json:"value"`
				Type   string `json:"type"`
			} `json:"modifiers"`
		} `json:"defense"`
		Speed struct {
			Value     int64 `json:"value"`
			Modifier  int   `json:"modifier"`
			Modifiers []struct {
				Effect string `json:"effect"`
				Value  int    `json:"value"`
				Type   string `json:"type"`
			} `json:"modifiers"`
		} `json:"speed"`
		Dexterity struct {
			Value     int64 `json:"value"`
			Modifier  int   `json:"modifier"`
			Modifiers []struct {
				Effect string `json:"effect"`
				Value  int    `json:"value"`
				Type   string `json:"type"`
			} `json:"modifiers"`
		} `json:"dexterity"`
		Total int64 `json:"total"`
	} `json:"battlestats"`
}

type UserBattlestatsParams struct {
	api.BaseParams
}

func CreateBattlestatsURL(t *UserBattlestatsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "battlestats")

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
