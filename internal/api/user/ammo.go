package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserAmmoResponse struct {
	Ammo []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Types []struct {
			Name     string `json:"name"`
			Quantity int    `json:"quantity"`
			Equipped bool   `json:"equipped"`
		} `json:"types"`
	} `json:"ammo"`
}

type UserAmmoParams struct {
	api.BaseParams
}

func CreateAmmoURL(t *UserAmmoParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "ammo")

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

func GetAmmo(t *UserAmmoParams) (*UserAmmoResponse, error) {
	url, err := CreateAmmoURL(t)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var ammoResp UserAmmoResponse
	var errorResp api.ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&errorResp); err == nil && errorResp.Error.Code != 0 {
		return nil, fmt.Errorf("API Error %d: %s", errorResp.Error.Code, errorResp.Error.Error)
	}

	if err := json.NewDecoder(resp.Body).Decode(&ammoResp); err != nil {
		return nil, err
	}

	return &ammoResp, nil
}
