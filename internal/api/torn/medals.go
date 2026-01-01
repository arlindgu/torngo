package torn

type TornMedalsResponse struct {
	Medals []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Type        struct {
			ID    string `json:"id"`
			Title string `json:"title"`
		} `json:"type"`
		Circulation   int    `json:"circulation"`
		Rarity        string `json:"rarity"`
		CrimesVersion string `json:"crimes_version"`
	} `json:"medals"`
}
