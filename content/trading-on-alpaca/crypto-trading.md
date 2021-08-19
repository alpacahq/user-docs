---
title: Crypto Trading
weight: 55
aliases:
    - /crypto-trading.md
---

We now offer crypto trading through our API and the Alpaca web dashboard! Trade all day, seven days a week, as frequently as you'd like. 

## Eligibility

Currently, crypto trading is enabled through an invite-only beta for participants in California, Massachusetts, Missouri, and Montana as well as international users in Hungary. We're working to broaden our eligibility criteria to more states and countries so stay tuned!

## Supported Assets

The list of supported cryptocurrencies includes BTC, BCH, ETH, and LTC and will continue to grow.

Tradable cryptocurrencies can be identified through the [assets API]({{< relref "/api-documentation/api-v2/assets.md" >}}) where the asset entity has `class = crypto` and `tradable = true`.

```json
{
    "id": "bbaf7b16-956e-45dc-96a6-c19f90036ecc",
    "class": "crypto",
    "exchange": "crypto",
    "symbol": "BTCUSD",
    "name": "",
    "status": "active",
    "tradable": true,
    "marginable": false,
    "shortable": false,
    "easy_to_borrow": false,
    "fractionable": true
}
```


## Supported Orders
When submitting crypto orders through the [orders API]({{< relref "/api-documentation/api-v2/orders.md" >}}) and the Alpaca web dashboard, the supported `type` values are `market`, `limit`, `stop_limit` while the supported `time_in_force` values are `day`, `gtc`, `ioc`, `fok`. We accept fractional orders as well with either `notional` or `qty` provided.

Learn more about [orders]({{< relref "/trading-on-alpaca/orders.md" >}}) and [fractional trading]({{< relref "/trading-on-alpaca/fractional-trading.md" >}}). 



## Margin and Short Selling
Currently, crypto is non-marginable and orders are evaluated against `non_marginable_buying_power`. Additionally, crypto is not shortable.

## Market Data
Crypto market data is provided through both the Free and Unlimited data plans.

## Transferring Crypto
At this time, we don't support transfers of crypto in or out of your account.