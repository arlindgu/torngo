package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserBarsResponse struct {
	Bars struct {
		Energy struct {
			Current   int `json:"current"`
			Maximum   int `json:"maximum"`
			Increment int `json:"increment"`
			Interval  int `json:"interval"`
			TickTime  int `json:"tick_time"`
			FullTime  int `json:"full_time"`
		} `json:"energy"`
		Nerve struct {
			Current   int `json:"current"`
			Maximum   int `json:"maximum"`
			Increment int `json:"increment"`
			Interval  int `json:"interval"`
			TickTime  int `json:"tick_time"`
			FullTime  int `json:"full_time"`
		} `json:"nerve"`
		Happy struct {
			Current   int `json:"current"`
			Maximum   int `json:"maximum"`
			Increment int `json:"increment"`
			Interval  int `json:"interval"`
			TickTime  int `json:"tick_time"`
			FullTime  int `json:"full_time"`
		} `json:"happy"`
		Life struct {
			Current   int `json:"current"`
			Maximum   int `json:"maximum"`
			Increment int `json:"increment"`
			Interval  int `json:"interval"`
			TickTime  int `json:"tick_time"`
			FullTime  int `json:"full_time"`
		} `json:"life"`
		Chain struct {
			ID       int `json:"id"`
			Current  int `json:"current"`
			Max      int `json:"max"`
			Timeout  int `json:"timeout"`
			Modifier int `json:"modifier"`
			Cooldown int `json:"cooldown"`
			Start    int `json:"start"`
			End      int `json:"end"`
		} `json:"chain"`
	} `json:"bars"`
}

type UserBarsParams struct {
	api.BaseParams
}

func createBarsURL(t *UserBarsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "bars")

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
