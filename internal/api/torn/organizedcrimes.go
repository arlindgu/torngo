package torn

type TornOrganizedCrimesResponse struct {
	Organizedcrimes []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Difficulty  int    `json:"difficulty"`
		Spawn       struct {
			Level int    `json:"level"`
			Name  string `json:"name"`
		} `json:"spawn"`
		Scope struct {
			Cost   int `json:"cost"`
			Return int `json:"return"`
		} `json:"scope"`
		Prerequisite string `json:"prerequisite"`
		Slots        []struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			RequiredItem struct {
				ID     int64  `json:"id"`
				Name   string `json:"name"`
				IsUsed bool   `json:"is_used"`
			} `json:"required_item"`
		} `json:"slots"`
	} `json:"organizedcrimes"`
}
