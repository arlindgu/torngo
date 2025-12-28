package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserBountiesResponse struct {
	Bounties []struct {
		TargetID    int    `json:"target_id"`
		TargetName  string `json:"target_name"`
		TargetLevel int    `json:"target_level"`
		ListerID    int    `json:"lister_id"`
		ListerName  string `json:"lister_name"`
		Reward      int64  `json:"reward"`
		Reason      string `json:"reason"`
		Quantity    int    `json:"quantity"`
		IsAnonymous bool   `json:"is_anonymous"`
		ValidUntil  int    `json:"valid_until"`
	} `json:"bounties"`
}

type UserBountiesParams struct {
	api.BaseParams
}

func CreateBountiesURL(t *UserBountiesParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "bounties")

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
