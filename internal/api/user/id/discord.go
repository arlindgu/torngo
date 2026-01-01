package id

type UserIdDiscordResponse struct {
	Discord struct {
		DiscordID string `json:"discord_id"`
		UserID    int    `json:"user_id"`
	} `json:"discord"`
}
