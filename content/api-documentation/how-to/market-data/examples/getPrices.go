package main

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func init() {
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func main() {
	// Get daily price data for AAPL over the last 5 trading days.
	barCount := 5
	bars, err := alpaca.GetSymbolBars("AAPL", alpaca.ListBarParams{
		Timeframe: "day",
		Limit: &barCount,
	})
	if err != nil {
		panic(err)
	}

	// See how much AAPL moved in that timeframe.
	startPrice := bars[0].Open
	endPrice := bars[len(bars)-1].Close
	percentChange := (startPrice - endPrice) / startPrice
	fmt.Printf("AAPL moved %v%% over the last 5 days.", percentChange)
}