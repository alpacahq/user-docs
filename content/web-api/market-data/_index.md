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
The authentication is done the same way as [Trading API]({{< relref "/web-api/_index.md#authentication" >}}),
and make sure you set the following fields in the HTTP header:

- `APCA-API-KEY-ID`
- `APCA-API-SECRET-KEY`


## Endpoint
The API endpoint is different from the Trading API's endpoint.

```
https://data.alpaca.markets/v1
```

This URL is the same between paper trading and live trading.


## Polygon Integration

All Alpaca customers with full brokerage accounts can access various kinds of market data in [Polygon](https://polygon.io/).
(This data is not available to paper-only accounts.) For the RESTful interface, you need to give the `apiKey`
parameter with the same API key ID you use for Alpaca, as demonstrated below.

```sh
$ curl "https://api.polygon.io/v1/historic/quotes/SPY/2018-06-01?apiKey=$APCA_API_KEY_ID"
```

(Note that your Alpaca API secret key is not required and should not be provided as a parameter in this call.)

Using Polygon, you can query quotes, bars, and fundamentals data from both historical and real-time datasets.
To see everything you have access to through their API, you can view [Polygon's documentation](https://polygon.io/docs/).

With the same API key ID, you can subscribe your algorithm to NATS streaming for market data updates.
In order for your program to be authenticated with the key ID, use the `CONNECT` method with the `auth_token` field.

```
[CONNECT {"auth_token":"<your-key-id>"}]\r\n
```

The NATS server addresses for Alpaca users are shown below.

- nats1.polygon.io:31101
- nats2.polygon.io:31102
- nats3.polygon.io:31103

For the further description and specification of each API endpoint, please find more details [here](https://polygon.io/docs/).
The details about NATS protocol can be found [here](https://nats.io/documentation/internals/nats-protocol/). Also, please
refer to the documentation provided by [each language SDK]({{< relref "/libraries/_index.md" >}}) for its I/O specification.

## Which API should I use?

You may wonder which of the two above - IEX or Polygon - to use. First of all, the full
volume data provided by Polygon is not immediately available to everyone. You first need
to create a live trading account - paper-only accounts must use IEX - and agree with the
exchange agreements.

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
