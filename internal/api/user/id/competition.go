package id

type UserIdCompetitionResponse struct {
	Competition struct {
		Name            string `json:"name"`
		TreatsCollected int    `json:"treats_collected"`
		Basket          struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"basket"`
	} `json:"competition"`
}
