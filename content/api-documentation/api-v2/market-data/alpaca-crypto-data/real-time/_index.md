---
title: Real-time data
weight: 120
summary: Alpaca Crypto Data API provides websocket streaming for trades, quotes, minute bars and more. This helps receive the most up to date market information that could help your trading strategy to act upon certain market movements.
 
---

# Real-time data

**Please note that Alpaca Crypto Data is in beta - we welcome any feedback to improve our offering.**


Crypto Data API provides websocket streaming for trades, quotes, minute bars and more. This helps receive the most up to date market information that could help your trading strategy to act upon certain market movements.

Once a connection is established and you have successfully authenticated yourself you can subscribe to tany event for a particular symbol or multiple symbols.

Alpaca provides crypto data from multiple venues and does not route orders to all venues even though it offers data.


## Subscription plans

Both subscription plans provide the same source and level of crypto data, the only difference on real-time data is how many symbols (channels) you can listen to.

**Free plan:**
- One concurrent connection is allowed.
- Subscription is limited to 30 channels at a time for trades (`trades`) and quotes (`quotes`).
- There is no limit for the number of channels with minute bars (`bars`).


**Unlimited plan:** 
- One concurrent connection is allowed.
- There is no limit for the number of channels at a time for trades, quotes and minute bars(`trades`,`quotes` and `bars`).


## Common behavior


### URL

To access real-time data use the URL below.

```
wss://stream.data.alpaca.markets/v1beta1/crypto

```



### Message format

Every message you receive from the server will be in the format:
```
[{"T": "{message_type}", {contents}},...]
```

Control messages (i.e. where `"T"` is `error`, `success` or `subscription`) always arrive in arrays of size one to make their processing easier.

Data points however may arrive in arrays that have a length that is greater than one. This is to facilitate clients whose connection is not fast enough to handle data points sent one by one. Our server buffers the outgoing messages but slow clients may get disconnected if their buffer becomes full.


### Encoding and compression

Messages over the websocket are in encoded as clear text.

To reduce bandwidth requirements we have implemented compression as per [RFC-7692](https://tools.ietf.org/html/rfc7692). Our SDKs handle this for you so in most cases you won't have to implement anything yourself.

### Timestamps

The timestamps for trades, quotes, and bars correspond to when a trade was executed or a quote was generated on the exchange or OTC desk.


## Communication flow

The communication can be thought of as two separate phases: **establishment** and **receiving data**.


### Establishment

To establish the connection **first you will need to connect** to our server using the URL above.

Upon successfully connecting, you will receive the welcome message:
```
[{"T":"success","msg":"connected"}]
```
Now you will need to **authenticate yourself using** your credentials by sending the following message:
```
{"action": "auth", "key": "{APCA-API-KEY-ID}", "secret": "{APCA-API-SECRET-KEY}"}
```

Please note that each account can have up to one concurrent websocket connection. Subsequent attempts to connect are rejected.

If you provided correct credentials you will receive another `success` message:
```
[{"T":"success","msg":"authenticated"}]
```

### Receiving data

Congratulations, you are ready to receive real-time crypto market data!

You can send one or more subscription messages (described [below]({{<
 relref "#subscribe" >}})) and after confirmation you will receive the corresponding market data.

At any time you can subscribe to or unsubscribe from symbols. Please note that due to the internal buffering mentioned above for a short while you may receive data points for symbols you have recently unsubscribed from.


## Client to server


### Authentication

After connecting you will have to authenticate as described above.
```
{"action":"auth","key":"PK******************","secret":"***********************************"}
```

### Subscribe

You can subscribe to `trades`, `quotes` and `bars` of a particular symbol (or `*` for every symbol in the case of `bars`). A `subscribe` message should contain what subscription you want to add to your current subscriptions in your session so you don't have to send what you're already subscribed to.
```
{"action":"subscribe","trades":["BTCUSD"],"quotes":["LTCUSD","ETHUSD"],"bars":["BCHUSD"]}
```
You can also omit either one of them (`trades`,`quotes` or `bars`) if you don't want to subscribe to any symbols in that category but be sure to include at least one of the three.

To subscribe to multiple exchanges' datafeed use `exchanges` parameter with comma seaparated delisting.
Example `wss://stream.data.alpaca.markets/v1beta1/crypto?exchanges=ERSX,GNSS`


### Unsubscribe

Much like `subscribe` you can also send an `unsubscribe` message that subtracts the list of subscriptions specified from your current set of subscriptions.
```
{"action":"unsubscribe","trades":["BTCUSD"],"quotes":["ETHUSD"],"bars":[]}
```

## Server to client


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
 > {"action": "subscribe", "trades": ["BTCUSD"], "quotes": ["BTCUSD", "LTCUSD"], "bars": ["*"]}
 < [{"T":"subscription","trades":["BTCUSD"],"quotes":["BTCUSD","LTCUSD"],"bars":["*"]}]
 > {"action": "unsubscribe", "bars": ["*"]}
 > [{"T":"subscription","trades":["BTCUSD"],"quotes":["BTCUSD","LTCUSD"],"bars":[]}]
{{< /snippet >}}



### Example

{{< snippet >}}
$ wscat -c wss://stream.data.alpaca.markets/v2/sip
connected (press CTRL+C to quit)
< [{"T":"success","msg":"connected"}]
> {"action": "auth", "key": "*****", "secret": "*****"}
< [{"T":"success","msg":"authenticated"}]
> {"action": "subscribe", "bars": ["BTCUSD"]}
< [{"T":"b","S":"BTCUSD","x":"CBSE","o":47993.71,"c":47982.6,"h":48000,"l":47976.69,"v":1.66283276,"t":"2021-09-17T02:02:00Z","n":162,"vw":47984.1688165374}]
{{< /snippet >}}
