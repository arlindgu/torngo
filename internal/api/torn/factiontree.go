package torn

type TornFactionTreeResponse struct {
	FactionTree []struct {
		Name     string `json:"name"`
		Branches []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Upgrades []struct {
				Name      string `json:"name"`
				Level     int    `json:"level"`
				Ability   string `json:"ability"`
				Cost      int    `json:"cost"`
				Challenge struct {
					Description    string `json:"description"`
					AmountRequired int64  `json:"amount_required"`
					Stat           string `json:"stat"`
				} `json:"challenge"`
			} `json:"upgrades"`
		} `json:"branches"`
	} `json:"factionTree"`
}
