package user

import (
	"fmt"
	"net/url"
	"path"
	"torngo/internal/api"
)

type UserReportsResponse struct {
	Reports []struct {
		Type       string `json:"type"`
		TargetID   int    `json:"target_id"`
		ReporterID int    `json:"reporter_id"`
		FactionID  int    `json:"faction_id"`
		Timestamp  int    `json:"timestamp"`
		Report     struct {
			Money int `json:"money"`
		} `json:"report"`
	} `json:"reports"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserReportsCategory string

const (
	UserReportsCategoryMostWanted        UserReportsCategory = "mostwanted"
	UserReportsCategoryMoney             UserReportsCategory = "money"
	UserReportsCategoryStats             UserReportsCategory = "stats"
	UserReportsCategoryReferences        UserReportsCategory = "references"
	UserReportsCategoryFriendOrFoe       UserReportsCategory = "friendorfoe"
	UserReportsCategoryCompanyFinancials UserReportsCategory = "companyfinancials"
	UserReportsCategoryTrueLevel         UserReportsCategory = "truelevel"
	UserReportsCategoryStockAnalysis     UserReportsCategory = "stockanalysis"
	UserReportsCategoryAnonymousBounties UserReportsCategory = "anonymousbounties"
	UserReportsCategoryInvestment        UserReportsCategory = "investment"
)

type UserReportsParams struct {
	Cat    UserReportsCategory
	Target int
	Limit  int
	Offset int
	api.BaseParams
}

func CreateReportsURL(t *UserReportsParams) (string, error) {
	if t.APIKey == "" {
		return "", fmt.Errorf("APIKey is required")
	}

	u, err := url.Parse(api.APIBaseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, "user", "reports")

	q := u.Query()
	q.Set("key", t.APIKey)

	if t.Comment != "" {
		q.Set("comment", t.Comment)
	}

	if t.Timestamp != 0 {
		q.Set("timestamp", fmt.Sprintf("%d", t.Timestamp))
	}

	if t.Cat != "" {
		q.Set("category", string(t.Cat))
	}

	if t.Target != 0 {
		q.Set("target", fmt.Sprintf("%d", t.Target))
	}

	if t.Offset != 0 {
		q.Set("offset", fmt.Sprintf("%d", t.Offset))
	}

	api.SetLimitParams(q, t.Limit, 20)

	u.RawQuery = q.Encode()
	return u.String(), nil
}
