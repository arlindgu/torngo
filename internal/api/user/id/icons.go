package id

type UserIdIconsResponse struct {
	Icons []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Until       int    `json:"until"`
	} `json:"icons"`
}
