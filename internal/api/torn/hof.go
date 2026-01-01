package torn

type TornHofResponse struct {
	Hof []struct {
		ID               int    `json:"id"`
		Username         string `json:"username"`
		FactionID        int    `json:"faction_id"`
		Level            int    `json:"level"`
		LastAction       int    `json:"last_action"`
		RankName         string `json:"rank_name"`
		RankNumber       int    `json:"rank_number"`
		Position         int    `json:"position"`
		SignedUp         int    `json:"signed_up"`
		AgeInDays        int    `json:"age_in_days"`
		Value            int    `json:"value"`
		Rank             string `json:"rank"`
		CriminalOffenses int    `json:"criminal_offenses,omitempty"`
	} `json:"hof"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}
