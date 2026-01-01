package torn

type TornEliminationResponse struct {
	Elimination []struct {
		ID                  int    `json:"id"`
		Name                string `json:"name"`
		Participants        int    `json:"participants"`
		Position            int    `json:"position"`
		Score               int    `json:"score"`
		Lives               int    `json:"lives"`
		Wins                int    `json:"wins"`
		Losses              int    `json:"losses"`
		Eliminated          bool   `json:"eliminated"`
		EliminatedTimestamp int    `json:"eliminated_timestamp"`
		Leaders             []struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Active bool   `json:"active"`
		} `json:"leaders"`
	} `json:"elimination"`
}
