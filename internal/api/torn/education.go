package torn

type TornEducationResponse struct {
	Education []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Courses []struct {
			ID          int    `json:"id"`
			Code        string `json:"code"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Duration    int    `json:"duration"`
			Rewards     struct {
				WorkingStats struct {
					ManualLabor  int `json:"manual_labor"`
					Intelligence int `json:"intelligence"`
					Endurance    int `json:"endurance"`
				} `json:"working_stats"`
				Effect string `json:"effect"`
				Honor  string `json:"honor"`
			} `json:"rewards"`
			Prerequisites struct {
				Cost    int   `json:"cost"`
				Courses []int `json:"courses"`
			} `json:"prerequisites"`
		} `json:"courses"`
	} `json:"education"`
}
