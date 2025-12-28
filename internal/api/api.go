package api

import (
	"fmt"
	"net/url"
)

type ErrorResponse struct {
	Error struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	} `json:"error"`
}

type SortParams string

const (
	Ascending  SortParams = "ASC"
	Descending SortParams = "DESC"
)

type BaseParams struct {
	Timestamp int
	Comment   string
	APIKey    string
}

type RangeParams struct {
	From int32
	To   int32
}

type AccessLevel int16

const (
	Public  AccessLevel = 1
	Minimal AccessLevel = 2
	Limited AccessLevel = 3
	Full    AccessLevel = 4
)

const APIBaseURL = "https://api.torn.com/v2/"

func SetRangeParams(q url.Values, from, to int32) error {
	if from != 0 && to != 0 && from > to {
		return fmt.Errorf("From parameter can't be greater than To parameter")
	}

	if to != 0 {
		q.Set("to", fmt.Sprintf("%d", to))
	}

	if from != 0 {
		q.Set("from", fmt.Sprintf("%d", from))
	}

	return nil
}

func SetLimitParams(q url.Values, limit, defaultLimit int) {
	if limit != 0 {
		q.Set("limit", fmt.Sprintf("%d", limit))
	} else {
		q.Set("limit", fmt.Sprintf("%d", defaultLimit))
	}
}
