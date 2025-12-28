package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserFactionResponse struct {
	Faction struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Tag           string `json:"tag"`
		TagImage      string `json:"tag_image"`
		Position      string `json:"position"`
		DaysInFaction int    `json:"days_in_faction"`
	} `json:"faction"`
}

type UserFactionParams struct {
	api.BaseParams
}

func CreateFactionURL(t *UserFactionParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "faction")

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
