package faction

import "torngo/internal/api"

type FactionApplicationsResponse struct {
	Applications []struct {
		ID   int `json:"id"`
		User struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Level int    `json:"level"`
			Stats struct {
				Strength  int64 `json:"strength"`
				Speed     int64 `json:"speed"`
				Dexterity int64 `json:"dexterity"`
				Defense   int64 `json:"defense"`
			} `json:"stats"`
		} `json:"user"`
		Message    string `json:"message"`
		ValidUntil int    `json:"valid_until"`
		Status     string `json:"status"`
	} `json:"applications"`
}

type FactionApplicationsParams struct {
	Timestamp api.ApiTimestamp
	Comment   api.ApiComment
}
