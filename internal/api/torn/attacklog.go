package torn

type TornAttacklogResponse struct {
	Attacklog struct {
		Log []struct {
			Text      string `json:"text"`
			Timestamp int    `json:"timestamp"`
			Action    string `json:"action"`
			Icon      string `json:"icon"`
			Attacker  struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Item struct {
					ID   int64  `json:"id"`
					Name string `json:"name"`
				} `json:"item"`
			} `json:"attacker"`
			Defender struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"defender"`
			AttackerItem struct {
				ID   int64  `json:"id"`
				Name string `json:"name"`
			} `json:"attacker_item"`
		} `json:"log"`
		Summary []struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Hits   int    `json:"hits"`
			Misses int    `json:"misses"`
			Damage int    `json:"damage"`
		} `json:"summary"`
	} `json:"attacklog"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}
