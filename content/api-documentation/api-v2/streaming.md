---
title: Streaming
weight: 40
---

## Overview
Alpaca's API offers WebSocket streaming for account and order updates which follows the [RFC6455 WebSocket protocol](https://tools.ietf.org/html/rfc6455).

To connect to the WebSocket follow the standard opening handshake as defined by the RFC specification to the `/stream`
endpoint of Alpaca's API. Alpaca's streaming service supports both JSON and MessagePack codecs.

Once the connection is authorized, the client can listen to one or more streams
to get updates on particular changes.  These are the streams the client can
choose to listen to.

- account_updates
- trade_updates

The details of each stream will be described later in this document.

In order to listen to streams, the client sends a `listen` message
to the server as follows.
```
{
    "action": "listen",
    "data": {
        "streams": ["trade_updates"]
    }
}
```

The server acknowledges by replying a message in the `listening` stream.

```
{
    "stream": "listening",
    "data": {
        "streams": ["trade_updates"]
    }
}
```

If some of the requested streams are not available, they will not appear
in the `streams` list in the acknowledgement.
Note that the `streams` field in the listen message is to tell
the set of streams to listen, so if you want to stop receiving
updates from the stream, you must send an empty list of streams
values as a listen message.  Similarly, if you want to add more
streams to get updates in addition to the ones you are already
doing so, you must send all the stream names not only the new
ones.

In order to maintain the state of their brokerage accounts at Alpaca, along with requesting from the REST API, clients can also
listen to the trade streams for their accounts. This will ensure any running algorithms will always have the most up-to-date
picture of any accounts they are trading with at Alpaca.

## Authentication
The WebSocket client can be authenticated using the same API key when making HTTP requests. Upon connecting to the WebSocket
client must send an authentication message over the WebSocket connection with the API key, and secret key as its payload:
```
{
    "action": "authenticate",
    "data": {
        "key_id": "{YOUR_API_KEY_ID}",
        "secret_key": "{YOUR_API_SECRET_KEY}"
    }
}
```

The server will then authorize the connection and respond with either a successful response:

```
{
    "stream": "authorization",
    "data": {
        "status": "authorized",
        "action": "authenticate"
    }
}
```

or an unathorized response:

```
{
    "stream": "authorization",
    "data": {
        "status": "unauthorized",
        "action": "authenticate"
    }
}
```

In the case the socket connection is not authorized yet, a new message under
the `authorization` stream is issued in response to the listen request.

```
{
    "stream": "authorization",
    "data": {
        "status": "unauthorized",
        "action": "listen"
    }
}
```

## Order Updates
Updates with regards to orders placed at Alpaca are dispatched over the WebSocket connection under the event `trade_updates`, and include
any data pertaining to orders that are executed with Alpaca. This includes order fills, partial fills, as well as cancellations and
rejections of orders. Clients may listen to this stream by sending a listen message:

```
{
    "action": "listen",
    "data": {
        "streams": ["trade_updates"]
    }
}
```

Any listen messages received by the server will be acknowledged via a message on the `listening` stream. The message's
data payload will include the list of streams the client is currently listening to:

```
{
    "stream": "listening",
    "data": {
        "streams": ["trade_updates"]
    }
}
```

The fields present in a message sent over the `trade_updates` stream depend on the type of event they are communicating. All messages
contain an `event` type and an `order` field, which is the same as the [order object](https://docs.alpaca.markets/api-documentation/api-v2/orders/#properties) that is returned from the REST API.
Potential event types and additional fields that will be in their messages are listed below.

### Common events:
These are the events that are the expected results of actions you may have taken by sending API requests.

- `new`: Sent when an order has been routed to exchanges for execution.
- `fill`: Sent when your order has been completely filled.
    - *timestamp*: The time at which the order was filled.
    - *price*: The average price per share at which the order was filled.
    - *position_qty*: The size of your total position, after this fill event, in shares. Positive for long positions, negative for short positions.
- `partial_fill`: Sent when a number of shares less than the total remaining quantity on your order has been filled.
    - *timestamp*: The time at which the shares were filled.
    - *price*: The average price per share at which the shares were filled.
    - *position_qty*: The size of your total position, after this fill event, in shares. Positive for long positions, negative for short positions.
- `canceled`: Sent when your requested cancelation of an order is processed.
    - *timestamp*: The time at which the order was canceled.
- `expired`: Sent when an order has reached the end of its lifespan, as determined by the order's time in force value.
    - *timestamp*: The time at which the order expired.
- `done_for_day`: Sent when the order is done executing for the day, and will not receive further updates until the next trading day.
- `replaced`: Sent when your requested replacement of an order is processed.
    - *timestamp*: The time at which the order was replaced.

### Rarer events:
These are events that may rarely be sent due to unexpected circumstances on the exchanges. It is unlikely you will need to design your
code around them, but you may still wish to account for the possibility that they will occur.

- `rejected`: Sent when your order has been rejected.
    - *timestamp*: The time at which the rejection occurred.
- `pending_new`: Sent when the order has been received by Alpaca and routed to the exchanges, but has not yet been accepted for execution.
- `stopped`: Sent when your order has been stopped, and a trade is guaranteed for the order, usually at a stated price or better, but has not yet occurred.
- `pending_cancel`: Sent when the order is awaiting cancelation. Most cancelations will occur without the order entering this state.
- `pending_replace`: Sent when the order is awaiting replacement.
- `calculated`: Sent when the order has been completed for the day - it is either "filled" or "done_for_day" - but remaining settlement calculations are still pending.
- `suspended`: Sent when the order has been suspended and is not eligible for trading.
- `order_replace_rejected`: Sent when the order replace has been rejected.
- `order_cancel_rejected`: Sent when the order cancel has been rejected.

### Example
An example message sent over the `trade_updates` stream would look like:
```
{
    "stream": "trade_updates",
    "data": {
        "event": "fill",
        "price": "179.08",
        "timestamp": "2018-02-28T20:38:22Z",
        "position_qty": "100",
        "order": {
            "id": "7b7653c4-7468-494a-aeb3-d5f255789473",
            "client_order_id": "7b7653c4-7468-494a-aeb3-d5f255789473",
            "asset_id": "904837e3-3b76-47ec-b432-046db621571b",
            "symbol": "AAPL",
            "exchange": "NASDAQ",
            "asset_class": "us_equity",
            "side": "buy",
            ...
        }
    }
}
```

## Account Updates
Users may also listen to the account updates stream under: `account_updates`. This stream provides clients with updates pertaining
to their brokerage accounts at Alpaca, including balance information. The account updates stream can be listened to in the same
way as the trade updates stream, and in fact, both streams can be listened to simultaneously:

```
{
    "action": "listen",
    "data": {
        "streams": ["account_updates", "trade_updates"]
    }
}
```

It is highly recommended that clients listen to both streams when using Alpaca's streaming API. Any time there is a state change to the listening user's account, an update is sent over the WebSocket:
```
{
    "stream": "account_updates",
    "data": {
        "id": "ef505a9a-2f3c-4b8a-be95-6b6f185f8a03",
        "created_at": "2018-02-26T19:22:31Z",
        "updated_at": "2018-02-27T18:16:24Z",
        "deleted_at": null,
        "status": "ACTIVE",
        "currency": "USD",
        "cash": "1241.54",
        "cash_withdrawable": "523.71"
    }
}
```
