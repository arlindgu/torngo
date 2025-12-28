package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserSkillsResponse struct {
	Skills []struct {
		Slug  string  `json:"slug"`
		Name  string  `json:"name"`
		Level float64 `json:"level"`
	} `json:"skills"`
}

type UserSkillsParams struct {
	api.BaseParams
}

func CreateSkillsURL(t *UserSkillsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "skills")

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
