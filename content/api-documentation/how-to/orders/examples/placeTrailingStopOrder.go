package main

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/shopspring/decimal"
)

func init() {
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func main() {
	// Submit a market order to buy 1 share of Apple at market price
	symbol := "AAPL"
	alpaca.PlaceOrder(alpaca.PlaceOrderRequest{
		AssetKey: &symbol,
		Qty: decimal.NewFromFloat(1),
		Side: alpaca.Buy,
		Type: alpaca.Market,
		TimeInForce: alpaca.Day,
	})

	// Submit a trailing stop order to sell 1 share of Apple at a
    // trailing stop of
	symbol = "AAPL"
	alpaca.PlaceOrder(alpaca.PlaceOrderRequest{
		AssetKey: &symbol,
		Qty: decimal.NewFromFloat(1),
		Side: alpaca.Sell,
		Type: alpaca.TrailingStop,
		StopPrice: decimal.NewFromFloat(1.00),  // stop price will be hwm - 1.00$
		TimeInForce: alpaca.Day,
	})

	// Alternatively, you could use trail_percent:
	symbol = "AAPL"
	alpaca.PlaceOrder(alpaca.PlaceOrderRequest{
		AssetKey: &symbol,
		Qty: decimal.NewFromFloat(1),
		Side: alpaca.Sell,
		Type: alpaca.TrailingStop,
		TrailPercent: decimal.NewFromFloat(1.0),  // stop price will be hwm*0.99
		TimeInForce: alpaca.Day,
	})
}