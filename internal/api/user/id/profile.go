package id

type UserIdProfileResponse struct {
	Profile struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Level         int    `json:"level"`
		Rank          string `json:"rank"`
		Title         string `json:"title"`
		DonatorStatus string `json:"donator_status"`
		Age           int    `json:"age"`
		SignedUp      int    `json:"signed_up"`
		FactionID     int    `json:"faction_id"`
		HonorID       int    `json:"honor_id"`
		Property      struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"property"`
		Image     string `json:"image"`
		Gender    string `json:"gender"`
		Revivable bool   `json:"revivable"`
		Role      string `json:"role"`
		Status    struct {
			Description    string `json:"description"`
			Details        string `json:"details"`
			State          string `json:"state"`
			Color          string `json:"color"`
			Until          int    `json:"until"`
			PlaneImageType string `json:"plane_image_type"`
		} `json:"status"`
		Spouse struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Status      string `json:"status"`
			DaysMarried int    `json:"days_married"`
		} `json:"spouse"`
		Awards     int `json:"awards"`
		Friends    int `json:"friends"`
		Enemies    int `json:"enemies"`
		ForumPosts int `json:"forum_posts"`
		Karma      int `json:"karma"`
		LastAction struct {
			Status    string `json:"status"`
			Timestamp int    `json:"timestamp"`
			Relative  string `json:"relative"`
		} `json:"last_action"`
		Life struct {
			Current int `json:"current"`
			Maximum int `json:"maximum"`
		} `json:"life"`
	} `json:"profile"`
}
