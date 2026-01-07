package trackid

import "torngo/internal/api"

type RacingTrackIdRecordsResponse struct {
	Records []struct {
		DriverID    int     `json:"driver_id"`
		DriverName  string  `json:"driver_name"`
		CarItemID   int64   `json:"car_item_id"`
		LapTime     float64 `json:"lap_time"`
		CarItemName string  `json:"car_item_name"`
	} `json:"records"`
}

type RacingTrackIdCategory string

const (
	RacingTrackIdCategoryA RacingTrackIdCategory = "A"
	RacingTrackIdCategoryB RacingTrackIdCategory = "B"
	RacingTrackIdCategoryC RacingTrackIdCategory = "C"
	RacingTrackIdCategoryD RacingTrackIdCategory = "D"
	RacingTrackIdCategoryE RacingTrackIdCategory = "E"
)

type RacingTrackIdRecordsParams struct {
	TrackId   int32
	Category  RacingTrackIdCategory
	Comment   api.ApiComment
	Timestamp api.ApiTimestamp
}
