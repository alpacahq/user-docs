package main

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func init() {
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func main() {
	// Check if the market is open now.
	clock, err := alpaca.GetClock()
	if err != nil {
		panic(err)
	}
	if clock.IsOpen {
		fmt.Println("The market is open.")
	} else {
		fmt.Println("The market is closed.")
	}

	// Check when the market was open on Dec. 1, 2018
	date := "2018-12-01"
	calendar, err := alpaca.GetCalendar(&date, &date)
	if err != nil {
		panic(err)
	}
	calendarDay := calendar[0]
	fmt.Printf("The market opened at %s and closed at %s on %s.\n",
		calendarDay.Open, calendarDay.Close, calendarDay.Date)
}