package user

import (
	"fmt"
	"net/url"
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
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserNotificationsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetNotifications(session *api.Session, params *UserNotificationsParams) (*UserNotificationsResponse, error) {
	var resp UserNotificationsResponse
	if err := session.Get("user/notifications", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
