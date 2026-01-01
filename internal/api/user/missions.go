package user

import (
	"fmt"
	"net/url"
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

type UserMissionParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserMissionParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetMissions(session *api.Session, params *UserMissionParams) (*UserMissionsResponse, error) {
	var resp UserMissionsResponse
	if err := session.Get("user/missions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
