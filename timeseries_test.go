package bls

import (
	"testing"
	"encoding/json"
)

func TestParamsSeriesID(t *testing.T) {
	params := &TimeSeriesParams{[]string{"CWUR0000SA0L1E"}, 2006, 2014}

	if len(params.SeriesID) != 1 {
		t.Errorf("%v has the wrong number of series.", params)
	}

	if params.SeriesID[0] != "CWUR0000SA0L1E" {
		t.Errorf("%v has the wrong series id.", params)
	}
}

func TestParamsStartYear(t *testing.T) {
	params := &TimeSeriesParams{[]string{"CWUR0000SA0L1E"}, 2006, 2014}

	if params.StartYear != 2006 {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsEndYear(t *testing.T) {
	params := &TimeSeriesParams{[]string{"CWUR0000SA0L1E"}, 2006, 2014}

	if params.EndYear != 2014 {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestJsonEncoding(t *testing.T) {
	params := &TimeSeriesParams{[]string{"CWUR0000SA0L1E"}, 2006, 2014}
	_, err := json.Marshal(params)

	if err != nil {
		t.Errorf("%v does not encode into JSON", params)
	}
}
