package user

import (
	"fmt"
	"net/url"
	"path"
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
	Striptags bool
	Limit     int
	Offset    int
	Sort      api.SortParams
	api.BaseParams
}

func CreateListURL(t *UserListParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}
	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "list")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Cat != "" {
		q.Set("cat", string(t.Cat))
	}

	api.SetLimitParams(q, t.Limit, 20)

	if t.Sort != "" {
		q.Set("sort", string(t.Sort))
	}

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}
	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
