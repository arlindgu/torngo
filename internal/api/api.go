package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type (
	ApiTimestamp int
	ApiComment   string
	ApiName      string
	ApiLimit     int32
	ApiOffset    int32
	ApiFrom      int32
	ApiTo        int32
	ApiSort      string
	ApiStriptags bool
	ApiId        uint32
	ApiDiscordId uint64
)

const (
	Ascending  ApiSort = "ASC"
	Descending ApiSort = "DESC"
)

type Session struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

func NewSession(apikey string) *Session {
	return &Session{
		httpClient: &http.Client{},
		baseURL:    APIBaseURL,
		apiKey:     apikey,
	}
}

type ParamEncoder interface {
	EncodeParams() url.Values
}

func (s *Session) Get(urlPath string, params ParamEncoder, result interface{}) error {
	u, err := url.Parse(s.baseURL)
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, urlPath)

	q := u.Query()
	q.Set("key", s.apiKey)

	if params != nil {
		customParams := params.EncodeParams()
		for k, v := range customParams {
			q[k] = v
		}
	}

	u.RawQuery = q.Encode()

	resp, err := s.httpClient.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Den Body lesen und loggen
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Jetzt decodieren
	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return err
	}

	return nil
}

type ErrorResponse struct {
	Error struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	} `json:"error"`
}

const APIBaseURL = "https://api.torn.com/v2"

func SetIfNotZero[T comparable](q url.Values, key string, value T) {
	var zero T
	if value != zero {
		q.Set(key, fmt.Sprintf("%v", value))
	}
}
