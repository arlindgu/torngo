package key

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"torngo/internal/api"
)

type KeyInfoResponse struct {
	*api.ErrorResponse
	Info struct {
		Selections struct {
			Company  []string `json:"company"`
			Faction  []string `json:"faction"`
			Market   []string `json:"market"`
			Property []string `json:"property"`
			Torn     []string `json:"torn"`
			User     []string `json:"user"`
			Racing   []string `json:"racing"`
			Forum    []string `json:"forum"`
			Key      []string `json:"key"`
		} `json:"selections"`
		User struct {
			FactionID int `json:"faction_id"`
			CompanyID int `json:"company_id"`
			ID        int `json:"id"`
		} `json:"user"`
		Access struct {
			Level   int    `json:"level"`
			Type    string `json:"type"`
			Faction bool   `json:"faction"`
			Company bool   `json:"company"`
			Log     struct {
				CustomPermissions bool `json:"custom_permissions"`
				Available         []struct {
					CategoryID int   `json:"category_id"`
					LogIds     []int `json:"log_ids"`
				} `json:"available"`
			} `json:"log"`
		} `json:"access"`
	} `json:"info"`
}

type KeyInfoParams struct {
	api.BaseParams
}

func CreateInfoURL(t *KeyInfoParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "key", "info")

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

func GetInfo(t *KeyInfoParams) (*KeyInfoResponse, error) {
	url, err := CreateInfoURL(t)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var infoResp KeyInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&infoResp); err != nil {
		return nil, err
	}

	return &infoResp, nil
}

func CheckAPIKeyLevel(info KeyInfoResponse, requiredLevel int) (bool, error) {
	if info.Info.Access.Level >= requiredLevel {
		return true, nil
	}
	return false, fmt.Errorf("APIKey does not have the required access level. Key Provided: %d, Required Level: %d", info.Info.Access.Level, requiredLevel)
}
