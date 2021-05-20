package main

import (
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v2"
)

func init() {
	alpaca.SetBaseUrl("https://data.alpaca.markets")
}

func main() {
	bars := alpaca.GetBars(
		"AAPL", v2.Day, v2.Raw, time.Now().Add(-7*24*time.Hour), time.Now().Add(-20*time.Minute), 7)
	var barset []v2.Bar

	for bar := range bars {
		if bar.Error != nil {
			panic(bar.Error)
		}
		barset = append(barset, bar.Bar)
	}

	// See how much AAPL moved in that timeframe.
	startPrice := barset[0].Open
	endPrice := barset[len(barset)-1].Close
	percentChange := ((endPrice - startPrice) / startPrice ) * 100
	fmt.Printf("AAPL moved %v%% over the last 7 days.", percentChange)
}