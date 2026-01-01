package torn

type TornMeritsResponse struct {
	Merits []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"merits"`
}
