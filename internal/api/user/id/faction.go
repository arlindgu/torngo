package id

type UserIdFactionResponse struct {
	Faction struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Tag           string `json:"tag"`
		TagImage      string `json:"tag_image"`
		Position      string `json:"position"`
		DaysInFaction int    `json:"days_in_faction"`
	} `json:"faction"`
}
