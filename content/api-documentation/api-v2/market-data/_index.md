---
title: Market Data
weight: 120
---

# Market Data

## Alpaca Data API

Alpaca Data API provides the market data available to the client user through
the REST and websocket streaming interfaces. Currently, the supported data
set is IEX price data, but we are working on adding more.


## Authentication
The authentication is done the same way as [Trading API]({{< relref "/api-documentation/api-v2/_index.md#authentication" >}}),
and make sure you set the following fields in the HTTP header:

- `APCA-API-KEY-ID`
- `APCA-API-SECRET-KEY`


## Endpoint
**Please note the Data API endpoint is different from the Trading API endpoint. Although the Trading API documented herein
references v2, the Data API endpoint should still point to v1 as follows:**

```
https://data.alpaca.markets/v1
```

This URL is the same between paper trading and live trading.


## Polygon Integration

All Alpaca customers with live brokerage accounts can access various kinds of market data in [Polygon](https://polygon.io/).
(This data is not available to users who have not yet set up live accounts.) For the RESTful interface,
you need to give the `apiKey` parameter with the same API key ID you use for Alpaca, as demonstrated below.
You will see an authentication error message saying "invalid API key" if your account does not have Polygon access.

```sh
$ curl "https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/2019-01-01/2019-02-01?apiKey=$APCA_API_KEY_ID"
```

(Note that your Alpaca API secret key is not required and should not be provided as a parameter in this call.)

Using Polygon, you can query quotes, bars, and fundamentals data from both historical and real-time datasets.
To see everything you have access to through their API, you can view [Polygon's documentation](https://polygon.io/docs/).

With the same API key ID, you can subscribe your algorithm to Polygon's Websocket streaming for market data updates.
Put your API key ID as the `params` value in the `auth` action at the beginning
of the communication as the example below.

```
$ wscat wss://alpaca.socket.polygon.io/stocks
[{"ev":"status","status":"connected","message":"Connected Successfully"}]
{"action":"auth","params":"$APCA_API_KEY_ID"}
[{"ev":"status","status":"success","message":"authenticated"}]
{"action":"subscribe","params":"T.MSFT"}
[{"ev":"status","status":"success","message":"subscribed to: T.MSFT"}]
[{"ev":"T","sym":"MSFT","p":114.170,"x":"10","s":467,"t":1577413831740, .... }]
```

The Websocket server address for Alpaca users is as follows.

- wss://alpaca.socket.polygon.io/stocks

For the further description and specification of each REST API endpoint, please find more details [here](https://polygon.io/docs/).
For Websocket specification, please read [Polygon's socket documentation](https://polygon.io/sockets),
as well as refer to the documentation provided by
[each language SDK]({{< relref "/api-documentation/client-sdk/_index.md" >}}).

## Which API should I use?

You may wonder which of the two above - IEX or Polygon - to use. First of all, the full
volume data provided by Polygon is not immediately available to everyone. You first need
to create a live trading account - users without live brokerage accounts must use IEX -\
and agree with the exchange agreements.

<object type="image/svg+xml" data="flowchart.svg">
</object>

If what you need is only daily bars, there is not much difference between the IEX
and the full volume Polygon data. However, for the intraday real-time bars, the IEX data
provides price information coming only from the IEX exchange trades,
while the full volume data reflects trades from all exchanges.

Aside from the actual data contents, Alpaca Data API, which uses IEX, provides more
flexible query capability such as multi-symbol and precise time range parameters.

### Can I get the full volume data from Alpaca Data API?

Currently, no. The plan is to eventually provide everything through Alpaca Data API at some point.
This is to achieve the best user experience for all Alpaca users.
But don't worry! We will have enough time for everyone to smoothly migrate from
Polygon API to Alpaca API. Through our community Slack and email notifications, we will
announce updates to available services in advance of major changes.
