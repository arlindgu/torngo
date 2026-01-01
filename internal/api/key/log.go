package key

type KeyLogResponse struct {
	Info struct {
		Selections struct {
			Company  []string `json:"company"`
			Faction  []string `json:"faction"`
			Market   []string `json:"market"`
			Property []string `json:"property"`
			Torn     []string `json:"torn"`
			User     []string `json:"user"`
			Racing   []string `json:"racing"`
			Forum    []string `json:"forum"`
			Key      []string `json:"key"`
		} `json:"selections"`
		User struct {
			FactionID int `json:"faction_id"`
			CompanyID int `json:"company_id"`
			ID        int `json:"id"`
		} `json:"user"`
		Access struct {
			Level   int    `json:"level"`
			Type    string `json:"type"`
			Faction bool   `json:"faction"`
			Company bool   `json:"company"`
			Log     struct {
				CustomPermissions bool `json:"custom_permissions"`
				Available         []struct {
					CategoryID int   `json:"category_id"`
					LogIds     []int `json:"log_ids"`
				} `json:"available"`
			} `json:"log"`
		} `json:"access"`
	} `json:"info"`
}
