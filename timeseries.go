package bls

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"errors"
)

// A TimeSeriesParams is a group of parameters passed into GetSeries
type TimeSeriesParams struct {
	SeriesID  []string `json:"seriesid"`
	StartYear int      `json:"startyear"`
	EndYear   int      `json:"endyear"`
}

// The TimeSeriesResponse is the parent object returned from the API
type TimeSeriesResponse struct {
	Results      TimeSeriesResults `json:"Results"`
	Message      []string          `json:"message"`
	ResponseTime int64             `json:"responseTime"`
	Status       string            `json:"status"`
}

// TimeSeriesResults holds the group of series returned
type TimeSeriesResults struct {
	Series []TimeSeries `json:"series"`
}

type DataSlice []TimeSeriesData

// TimeSeries represents each series and it's data
type TimeSeries struct {
	Data     DataSlice `json:"data"`
	SeriesID string    `json:"seriesID"`
}

// TimeSeriesData holds the time based data points in the series
type TimeSeriesData struct {
	Footnotes  []TimeSeriesFootnote `json:"footnotes"`
	Period     string               `json:"period"`
	PeriodName string               `json:"periodName"`
	Value      string               `json:"value"`
	Year       string               `json:"year"`
}

func (ds DataSlice) Less(i, j int) bool {
	if ds[i].Year == ds[j].Year {
		t_i, _ := time.Parse("January", ds[i].PeriodName)
		t_j, _ := time.Parse("January", ds[j].PeriodName)
		return t_i.Month() < t_j.Month()
	}
	return ds[i].Year < ds[j].Year
}

func (ds DataSlice) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

func (ds DataSlice) Len() int {
	return len(ds)
}

// TimeSeriesFootnote is a footnote related to the data point
type TimeSeriesFootnote struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

// GetSeries returns one or more TimeSeries in a TimeSeriesResponse for a given TimeSeriesParams
func GetSeries(params TimeSeriesParams) ([]TimeSeries, error) {
	var series []TimeSeries
	var response TimeSeriesResponse

	url := "http://api.bls.gov/publicAPI/v1/timeseries/data/"

	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return series, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(paramsJSON))
	if err != nil {
		return series, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return series, err
	}

	if len(response.Message) > 0 {
		return series, errors.New(response.Message[0])
	}

	series = response.Results.Series

	return series, nil
}
