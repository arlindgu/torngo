package user

import (
	"fmt"
	"net/url"
	"torngo/internal/api"
)

type UserPropertiesResponse struct {
	Properties []struct {
		ID    int64 `json:"id"`
		Owner struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"owner"`
		Property struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"property"`
		Happy  int `json:"happy"`
		Upkeep struct {
			Property int `json:"property"`
			Staff    int `json:"staff"`
		} `json:"upkeep"`
		MarketPrice   int64    `json:"market_price"`
		Modifications []string `json:"modifications"`
		Staff         []struct {
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"staff"`
		UsedBy []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"used_by,omitempty"`
		Status                string `json:"status,omitempty"`
		Cost                  int64  `json:"cost,omitempty"`
		CostPerDay            int64  `json:"cost_per_day,omitempty"`
		RentalPeriod          int    `json:"rental_period,omitempty"`
		RentalPeriodRemaining int    `json:"rental_period_remaining,omitempty"`
		RentedBy              struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"rented_by,omitempty"`
		LeaseExtension struct {
			Cost      int64 `json:"cost"`
			Period    int   `json:"period"`
			CreatedAt int   `json:"created_at"`
		} `json:"lease_extension,omitempty"`
		RenterAsked struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"renter_asked,omitempty"`
	} `json:"properties"`
	Metadata struct {
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"_metadata"`
}

type UserPropertiesFilter string

const (
	PropertiesOwnedByUser   UserPropertiesFilter = "ownedByUser"
	PropertiesOwnedBySpouse UserPropertiesFilter = "ownedBySpouse"
)

type UserPropertiesParams struct {
	Filter    UserPropertiesFilter
	Offset    api.ApiOffset
	Limit     api.ApiLimit
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}

func (p *UserPropertiesParams) EncodeParams() url.Values {
	q := url.Values{}
	api.SetIfNotZero(q, "filter", string(p.Filter))
	api.SetIfNotZero(q, "offset", int(p.Offset))
	api.SetIfNotZero(q, "limit", int(p.Limit))
	api.SetIfNotZero(q, "timestamp", fmt.Sprintf("%d", p.Timestamp))
	api.SetIfNotZero(q, "comment", string(p.Comment))
	return q
}

func GetProperties(session *api.Session, params *UserPropertiesParams) (*UserPropertiesResponse, error) {
	var resp UserPropertiesResponse
	if err := session.Get("user/properties", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
