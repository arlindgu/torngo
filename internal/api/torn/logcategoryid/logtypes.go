package logcategoryid

type TornCategoryIdLogResponse struct {
	Logtypes []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"logtypes"`
}
