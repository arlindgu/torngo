package torn

type TornItemAmmoResponse struct {
	Itemammo []struct {
		ID    int      `json:"id"`
		Name  string   `json:"name"`
		Price int64    `json:"price"`
		Types []string `json:"types"`
	} `json:"itemammo"`
}
