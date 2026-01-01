package id

type TornIdEliminationTeamResponse struct {
	Eliminationteam []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Level      int    `json:"level"`
		LastAction struct {
			Status    string `json:"status"`
			Timestamp int    `json:"timestamp"`
			Relative  string `json:"relative"`
		} `json:"last_action"`
		Status struct {
			Description    string `json:"description"`
			Details        string `json:"details"`
			State          string `json:"state"`
			Color          string `json:"color"`
			Until          int    `json:"until"`
			PlaneImageType string `json:"plane_image_type"`
		} `json:"status"`
		Attacks int `json:"attacks"`
		Score   int `json:"score"`
	} `json:"eliminationteam"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}
