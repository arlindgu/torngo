package user

import (
	"fmt"
	"net/url"
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
	Cat       UserReportsCategory
	Target    int
	Limit     api.ApiLimit
	Offset    api.ApiOffset
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserReportsParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "cat", string(p.Cat))
	api.SetIfNotZero(q, "target", fmt.Sprintf("%d", p.Target))
	api.SetIfNotZero(q, "limit", string(p.Limit))
	api.SetIfNotZero(q, "offset", string(p.Offset))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetReports(session *api.Session, params *UserReportsParams) (*UserReportsResponse, error) {
	var resp UserReportsResponse
	if err := session.Get("user/reports", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
