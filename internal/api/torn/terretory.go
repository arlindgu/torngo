package torn

type TornTerritoryResponse struct {
	Territory []struct {
		ID          string `json:"id"`
		Sector      int    `json:"sector"`
		Size        int    `json:"size"`
		Density     int    `json:"density"`
		Slots       int    `json:"slots"`
		Respect     int    `json:"respect"`
		Coordinates struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"coordinates"`
		Neighbors []string `json:"neighbors"`
	} `json:"territory"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}
