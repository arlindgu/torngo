package torn

type TornCalendarResponse struct {
	Calendar struct {
		Competitions []struct {
			Title          string `json:"title"`
			Description    string `json:"description"`
			Start          int    `json:"start"`
			End            int    `json:"end"`
			FixedStartTime bool   `json:"fixed_start_time"`
		} `json:"competitions"`
		Events []struct {
			Title          string `json:"title"`
			Description    string `json:"description"`
			Start          int    `json:"start"`
			End            int    `json:"end"`
			FixedStartTime bool   `json:"fixed_start_time"`
		} `json:"events"`
	} `json:"calendar"`
}
