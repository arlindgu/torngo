package torn

type TornLogTypesResponse struct {
	Logtypes []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"logtypes"`
}
