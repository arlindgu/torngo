package torn

type TornFactionHofResponse struct {
	Factionhof []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Members  int    `json:"members"`
		Position int    `json:"position"`
		Rank     string `json:"rank"`
		Values   struct {
			Chain         int `json:"chain"`
			ChainDuration int `json:"chain_duration"`
			Respect       int `json:"respect"`
		} `json:"values"`
	} `json:"factionhof"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}
