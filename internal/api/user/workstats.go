package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserWorkstatsResponse struct {
	Workstats struct {
		Endurance    int `json:"endurance"`
		Intelligence int `json:"intelligence"`
		ManualLabor  int `json:"manual_labor"`
		Total        int `json:"total"`
	} `json:"workstats"`
}

type UserWorkstatsParams struct {
	api.BaseParams
}

func CreateWorkstatsURL(t *UserWorkstatsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "workstats")

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
