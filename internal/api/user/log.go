package user

import (
	"fmt"
	"net/url"
	"path"
	"strings"
	"torngo/internal/api"
)

type UserLogResponse struct {
	Log []struct {
		ID        string `json:"id"`
		Timestamp int    `json:"timestamp"`
		Details   struct {
			ID       int    `json:"id"`
			Title    string `json:"title"`
			Category string `json:"category"`
		} `json:"details"`
		Data struct {
		} `json:"data"`
		Params struct {
		} `json:"params"`
	} `json:"log"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserLogParams struct {
	Log    []int
	Cat    int
	Target int
	Limit  int
	api.RangeParams
	api.BaseParams
}

func CreateLogURL(t *UserLogParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}
	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "log")

	q := u.Query()
	q.Set("key", t.APIKey)

	if len(t.Log) > 0 {
		var logValues []string
		for _, l := range t.Log {
			logValues = append(logValues, fmt.Sprintf("%d", l))
		}
		q.Set("log", strings.Join(logValues, ","))
	}

	if t.Cat != 0 {
		q.Set("cat", fmt.Sprintf("%d", t.Cat))
	}

	if t.Target != 0 {
		q.Set("target", fmt.Sprintf("%d", t.Target))
	}

	api.SetLimitParams(q, t.Limit, 20)

	if err := api.SetRangeParams(q, t.From, t.To); err != nil {
		return "", err
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
