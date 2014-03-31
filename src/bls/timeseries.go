package bls

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TimeSeries struct {
	Results struct {
		Series []struct {
		Data []struct {
		Footnotes []struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"footnotes"`
		Period     string `json:"period"`
		PeriodName string `json:"periodName"`
		Value      string `json:"value"`
		Year       string `json:"year"`
	} `json:"data"`
		SeriesID string `json:"seriesID"`
	} `json:"series"`
	} `json:"Results"`
	Message      []interface{} `json:"message"`
	ResponseTime int64         `json:"responseTime"`
	Status       string        `json:"status"`
}

func GetSeries(params TimeSeriesParams) (series *TimeSeries, err error) {
	url := "http://api.bls.gov/publicAPI/v1/timeseries/data/"

	params_json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(params_json))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	series = new(TimeSeries)
	err = json.Unmarshal(body, &series)
	if err != nil {
		return nil, err
	}

	return
}
