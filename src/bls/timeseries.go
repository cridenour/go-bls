package bls

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TimeSeriesParams struct {
	SeriesID  []string `json:"seriesid"`
	StartYear int      `json:"startyear"`
	EndYear   int      `json:"endyear"`
}

type TimeSeriesResponse struct {
	Results      TimeSeriesResults           `json:"Results"`
	Message      []TimeSeriesResponseMessage `json:"message"`
	ResponseTime int64                       `json:"responseTime"`
	Status       string                      `json:"status"`
}

type TimeSeriesResults struct {
	Series []TimeSeries `json:"series"`
}

type TimeSeries struct {
	Data     []TimeSeriesData `json:"data"`
	SeriesID string           `json:"seriesID"`
}

type TimeSeriesData struct {
	Footnotes  []TimeSeriesFootnote `json:"footnotes"`
	Period     string               `json:"period"`
	PeriodName string               `json:"periodName"`
	Value      string               `json:"value"`
	Year       string               `json:"year"`
}

type TimeSeriesFootnote struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type TimeSeriesResponseMessage struct {}

func GetSeries(params TimeSeriesParams) (TimeSeriesResponse, error) {
	var series TimeSeriesResponse

	url := "http://api.bls.gov/publicAPI/v1/timeseries/data/"

	params_json, err := json.Marshal(params)
	if err != nil {
		return series, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(params_json))
	if err != nil {
		return series, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return series, err
	}

	err = json.Unmarshal(body, &series)
	if err != nil {
		return series, err
	}

	return series, nil
}
