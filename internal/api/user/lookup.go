package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type LookupResponse struct {
	Selections []string `json:"selections"`
}

type LookupParams struct {
	api.BaseParams
}

func CreateLookupURL(t *LookupParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "lookup")

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
