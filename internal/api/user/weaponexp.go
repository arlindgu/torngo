package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserWeaponExpResponse struct {
	Weaponexp []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Exp  int    `json:"exp"`
	} `json:"weaponexp"`
}

type UserWeaponExpParams struct {
	api.BaseParams
}

func CreateWeaponExpURL(t *UserWeaponExpParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "weaponexp")

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
