package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserOrganizedCrimeResponse struct {
	OrganizedCrime struct {
		ID              int64  `json:"id"`
		PreviousCrimeID int64  `json:"previous_crime_id"`
		Name            string `json:"name"`
		Difficulty      int    `json:"difficulty"`
		Status          string `json:"status"`
		CreatedAt       int    `json:"created_at"`
		PlanningAt      int    `json:"planning_at"`
		ReadyAt         int    `json:"ready_at"`
		ExpiredAt       int    `json:"expired_at"`
		ExecutedAt      int    `json:"executed_at"`
		Slots           []struct {
			Position        string `json:"position"`
			PositionID      string `json:"position_id"`
			PositionNumber  int    `json:"position_number"`
			ItemRequirement struct {
				ID          int64 `json:"id"`
				IsReusable  bool  `json:"is_reusable"`
				IsAvailable bool  `json:"is_available"`
			} `json:"item_requirement"`
			User struct {
				ID              int    `json:"id"`
				Outcome         string `json:"outcome"`
				OutcomeDuration int    `json:"outcome_duration"`
				ItemOutcome     struct {
					OwnedBy string `json:"owned_by"`
					ItemID  int64  `json:"item_id"`
					ItemUID int64  `json:"item_uid"`
					Outcome string `json:"outcome"`
				} `json:"item_outcome"`
				JoinedAt int     `json:"joined_at"`
				Progress float64 `json:"progress"`
			} `json:"user"`
			CheckpointPassRate int `json:"checkpoint_pass_rate"`
		} `json:"slots"`
		Rewards struct {
			Money int `json:"money"`
			Items []struct {
				ID       int64 `json:"id"`
				Quantity int   `json:"quantity"`
			} `json:"items"`
			Respect int `json:"respect"`
			Scope   int `json:"scope"`
			Payout  struct {
				Type       string `json:"type"`
				Percentage int    `json:"percentage"`
				PaidBy     int    `json:"paid_by"`
				PaidAt     int    `json:"paid_at"`
			} `json:"payout"`
		} `json:"rewards"`
	} `json:"organizedCrime"`
}

type UserOrganizedCrimeParams struct {
	api.BaseParams
}

func CreateOrganizedCrimeURL(t *UserOrganizedCrimeParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "organized_crime")

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
