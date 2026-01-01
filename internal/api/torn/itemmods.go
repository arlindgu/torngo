package torn

type TornItemModsResponse struct {
	Itemmods []struct {
		ID          int      `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		DualFit     bool     `json:"dual_fit"`
		Weapons     []string `json:"weapons"`
	} `json:"itemmods"`
}