package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserJobranksResponse struct {
	Jobranks struct {
		Army      string `json:"army"`
		Grocer    string `json:"grocer"`
		Casino    string `json:"casino"`
		Medical   string `json:"medical"`
		Law       string `json:"law"`
		Education string `json:"education"`
	} `json:"jobranks"`
}

type UserJobranksParams struct {
	api.BaseParams
}

func CreateJobranksURL(t *UserJobranksParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "jobranks")

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
