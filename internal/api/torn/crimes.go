package torn

type TornCrimesResponse struct {
	Crimes []struct {
		ID                  int      `json:"id"`
		Name                string   `json:"name"`
		CategoryID          int      `json:"category_id"`
		CategoryName        string   `json:"category_name"`
		EnhancerID          int      `json:"enhancer_id"`
		EnhancerName        string   `json:"enhancer_name"`
		UniqueOutcomesCount int      `json:"unique_outcomes_count"`
		UniqueOutcomesIds   []int    `json:"unique_outcomes_ids"`
		Notes               []string `json:"notes"`
	} `json:"crimes"`
}
