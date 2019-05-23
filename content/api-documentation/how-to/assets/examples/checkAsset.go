package main

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func init() {
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func main() {
	// Check if AAPL is tradable on the Alpaca platform.
	asset, err := alpaca.GetAsset("AAPL")
	if err != nil {
		fmt.Println("Asset not found for AAPL.")
	} else if asset.Tradable {
		fmt.Println("We can trade AAPL.")
	}
}