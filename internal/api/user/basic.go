package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserBasicResponse struct {
	Profile struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Level  int    `json:"level"`
		Gender string `json:"gender"`
		Status struct {
			Description    string `json:"description"`
			Details        string `json:"details"`
			State          string `json:"state"`
			Color          string `json:"color"`
			Until          int    `json:"until"`
			PlaneImageType string `json:"plane_image_type"`
		} `json:"status"`
	} `json:"profile"`
}

type UserBasicParams struct {
	Striptags bool
	api.BaseParams
}

func createBasicURL(t *UserBasicParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "basic")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
