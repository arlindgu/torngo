package id

type TornIdItemDetailsResponse struct {
	Itemdetails struct {
		UID   int64 `json:"uid"`
		Stats struct {
			Damage   float64 `json:"damage"`
			Accuracy float64 `json:"accuracy"`
			Armor    float64 `json:"armor"`
			Quality  float64 `json:"quality"`
		} `json:"stats"`
		Bonuses []struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Value       int    `json:"value"`
		} `json:"bonuses"`
		Rarity  string `json:"rarity"`
		ID      int64  `json:"id"`
		Name    string `json:"name"`
		Type    string `json:"type"`
		SubType string `json:"sub_type"`
	} `json:"itemdetails"`
}
