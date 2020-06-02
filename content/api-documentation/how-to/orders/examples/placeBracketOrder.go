package main

import (
    "github.com/alpacahq/alpaca-trade-api-go/alpaca"
    "github.com/alpacahq/alpaca-trade-api-go/common"
    "github.com/shopspring/decimal"
)

func init() {
    API_KEY := "YOUR_API_KEY_HERE"
    API_SECRET := "YOUR_API_SECRET_HERE"
    BASE_URL := "https://paper-api.alpaca.markets"
    // Check for environment variables
    if common.Credentials().ID == "" {
        os.Setenv(common.EnvApiKeyID, API_KEY)
    }
    if common.Credentials().Secret == "" {
        os.Setenv(common.EnvApiSecretKey, API_SECRET)
    }
    alpaca.SetBaseUrl(BASE_URL)
}

func main() {
    // Submit a limit order to buy 1 share of Apple and add
    // StopLoss and TakeProfit orders.
    client := alpaca.NewClient(common.Credentials())
    symbol := "AAPL"
    tpp := decimal.NewFromFloat(320.)
    spp := decimal.NewFromFloat(317.)
    limit := decimal.NewFromFloat(318.)
    tp := &alpaca.TakeProfit{LimitPrice: &tpp}
    sl := &alpaca.StopLoss{
        LimitPrice: nil,
        StopPrice:  &spp,
    }
    req := alpaca.PlaceOrderRequest{
        AccountID:   common.Credentials().ID,
        AssetKey:    &symbol,
        Qty:         decimal.New(1, 0),
        Side:        alpaca.Buy,
        LimitPrice:  &limit,
        TimeInForce: alpaca.GTC,
        Type:        alpaca.Limit,
        OrderClass:  alpaca.Bracket,
        TakeProfit:  tp,
        StopLoss:    sl,
    }
    order, err := client.PlaceOrder(req)
    fmt.Println(order)
    fmt.Println(err)

    // We could buy a position and just add a stop loss (OTO Orders)
    spp := decimal.NewFromFloat(317.)
    limit := decimal.NewFromFloat(318.)
    sl := &alpaca.StopLoss{
        StopPrice:  &spp,
    }
    req := alpaca.PlaceOrderRequest{
        AccountID:   common.Credentials().ID,
        AssetKey:    &symbol,
        Qty:         decimal.New(1, 0),
        Side:        alpaca.Buy,
        LimitPrice:  &limit,
        TimeInForce: alpaca.GTC,
        Type:        alpaca.Limit,
        OrderClass:  alpaca.Oto,
        StopLoss:    sl,
    }
    order, err := client.PlaceOrder(req)
    fmt.Println(order)
    fmt.Println(err)

    // We could split it to 2 orders. first buy a stock,
    // and then add the stop/profit prices (OCO Orders)
    limit := decimal.NewFromFloat(318.)
    req := alpaca.PlaceOrderRequest{
        AccountID:   common.Credentials().ID,
        AssetKey:    &symbol,
        Qty:         decimal.New(1, 0),
        Side:        alpaca.Buy,
        LimitPrice:  &limit,
        TimeInForce: alpaca.GTC,
        Type:        alpaca.Limit,
        OrderClass:  alpaca.Simple,
    }
    order, err := client.PlaceOrder(req)
    fmt.Println(order)
    fmt.Println(err)

    // wait for it to complete and then
    tpp := decimal.NewFromFloat(320.)
    spp := decimal.NewFromFloat(317.)
    tp := &alpaca.TakeProfit{LimitPrice: &tpp}
    sl := &alpaca.StopLoss{
        LimitPrice: nil,
        StopPrice:  &spp,
    }
    req := alpaca.PlaceOrderRequest{
        AccountID:   common.Credentials().ID,
        AssetKey:    &symbol,
        Qty:         decimal.New(1, 0),
        Side:        alpaca.Sell,
        TimeInForce: alpaca.GTC,
        Type:        alpaca.Limit,
        OrderClass:  alpaca.Oco,
        TakeProfit:  tp,
        StopLoss:    sl,
    }
    order, err := client.PlaceOrder(req)
    fmt.Println(order)
    fmt.Println(err)
}
