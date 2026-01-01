package torn

type TornLogCategoriesResponse struct {
	Logcategories []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"logcategories"`
}
