---
title: Alpaca Data API v1
weight: 120
---

# Market Data

## Alpaca Data API v1

Alpaca Data API provides the market data available to the client user through
the REST and websocket streaming interfaces. Alpaca Data API consolidates
data sources from five different exchanges.

- IEX (Investors Exchange LLC)
- NYSE National, Inc.
- Nasdaq BX, Inc.
- Nasdaq PSX
- NYSE Chicago, Inc.

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

## Streaming trades, quotes and bars

Alpaca Data API provides websocket streaming for trades,
quotes and minute bars with the same API key. For the details about
the streaming, please see [the reference page]({{<
 relref "#market-data-streaming" >}})


## Polygon Integration

Alpaca integrates with Polygon service, provided for the users who
have funded live trading accounts. For details, please see
[another page]({{< relref "/api-documentation/api-v2/polygon-integration/_index.md" >}})

## Which API should I use?

The biggest difference between Alpaca Data API and Polygon is the
density of the data. While Polygon data is full volume consolidated
from all exchanges in the U.S., Alpaca Data API consists of five
exchanges listed above. That said, Alpaca Data API satisfies most of
the daily uses with enough accuracy as to the real time price needs.

Due to the exchange requirements, it is necessary to have a live trading
account and sign the agreement to start accessing Polygon data.
Users without a live brokerage account must use Alpaca Data API.

For those who have Polygon access via Alpaca API key, Alpaca Data
API still works with the same key, and applies separate constraint
for the concurrent connection. You can also use Alpaca Data API
streaming for the purpose of backup and redundancy.

Alpaca Data API provides more flexible query capability such as multi-symbol and precise time range parameters, in addition to
the real time trade & quotes data. Alpaca's intention is always
to extend the API such that different use cases are easily met.
If you find different needs for API design, please send us
any feedbacks, too.

### Can I get the full volume data from Alpaca Data API?

Currently, no. The plan is to eventually provide everything through Alpaca Data API at some point.
This is to achieve the best user experience for all Alpaca users.
But don't worry! We will have enough time for everyone to smoothly migrate from
Polygon API to Alpaca API. Through our community Slack and email notifications, we will
announce updates to available services in advance of major changes.

## Bars

The bars API provides time-aggregated price and volume data.

{{< rest-endpoint resource="bars" method="GET" path="/v1/bars/{timeframe}" >}}

### Bars Entity

#### Example
{{< rest-entity-example name="bars" >}}

#### Properties
{{< rest-entity-desc name="bars" >}}

## Last Trade

The Last Trade API provides last trade details for a symbol.

{{< rest-endpoint resource="last-trade" method="GET" path="/v1/last/stocks/{symbol}" >}}

### Response

#### Example
{{< rest-entity-example name="last-trade-response" >}}

### Last Trade Entity

#### Properties
{{< rest-entity-desc name="last-trade" >}}

## Last Quote

The Last Quote API provides last quote details for a symbol.

{{< rest-endpoint resource="last-quote" method="GET" path="/v1/last_quote/stocks/{symbol}" >}}

### Response

#### Example
{{< rest-entity-example name="last-quote-response" >}}

### Last Quote Entity

#### Properties
{{< rest-entity-desc name="last-quote" >}}

## Market Data Streaming

Alpaca Data API provides websocket streaming for trades,
quotes and minute bars. This helps receive the most up to date
market information that could help your trading strategy to act
upon certain market movement.

#### Specifications

- Each account can have up to one concurrent websocket connection.
- Trades, quotes and minute bars are supported.
- Subscription is limited to 30 channels at a time for trades and quotes (`T.` and `Q.`). This limit is temporary and we may support more channels in the future.
- There is no limit for the number of channels with minute bars (`AM.`).
- Trades and quotes are from the 5 exchanges.
- Minute bars are also based on the trades in the 5 exchanges.

#### How to connect the websocket streaming

The message protocol is almost the same as trade API stream, but
please note the endpoint is `wss://data.alpaca.markets/stream` (papertrading API key works with this URL, too). At the
beginning of connection, send API key as part of the "authenticate" `action`.

{{< snippet >}}
{
    "action": "authenticate",
    "data": {
        "key_id": "<YOUR_KEY_ID>",
        "secret_key": "<YOUR_SECRET_KEY>"
    }
}
{{< /snippet >}}

Then the server returns the result of authentication.

{{< snippet >}}
{
    "stream": "authorization",
    "data": {
        "action": "authenticate",
        "status": "authorized"
    }
}
{{< /snippet >}}

If the authentication is already done for the connection, it returns error.

{{< snippet >}}
{
  "stream": "listening",
  "data": {
    "error": "your connection is already authenticated"
  }
}
{{< /snippet >}}

If the authentication fails, it also returns error.

{{< snippet >}}
{
  "stream": "authorization",
  "data": {
    "action": "authenticate",
    "status": "unauthorized"
  }
}
{{< /snippet >}}

When another connection is already open, you will receive the error below.

{{< snippet >}}
{
  "stream": "listening",
  "data": {
    "error": "your connection is rejected while another connection is open under the same account"
  }
}
{{< /snippet >}}

Once the authentication is done, send a `listen` message to start
receiving the messages for the channels you want to receive.

{{< snippet >}}
{
    "action": "listen",
    "data": {
        "streams": ["T.SPY", "Q.SPY", "AM.SPY"]
    }
}
{{< /snippet >}}

The stream names can optionally be prefixed by `alpacadatav1/`
{{< snippet >}}
{
    "action": "listen",
    "data": {
        "streams": ["alpacadatav1/T.SPY", "alpacadatav1/Q.SPY", "alpacadatav1/AM.SPY"]
    }
}
{{< /snippet >}}

The server responds with acknowledgement (the prefix is trimmed).

{{< snippet >}}
{
    "stream": "listening",
    "data": {
        "streams": ["T.SPY", "Q.SPY", "AM.SPY"]
    }
}
{{< /snippet >}}

When you want to stop certain channels to stream, send a `unlisten` message.

{{< snippet >}}
{
    "action": "unlisten",
    "data": {
        "streams": ["T.SPY", "Q.SPY"]
    }
}
{{< /snippet >}}

It will again responds with acknowledgement.

{{< snippet >}}
{
    "stream": "listening",
    "data": {
        "streams": ["AM.SPY"]
    }
}
{{< /snippet >}}

The channel names follow the rules below.

- `T.{symbol}`: trades for the symbol
- `Q.{symbol}`: quotes for the symbol
- `AM.{symbol}`: minute bars for the symbol. (`AM.*` will provide all symbols)

#### T = Trade schema:

{{< rest-entity-desc name="stream-trade" >}}

Example: 

{{< rest-entity-example name="stream-trade" >}}

#### Q = Quote schema:

{{< rest-entity-desc name="stream-quote" >}}

Example: 

{{< rest-entity-example name="stream-quote" >}}

#### AM = Bar schema:

{{< rest-entity-desc name="stream-bar" >}}

Example: 

{{< rest-entity-example name="stream-bar" >}}


Below is the example of full message exchanges.

{{< snippet >}}
$ wscat -c wss://data.alpaca.markets/stream
connected (press CTRL+C to quit)
>  {"action": "authenticate","data": {"key_id": "<YOUR_KEY_ID>", "secret_key": "<YOUR_SECRET_KEY>"}}
< {"stream":"authorization","data":{"action":"authenticate","status":"authorized"}}
>  {"action": "listen", "data": {"streams": ["T.SPY"]}}
< {"stream":"listening","data":{"streams":["T.SPY"]}}
{{< /snippet >}}



