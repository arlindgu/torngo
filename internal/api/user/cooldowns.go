package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserCooldownsResponse struct {
	Cooldowns struct {
		Drug    int `json:"drug"`
		Medical int `json:"medical"`
		Booster int `json:"booster"`
	} `json:"cooldowns"`
}

type UserCooldownsParams struct {
	api.BaseParams
}

func CreateCooldownsURL(t *UserCooldownsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "cooldowns")

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
