package user

import (
	"fmt"
	"net/url"
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
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserOrganizedCrimeParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetOrganizedCrime(session *api.Session, params *UserOrganizedCrimeParams) (*UserOrganizedCrimeResponse, error) {
	var resp UserOrganizedCrimeResponse
	if err := session.Get("user/organizedcrime", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
