---
title: Alpaca Data API v2
weight: 120
summary: Please note that Alpaca Data API v2 will be live starting on Feb 27, 2021. Alpaca Data API v2 provides market data in 2 two different plans, Basic and Pro. The Basic plan is included in both paper-only and live trading accounts as the default plan for free. 
---

# Market Data

## Overview

Alpaca Data API v2 provides market data through an easy to use HTTP API for historical data and through WebSocket for real-time data.

We provide easy to use SDKs written in Python, Go and NodeJS.

**Please note that Alpaca Data API v2 will be live starting on Feb 27, 2021.**

## Subscription Plans

Alpaca Data API v2 provides market data in 2 two different plans: **Basic** and **Pro**.

The Basic plan is included in both paper-only and live trading accounts as the default plan for free.
You can subscribe to the Pro plan from the Dashbaord, to sign up click [here](https://app.alpaca.markets/signup). 

|  | Basic | Pro |
| -------- | -------- | -------- |
| Pricing    | Free     | $49/mo     |
| Securities coverage    | US Stocks & ETFs     | US Stocks & ETFs    |
| Real-time market coverage    | IEX     | All US Stock Exchanges     |
| Websocket subscriptions   | 30 symbols    | Unlimited    |
| Historical data timefrane   | 5+ years | 5+ years     |
| Historical data delay| 15 minutes | 1 minute
| Historical API calls    | 200/min     | Unlimited |


The **Basic plan** consists data from IEX (Investors Exchange LLC).

For the **Pro plan**, we receive direct feeds from the CTA (administered by NYSE) and UTP (administered by Nasdaq) SIPs. These 2 feeds combined offer 100% market volume. 
Data is received from all 19 exchanges:
* New York Stock Exchange LLC
* NYSE National, Inc.
* NYSE Chicago, Inc.
* NYSE Arca, Inc.
* NYSE American LLC
* Nasdaq Stock Market LLC
* Nasdaq BX, Inc.
* Nasdaq PHLX LLC
* Nasdaq ISE, LLC
* Investors Exchange LLC
* Financial Industry Regulatory Authority, Inc.
* Cboe Exchange, Inc.
* Cboe EDGX Exchange, Inc.
* Cboe EDGA Exchange, Inc.
* Cboe BZX Exchange, Inc.
* Cboe BYX Exchange, Inc.
* Long Term Stock Exchange, Inc.
* Members Exchange (MEMX)
* MIAX Pearl, LLC (MIAX).

Please visit our [Support](https://alpaca.markets/support/) page to learn more about the different plans.

## Historical Data

Alpaca Data API v2 provides three types of historical data: [Trades]({{<
 relref "#trades" >}}), [quotes]({{<
 relref "#quotes" >}}) and [bars]({{<
 relref "#bars" >}}).

### Common behavior

**Base URL**

Alpaca Data API v2 provides historical data through multiple endpoints. These endpoints have the same URL prefix (omitted from now on):

```
https://data.alpaca.markets/v2
```

This URL is the **same for both subscription plans** but users with Basic subscription will receive an error when trying to access data that is too recent.


**Authentication**
The authentication is done the same way as with the [Trading API](https://alpaca.markets/docs/api-documentation/api-v2/#authentication), simply set the following HTTP headers:

- `APCA-API-KEY-ID`
- `APCA-API-SECRET-KEY`


**Parameters**

You will see two kinds of parameters in the API documentation: **path parameters and query parameters**.

Path parameters are documented using the `{parameter}` convention where you will have to replace the bracket and its contents with the value you wish to use. Path parameters are always required.

Query parameters are key value pairs that follow a `?` and the end of the URL. Althoug more difficult to use than path parameters, this [Wikipedia article](https://en.wikipedia.org/wiki/Query_string) should be enough to get you started. Some query parameters are optional while some are required. Look for an indicator (i.e. `*` for required) when you are unsure.


**Limiting**

Depending on your goal you might like to receive more than one data point in the response for a query. Actually, this is almost always true after you're done experimenting with the API. To accomodate this you can include the `limit` query parameter. The value should be in the range **1 - 10000** (including them both) with **1000 being the default** if unspecified.


**Paging**

To support querying long timespans continuously we support paging in our API. If the result you have received contains a `next_page_token` that is **not `null`** there may be more data available in the timeframe you have chosen. Include the token you have received as the `page_token` query parameter for the next request you make while leaving the other parameters unchanged to continue where the previous response left off.


**Ordering**

The results are ordered in ascending order by time.


## Trades

The Trades API provides historcial trade data for a given ticker symbol on a specified date.

{{< rest-endpoint resource="trades-v2" method="GET" path="/v2/stocks/{symbol}/trades" >}}


### Example of one trade

{{< rest-entity-example name="trades-v2-item" >}}


### Properties

{{< rest-entity-desc name="trades-v2-item" >}}


### Example of multiple trades

{{< rest-entity-example name="trades-v2" >}}


### Properties

{{< rest-entity-desc name="trades-v2" >}}


## Quotes

The Quotes API provides NBBO quotes for a given ticker symbol at a specified date.

{{< rest-endpoint resource="quotes-v2" method="GET" path="/v2/stocks/{symbol}/quotes" >}}


### Example of one quote

{{< rest-entity-example name="quotes-v2-item" >}}


### Properties

{{< rest-entity-desc name="quotes-v2-item" >}}


### Example of multiple quotes

{{< rest-entity-example name="quotes-v2" >}}


### Properties

{{< rest-entity-desc name="quotes-v2" >}}


## Bars

The bars API returns aggregate historical data for the requested securities.

{{< rest-endpoint resource="bars-v2" method="GET" path="/v2/stocks/{symbol}/bars" >}}


### Example of one bar

{{< rest-entity-example name="bars-v2-item" >}}


### Properties

{{< rest-entity-desc name="bars-v2-item" >}}


### Example of multiple bars

{{< rest-entity-example name="bars-v2" >}}


### Properties

{{< rest-entity-desc name="bars-v2" >}}


## Real-time data

Alpaca Data API v2 provides websocket streaming for trades, quotes and minute bars. This helps receive the most up to date market information that could help your trading strategy to act upon certain market movements.

Once a connection is established and you have successfully authenticated yourself you can subscribe to trades, quotes and minute bars for a particular symbol or multiple symbols.


### Subscription plans
**Basic plan:**
- You can only connect to IEX data source. One concurrent connection is allowed.
- Subscription is limited to 30 channels at a time for trades (`trades`) and quotes (`quotes`). We may support more channels in the future.
- There is no limit for the number of channels with minute bars (`bars`).
- Minute bars are based on the trades from IEX.

**Pro plan:** 
- There is no limit for the number of channels at a time for trades, quotes and minute bars(`trades`,`quotes` and `bars`).
- Trades, quotes and mintue bars are direct feeds from the CTA (administered by NYSE) and UTP (administered by Nasdaq) SIPs.


### Common behavior


### URL

To access real-time data use the URL below, substituting `iex` or `sip` to `{source}` depending on your subscription.

```
ACTUALURL/{source}
```

Attemption to access a data source not available for your subscription will result in an error during authentication.


### Message format

Every message you receive from the server will be in the format:
```
[{"T": "{message_type}", {contents}},...]
```

Control messages (i.e. where `"T"` is `error`, `success` or `subscription`) always arrive in arrays of size one to make their processing easier.

Data points however may arrive in arrays that have a length that is greater than one. This is to facilitate clients whose connection is not fast enough to handle data points sent one by one. Our server buffers the outgoing messages but slow clients may get disconnected if their buffer becomes full.


### Encoding and compression

Messages over the WebSocket are in encoded as clear text.

To reduce bandwidth requirements we have implemented compression as per [RFC-7692](https://tools.ietf.org/html/rfc7692). Our SDKs handle this for you so in most cases you won't have to implement anything yourself.


### Communication flow

The communication can be thought of as two separate phases: **establishment** and **receiving data**.


### Establishment**

To establish the connection **first you will need to connect** to our server using the URL above.

Upon successfully connecting, you will receive the welcome message:
```
[{"T":"success","msg":"connected"}]
```
Now you will need to **authenticate yourself using** your credentials by sending the following message:
```
{"action": "auth", "key": "{APCA-API-KEY-ID}", "secret": "{APCA-API-SECRET-KEY}"}
```

Please not that each account can have up to one concurrent websocket connection. Subsequent attempts to connect are rejected.

If you provided correct credentials you will receive another `success` message:
```
[{"T":"success","msg":"authenticated"}]
```

### Receiving data

Congratulations, you are ready to receive real-time market data!

You can send one or more subscription messages (described below) and after confirmation you will receive the corresponding market data.

At any time you can subscribe to or unsubscribe from symbols. Please note that due to the internal buffering mentioned above for a short while you may receive data points for symbols you have recently unsubscribed from.


### Message reference - from client to server


### Authentication

After connecting you will have to authenticate as described above.
```
{"action":"auth","key":"PK******************","secret":"***********************************"}
```

### Subscribe

You can subscribe to `trades`, `quotes` and `bars` of a particular symbol (or `*` for evey symbol in the case of `bars`). A `subscribe` message should contain what subscription you want to add to your current subscriptions in your session so you don't have to send what you're already subscribed to.
```
{"action":"subscribe","trades":["AAPL"],quotes":["AMD","CLDR"],"bars":["AAPL","VOO"]}
```
You can also omit either one of them (`trades`,`quotes` or `bars`) if you don't want to subscribe to any symbols in that category but be sure to include at least one of the three.


### Unsubscribe

Much like `subscribe` you can also send an `unsubscribe` message that subtracts the list of subscriptions specified from your current set of subscriptions.

```
{"action":"unsubscribe","trades":["VOO"],quotes":["IBM"],"bars":[]}
```

### Message reference - from server to client


### Control messages

You may receive the following control messages during your session. 
```
[{"T":"success","msg":"connected"}]
```
You have successfully connected to our server.
```
[{"T":"success","msg":"authenticated"}]
```
You have successfully authenticated.


### Errors

You may receive an error during your session. You can differentiate between them using the list below.
```
 [{"T":"error","code":400,"msg":"invalid syntax"}]
```
The message you sent to the server did not follow the specification
```
 [{"T":"error","code":401,"msg":"not authenticated"}]
```
You have attempted to subscribe or unsubscribe before authentication.
```
[{"T":"error","code":402,"msg":"auth failed"}]
```
You have provided invalid authentication credentials.
```
[{"T":"error","code":403,"msg":"already authenticated"}]
```
You have already successfully authenticated during your current session.
```
[{"T":"error","code":404,"msg":"auth timeout"}]
```
You failed to successfully authenticate after connecting. You have a few seconds to authenticate after connecting.
```
[{"T":"error","code":405,"msg":"symbol limit exceeded"}]
```
The symbol subscription request you sent would put you over the limit set by your subscription package. If this happens your symbol subscriptions are the same as they were before you sent the request that failed.
```
[{"T":"error","code":406,"msg":"connection limit exceeded"}]
```
You already have an ongoing authenticated session.
```
[{"T":"error","code":407,"msg":"slow client"}]
```
You may receive this if you are too slow to process the messages sent by the server. Please note that this is not guaranteed to arrive before you are disconnected to avoid keeping slow connections active forever.
```
[{"T":"error","code":408,"msg":"v2 not enabled"}]
```
Your account does not have access to Data v2.
```
[{"T":"error","code":409,"msg":"insufficient subscription"}]
```
You have attempted to access a data source not available in your subscription package.
```
[{"T":"error","code":500,"msg":"internal error"}
```
An unexpected error occurred on our end and we are investigating the issue.

#### Subscription confirmation

After subscribing or unsubscribing you will receive a message that describes your current list of subscriptions.
```
[{"T":"subscription","trades":["AAPL"],"quotes":["AMD","CLDR"],"bars":["IBM","AAPL","VOO"]}]
```
You will always receive your entire list of subscriptions, as illustrated by the sample communication excerpt below:
{{< snippet >}}
 > {"action": "subscribe", "trades": ["AAPL"], "quotes": ["AMD", "CLDR"], "bars": ["*"]}
 < [{"T":"subscription","trades":["AAPL"],"quotes":["AMD","CLDR"],"bars":["*"]}]
 > {"action": "unsubscribe", "bars": ["*"]}
 > [{"T":"subscription","trades":["AAPL"],"quotes":["AMD","CLDR"],"bars":[]}]
{{< /snippet >}}


## Data points


### T = Trade schema:

{{< rest-entity-desc name="stream-trade-v2" >}}

Example: 

{{< rest-entity-example name="stream-trade-v2" >}}


### Q = Quote schema:

{{< rest-entity-desc name="stream-quote-v2" >}}

Example: 

{{< rest-entity-example name="stream-quote-v2" >}}


### AM = Bar schema:

{{< rest-entity-desc name="stream-bar-v2" >}}

Example: 

{{< rest-entity-example name="stream-bar-v2" >}}


### Example

{{< snippet >}}
$ wscat -c wss://REPLACETHISURL
connected (press CTRL+C to quit)
< [{"T":"success","msg":"connected"}]
> {"action": "auth", "key": "*****", "secret": "*****"}
< [{"T":"success","msg":"authenticated"}]
> {"action": "subscribe", "trades": ["AAPL"], "quotes": ["AMD", "CLDR"], "bars": ["*"]}
< [{"T":"t","i":96921,"S":"AAPL","x":"D","p":126.55,"s":1,"t":"2021-02-22T15:51:44.208Z","c":["@","I"],"z":"C"}]
< [{"T":"q","S":"AMD","bx":"U","bp":87.66,"bs":1,"ax":"X","ap":87.67,"as":1,"t":"2021-02-22T15:51:45.3355677Z","c":["R"],"z":"C"},{"T":"q","S":"AMD","bx":"U","bp":87.66,"bs":1,"ax":"Q","ap":87.68,"as":4,"t":"2021-02-22T15:51:45.335689322Z","c":["R"],"z":"C"},{"T":"q","S":"AMD","bx":"U","bp":87.66,"bs":1,"ax":"X","ap":87.67,"as":1,"t":"2021-02-22T15:51:45.335806018Z","c":["R"],"z":"C"}]
{{< /snippet >}}
