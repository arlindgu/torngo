package torn

type TornPropertiesResponse struct {
	Properties []struct {
		ID            int      `json:"id"`
		Name          string   `json:"name"`
		Cost          int      `json:"cost"`
		Happy         int      `json:"happy"`
		Upkeep        int      `json:"upkeep"`
		Modifications []string `json:"modifications"`
		Staff         []string `json:"staff"`
	} `json:"properties"`
}
