package racing

import "torngo/internal/api"

type RacingCarUpgradesResponse struct {
	Carupgrades []struct {
		ID            int    `json:"id"`
		ClassRequired string `json:"class_required"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		Category      string `json:"category"`
		Subcategory   string `json:"subcategory"`
		Effects       struct {
			TopSpeed     int `json:"top_speed"`
			Acceleration int `json:"acceleration"`
			Braking      int `json:"braking"`
			Handling     int `json:"handling"`
			Safety       int `json:"safety"`
			Dirt         int `json:"dirt"`
			Tarmac       int `json:"tarmac"`
		} `json:"effects"`
		Cost struct {
			Points int `json:"points"`
			Cash   int `json:"cash"`
		} `json:"cost"`
	} `json:"carupgrades"`
}

type RacingCarUpgradesParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
