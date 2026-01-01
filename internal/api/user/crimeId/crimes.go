package crimeid

type UserCrimeIdCrimesResponse struct {
	Crimes struct {
		NerveSpent       int `json:"nerve_spent"`
		Skill            int `json:"skill"`
		ProgressionBonus int `json:"progression_bonus"`
		Rewards          struct {
			Money int64 `json:"money"`
			Ammo  struct {
				Standard int `json:"standard"`
				Special  int `json:"special"`
			} `json:"ammo"`
			Items []struct {
				ID     int `json:"id"`
				Amount int `json:"amount"`
			} `json:"items"`
		} `json:"rewards"`
		Attempts struct {
			Total        int `json:"total"`
			Success      int `json:"success"`
			Fail         int `json:"fail"`
			CriticalFail int `json:"critical_fail"`
			Subcrimes    []struct {
				ID      int `json:"id"`
				Total   int `json:"total"`
				Success int `json:"success"`
				Fail    int `json:"fail"`
			} `json:"subcrimes"`
		} `json:"attempts"`
		Uniques []struct {
			ID      int64 `json:"id"`
			Rewards struct {
				Items []struct {
					ID     int `json:"id"`
					Amount int `json:"amount"`
				} `json:"items"`
				Money struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"money"`
				Ammo struct {
					Amount int    `json:"amount"`
					Type   string `json:"type"`
				} `json:"ammo"`
			} `json:"rewards"`
		} `json:"uniques"`
		Miscellaneous struct {
			OnlineStore struct {
				Earnings  int `json:"earnings"`
				Visits    int `json:"visits"`
				Customers int `json:"customers"`
				Sales     int `json:"sales"`
			} `json:"online_store"`
			DvdSales struct {
				Action   int `json:"action"`
				Comedy   int `json:"comedy"`
				Drama    int `json:"drama"`
				Fantasy  int `json:"fantasy"`
				Horror   int `json:"horror"`
				Romance  int `json:"romance"`
				Thriller int `json:"thriller"`
				SciFi    int `json:"sci_fi"`
				Total    int `json:"total"`
				Earnings int `json:"earnings"`
			} `json:"dvd_sales"`
			DvdsCopied int `json:"dvds_copied"`
		} `json:"miscellaneous"`
	} `json:"crimes"`
}
