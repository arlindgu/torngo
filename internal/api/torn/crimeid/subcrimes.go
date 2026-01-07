package crimeid

type TornCrimeIdSubcrimesResponse struct {
	Subcrimes []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		NerveCost int    `json:"nerve_cost"`
	} `json:"subcrimes"`
}

