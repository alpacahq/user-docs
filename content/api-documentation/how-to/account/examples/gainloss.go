package main

import (
	"fmt"
	"log"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func main() {
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")

	// Get account information.
	account, err := alpaca.GetAccount()
	if err != nil {
		log.Fatalln(err)
	}

	// Calculate the difference between current balance and balance at the last market close.
	balanceChange := account.Equity.Sub(account.LastEquity)

	fmt.Println("Today's portfolio balance change:", balanceChange)
}
