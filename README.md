go-bls
======

A golang wrapper around the public BLS API. Expect breaking changes if actively developed.


Upcoming Changes
----------------
* Only return TimeSeries types instead of the full result
* Check for error messages and return expected types and codes


Example
-------

A basic example getting one time series for 2012-2014 and printing them out.


```go
package main

import (
	"fmt"
	bls "github.com/cridenour/go-bls"
)

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	params := bls.TimeSeriesParams{[]string{"EIUIR"}, 2012, 2014}
	data, err := bls.GetSeries(params)
	perror(err)

	if len(data.Message) > 0 {
		for _, message := range data.Message {
			fmt.Printf("Error: %s", message)
		}
	}

	for _, series := range data.Results.Series {
		fmt.Printf("SeriesID: %s\n", series.SeriesID)
		for _, data := range series.Data {
			fmt.Printf("[%s %s] %s\n", data.Year, data.PeriodName, data.Value)
		}
		fmt.Printf("\n")
	}
}
```