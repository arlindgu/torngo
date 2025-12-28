package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserProfileResponse struct {
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

type UserProfileParams struct {
	Striptags bool
	api.BaseParams
}

func CreateProfileURL(t *UserProfileParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "profile")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}
	if t.Striptags {
		q.Set("striptags", "1")
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
