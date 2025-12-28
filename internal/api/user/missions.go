package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserMissionsResponse struct {
	Missions struct {
		Credits int `json:"credits"`
		Givers  []struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Contracts []struct {
				Title       string `json:"title"`
				Difficulty  string `json:"difficulty"`
				Status      string `json:"status"`
				CreatedAt   int    `json:"created_at"`
				StartedAt   int    `json:"started_at"`
				ExpiresAt   int    `json:"expires_at"`
				CompletedAt int    `json:"completed_at"`
				Rewards     struct {
					Money   int `json:"money"`
					Credits int `json:"credits"`
				} `json:"rewards"`
			} `json:"contracts"`
		} `json:"givers"`
		Rewards []struct {
			Type    string `json:"type"`
			Details struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"details"`
			Amount    int `json:"amount"`
			Cost      int `json:"cost"`
			ExpiresAt int `json:"expires_at"`
		} `json:"rewards"`
	} `json:"missions"`
}

type UserMissionsParams struct {
	api.BaseParams
}

func CreateMissionsURL(t *UserMissionsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "missions")

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
