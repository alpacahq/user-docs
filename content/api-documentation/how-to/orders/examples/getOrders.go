package main

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func init() {
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func main() {
	// Get the last 100 of our closed orders
	status := "closed"
	limit := 100
	closed_orders, err := alpaca.ListOrders(&status, nil, &limit)
	if err != nil {
		panic(err)
	}

	// Get only the closed orders for a particular stock
	var aaplOrders []alpaca.Order
	for _, order := range closed_orders {
		if order.Symbol == "AAPL" {
			aaplOrders = append(aaplOrders, order)
		}
	}
}