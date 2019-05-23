package main

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func init() {
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func main() {
	// Get our account information.
	account, err := alpaca.GetAccount()
	if err != nil {
		panic(err)
	}

	// Check if our account is restricted from trading.
	if account.TradingBlocked {
		fmt.Println("Account is currently restricted from trading.")
	}

	// Check how much money we can use to open new positions.
	fmt.Printf("%v is available as buying power.\n", account.BuyingPower)
}