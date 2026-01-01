package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserCalendarResponse struct {
	Calendar struct {
		StartTime string `json:"start_time"`
	} `json:"calendar"`
}

type UserCalendarParams struct {
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserCalendarParams) EncodeParams() url.Values {
	q := url.Values{}

	api.SetIfNotZero(q, "comment", string(p.Comment))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))

	return q
}

func GetCalendar(session *api.Session, params *UserCalendarParams) (*UserCalendarResponse, error) {
	var resp UserCalendarResponse
	if err := session.Get("user/calendar", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
