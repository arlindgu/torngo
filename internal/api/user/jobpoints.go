package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserJobpointsResponse struct {
	Jobpoints struct {
		Jobs struct {
			Army      int `json:"army"`
			Casino    int `json:"casino"`
			Education int `json:"education"`
			Grocer    int `json:"grocer"`
			Law       int `json:"law"`
			Medical   int `json:"medical"`
		} `json:"jobs"`
		Companies []struct {
			Company struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"company"`
			Points int `json:"points"`
		} `json:"companies"`
	} `json:"jobpoints"`
}

type UserJobpointsParams struct {
	api.BaseParams
}

func CreateJobpointsURL(t *UserJobpointsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "jobpoints")

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
