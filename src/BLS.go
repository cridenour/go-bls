package main

import (
	"fmt"
	"bls"
)

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	params := bls.TimeSeriesParams{[]string{"EIUIR"}, 2002, 2014}
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
