package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserEducationResponse struct {
	Education struct {
		Complete []int `json:"complete"`
		Current  struct {
			ID    int `json:"id"`
			Until int `json:"until"`
		} `json:"current"`
	} `json:"education"`
}

type UserEducationParams struct {
	api.BaseParams
}

func CreateEducationURL(t *UserEducationParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "education")

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
