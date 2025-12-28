package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserNotificationsResponse struct {
	Notifications struct {
		Messages    int `json:"messages"`
		Events      int `json:"events"`
		Awards      int `json:"awards"`
		Competition int `json:"competition"`
	} `json:"notifications"`
}

type UserNotificationsParams struct {
	api.BaseParams
}

func CreateNotificationsURL(t *UserNotificationsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "notifications")

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
