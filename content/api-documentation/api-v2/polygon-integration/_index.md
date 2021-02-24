---
title: Polygon Integration
weight: 50
---

# Polygon Integration


{{< warning >}} Please note that the Polygon integration will be discontinued from Feb 28, 2021. {{< /warning >}}


All Alpaca customers with live brokerage accounts can access various
kinds of market data provided by [Polygon](https://polygon.io/).
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
$ wscat wss://socket.polygon.io/stocks
[{"ev":"status","status":"connected","message":"Connected Successfully"}]
{"action":"auth","params":"$APCA_API_KEY_ID"}
[{"ev":"status","status":"success","message":"authenticated"}]
{"action":"subscribe","params":"T.MSFT"}
[{"ev":"status","status":"success","message":"subscribed to: T.MSFT"}]
[{"ev":"T","sym":"MSFT","p":114.170,"x":"10","s":467,"t":1577413831740, .... }]
```

Note the Websocket server address for Alpaca users is as follows.

- wss://socket.polygon.io/stocks

For the further description and specification of each REST API endpoint, please find more details [here](https://polygon.io/docs/).
For Websocket specification, please read [Polygon's socket documentation](https://polygon.io/sockets),
as well as refer to the documentation provided by
[each language SDK]({{< relref "/api-documentation/client-sdk/_index.md" >}}).

Polygon is a separate service integrated with Alpaca API authentication.
If you have questions about Polygon API, you can ask questions in the
#dev-polygon channel in [Alpaca community Slack](https://alpaca.markets/slack).