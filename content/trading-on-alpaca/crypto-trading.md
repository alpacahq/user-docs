---
title: Crypto Trading
weight: 55
aliases:
    - /crypto-trading.md
---

We now offer crypto trading through our API and the Alpaca web dashboard! Trade all day, seven days a week, as frequently as you'd like. 

{{< note title="Crypto trading is currently in open beta!" >}}
Sign up for an account [here](https://app.alpaca.markets/signup) to start trading now!
{{< /note >}}

## Eligibility

Currently, crypto trading is enabled through an invite-only beta for participants in California, Massachusetts, Missouri, and Montana as well as international users in Hungary. We're working to broaden our eligibility criteria to more states and countries so stay tuned!

## Supported Assets

The list of supported cryptocurrencies includes BTC, BCH, ETH, and LTC. We
constantly evaluate the list and aim to to grow the number of supported
currencies.

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

Please note that the symbol appears with `USD`, such as `BTCUSD` instead of `BTC`.

## Supported Orders
When submitting crypto orders through the [orders API]({{< relref
"/api-documentation/api-v2/orders.md" >}}) and the Alpaca web dashboard, all the order types, including advanced order, are supported while the supported
`time_in_force` values are `day`, `gtc`, `ioc`, and `fok`. We accept fractional
orders as well with either `notional` or `qty` provided.

Learn more about [orders]({{< relref "/trading-on-alpaca/orders.md" >}}) and [fractional trading]({{< relref "/trading-on-alpaca/fractional-trading.md" >}}). 

All cryptocurrency assets are fractionable but the supported decimal points vary depending on the cryptocurrency.

| Symbol | Minimum Qty  | Qty Increment | Price Increment  |
|--------|--------------|---------------|------------------|
| BTCUSD | 0.001 | 0.0001 | $0.10 |
| BCHUSD | 0.01  | 0.0001 | $0.01 |
| ETHUSD | 0.01  | 0.0001 | $0.01 |
| LTCUSD | 0.1   | 0.0001 | $0.01 |

Note these values could change in the future.

## Margin and Short Selling
Cryptocurrencies are non-marginable. This means that you cannot use leverage to
buy them and orders are evaluated against `non_marginable_buying_power`.

Cryptocurrencies are not shortable.

## Market Data
Crypto market data is provided through both the Free and Unlimited data plans.

## Transferring Crypto
At this time, we don't support transfers of crypto in or out of your account.
You have to fund your Alpaca account in USD first through the regular funding
methods.

## Trading Hours
Crypto trading is offered for 24 hours everyday and your orders will be executed
throughout the day.

----

{{< info-box >}} Cryptocurrency trading is offered through an account with
Alpaca Crypto LLC. Alpaca Crypto LLC is not a member of FINRA or SIPC.
Cryptocurrencies are not stocks and your cryptocurrency investments are not
protected by either FDIC or SIPC. {{< /info-box >}}
