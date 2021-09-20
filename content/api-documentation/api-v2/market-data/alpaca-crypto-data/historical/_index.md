---
title: Historical data
weight: 100
summary: Alpaca Crypto Data API provides historical data for more than 5 years.
---

# Historical Data

**Please note that Alpaca Crypto Data is in beta - we welcome any feedback to improve our offering.**

Alpaca provides crypto data from multiple venues and does not route orders to all venues even though it offers data.

List of **crypto** exchanges which are supported by Alpaca.

| Exchange Code | Name of Exchange      |
| ------------- | --------------------- |     
| ERSX          | ErisX                 |
| GNSS          | Genesis               |
| CBSE          | Coinbase              |


## Common behavior

**Base URL**

The Crypto Data API provides historical data through multiple endpoints. These endpoints have the same URL prefix (omitted from now on):

```
https://data.alpaca.markets/v1beta1/crypto
```

This URL is the **same for both subscription plans**, there is no limitation 


**Authentication**
The authentication is done the same way as with the [Trading API](https://alpaca.markets/docs/api-documentation/api-v2/#authentication), simply set the following HTTP headers:

- `APCA-API-KEY-ID`
- `APCA-API-SECRET-KEY`


**Limiting**

Use the `limit` query parameter. The value should be in the range **1 - 10000** (endpoints included) with **1000 being the default** if unspecified.


**Paging**

To support querying long timespans continuously we support paging in our API. If the result you have received contains a `next_page_token` that is **not `null`** there may be more data available in the timeframe you have chosen. Include the token you have received as the `page_token` query parameter for the next request you make while leaving the other parameters unchanged to continue where the previous response left off.


**Ordering**

The results are ordered in ascending order by time.

**Timestamps**

The timestamps for trades, quotes, and bars correspond to when a trade was executed or a quote was generated on the exchange or OTC desk.

## Trades

The Trades API provides historcial trade data for a given crypto symbol on a specified date.

{{< rest-endpoint resource="crypto-trades" method="GET" path="/v1beta1/crypto/{symbol}/trades" useh3="true" >}}


### Response

{{< rest-entity-example name="crypto-trades" >}}


### Properties

{{< rest-entity-desc name="crypto-trades" >}}



## Latest trade

The Latest trade API provides the latest trade data for a given crypto symbol.

{{< rest-endpoint resource="crypto-latest-trade" method="GET" path="/v1beta1/crypto/{symbol}/trades/latest" useh3="true" >}}


### Response

{{< rest-entity-example name="crypto-latest-trade" >}}


### Properties

{{< rest-entity-desc name="crypto-latest-trade" >}}



## Quotes

The Quotes API provides quotes for a given crypto symbol at a specified date.

{{< rest-endpoint resource="crypto-quotes" method="GET" path="/v1beta1/crypto/{symbol}/quotes" useh3="true" >}}


### Response

{{< rest-entity-example name="crypto-quotes" >}}


### Properties

{{< rest-entity-desc name="crypto-quotes" >}}



## Latest quote

The Latest quote API provides the latest quote data for a given ticker symbol.

{{< rest-endpoint resource="crypto-latest-quote" method="GET" path="/v1beta1/crypto/{symbol}/quotes/latest" useh3="true" >}}


### Response

{{< rest-entity-example name="crypto-latest-quote" >}}


### Properties

{{< rest-entity-desc name="crypto-latest-quote" >}}



## Bars

The Bars API returns aggregate historical data for the requested securities.

{{< rest-endpoint resource="crypto-bars" method="GET" path="/v1beta1/crypto/{symbol}/bars" useh3="true" >}}


### Response

{{< rest-entity-example name="crypto-bars" >}}


### Properties

{{< rest-entity-desc name="crypto-bars" >}}


## XBBO

The XBBO API best bid and offer across venues.

{{< rest-endpoint resource="crypto-xbbo" method="GET" path="/v1beta1/crypto/{symbol}/xbbo/latest" useh3="true" >}}


### Response

{{< rest-entity-example name="crypto-xbbo" >}}


### Properties

{{< rest-entity-desc name="crypto-xbbo" >}}




