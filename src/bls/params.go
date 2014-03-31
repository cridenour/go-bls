package bls

type TimeSeriesParams struct {
	SeriesID  []string `json:"seriesid"`
	StartYear int      `json:"startyear"`
	EndYear   int      `json:"endyear"`
}
