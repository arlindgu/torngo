package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserListResponse struct {
	List []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Level      int    `json:"level"`
		FactionID  int    `json:"faction_id"`
		LastAction struct {
			Status    string `json:"status"`
			Timestamp int    `json:"timestamp"`
			Relative  string `json:"relative"`
		} `json:"last_action"`
		Status struct {
			Description    string `json:"description"`
			Details        string `json:"details"`
			State          string `json:"state"`
			Color          string `json:"color"`
			Until          int    `json:"until"`
			PlaneImageType string `json:"plane_image_type"`
		} `json:"status"`
	} `json:"list"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type Category string

const (
	UserListCategoryFriends Category = "friends"
	UserListCategoryEnemies Category = "enemies"
	UserListCategoryTargets Category = "targets"
)

type UserListParams struct {
	Cat       Category
	Striptags api.ApiStriptags
	Limit     api.ApiLimit
	Offset    api.ApiOffset
	Sort      api.ApiSort
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserListParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "cat", string(p.Cat))
	api.SetIfNotZero(q, "striptags", fmt.Sprintf("%t", p.Striptags))
	api.SetIfNotZero(q, "limit", p.Limit)
	api.SetIfNotZero(q, "offset", p.Offset)
	api.SetIfNotZero(q, "sort", string(p.Sort))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetList(session *api.Session, params *UserListParams) (*UserListResponse, error) {
	var resp UserListResponse
	if err := session.Get("user/list", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
