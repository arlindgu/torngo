package torn

type TornBountiesResponse struct {
	Bounties []struct {
		TargetID    int    `json:"target_id"`
		TargetName  string `json:"target_name"`
		TargetLevel int    `json:"target_level"`
		ListerID    int    `json:"lister_id"`
		ListerName  string `json:"lister_name"`
		Reward      int64  `json:"reward"`
		Reason      string `json:"reason"`
		Quantity    int    `json:"quantity"`
		IsAnonymous bool   `json:"is_anonymous"`
		ValidUntil  int    `json:"valid_until"`
	} `json:"bounties"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}
