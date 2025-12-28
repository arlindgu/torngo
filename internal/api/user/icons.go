package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserIconsResponse struct {
	Icons []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Until       int    `json:"until"`
	} `json:"icons"`
}

type UserIconsParams struct {
	api.BaseParams
}

func CreateIconsURL(t *UserIconsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "icons")

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
