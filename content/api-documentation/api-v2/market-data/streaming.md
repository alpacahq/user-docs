---
title: Market Data Streaming
weight: 40
---

{{< info-box >}}
Alpaca Data API streaming is currently beta and invite-only.
{{< /info-box >}}

Alpaca Data API provides websocket streaming for trades,
quotes and minute bars. This helps receive the most up to date
market information that could help your trading strategy to act
upon certain market movement.

### Specifications

- Each account can have up to one concurrent websocket connection.
- Trades, quotes and minute bars are supported.
- Subscription is limited to 30 channels at a time for trades and quotes (`T.` and `Q.`). This limit is temporary and we may support more channels in the future.
- There is no limit for the number of channels with minute bars (`AM.`).
- Trades and quotes are from the 5 exchanges.
- Minute bars are also based on the trades in the 5 exchanges.

### How to connect the websocket streaming

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

Once the authentication is done, send a `listen` messsge to start
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

### T = Trade schema:

{{< rest-entity-desc name="stream-trade" >}}

Example: 

{{< rest-entity-example name="stream-trade" >}}

### Q = Quote schema:

{{< rest-entity-desc name="stream-quote" >}}

Example: 

{{< rest-entity-example name="stream-quote" >}}

### AM = Bar schema:

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
