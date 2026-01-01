package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserDiscordResponse struct {
	Discord struct {
		DiscordID string `json:"discord_id"`
		UserID    int    `json:"user_id"`
	} `json:"discord"`
}

type UserDiscordParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserDiscordParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetDiscord(session *api.Session, params *UserDiscordParams) (*UserDiscordResponse, error) {
	var resp UserDiscordResponse
	if err := session.Get("user/discord", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
