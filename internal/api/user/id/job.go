package id

type UserIdJobResponse struct {
	Job struct {
		Type     string `json:"type"`
		Name     string `json:"name"`
		Position string `json:"position"`
	} `json:"job"`
}
