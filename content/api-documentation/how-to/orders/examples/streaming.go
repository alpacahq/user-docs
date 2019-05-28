package main

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
    "github.com/alpacahq/alpaca-trade-api-go/stream"
)

func init() {
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func main() {
	if err := stream.Register(alpaca.TradeUpdates, tradeHandler); err != nil {
        panic(err)
    }
}

func tradeHandler(msg interface{}) {
	tradeupdate := msg.(alpaca.TradeUpdate)
	if tradeupdate.Order.ClientOrderID == "my_client_order_id" {
		fmt.Printf("Update for {}. Event: {}.\n", tradeupdate.Order.ClientOrderID, tradeupdate.Event)
	}
}