package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserDiscordResponse struct {
	Discord struct {
		DiscordID string `json:"discord_id"`
		UserID    int    `json:"user_id"`
	} `json:"discord"`
}

type UserDiscordParams struct {
	api.BaseParams
}

func CreateDiscordURL(t *UserDiscordParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "discord")

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
