---
title: Alpaca Data API v1
weight: 120
summary: Alpaca Data API v1 provides the market data available to the client user through the REST and websocket streaming interfaces and it consolidates data sources from five different exchanges.
---

# Market Data


**The Data API v1 will be deprecated on August 26th, 2021, so please make sure to update any existing systems to use the new endpoint before that date.**

## Overview

Alpaca Data API v1 provides the market data available to the client user through
the REST and websocket streaming interfaces. The Alpaca Data API v1 consolidates
data sources from a single exchange.

- IEX (Investors Exchange LLC)


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
 relref "/api-documentation/api-v2/market-data/alpaca-data-api-v1/streaming.md" >}})


### Can I get the full volume data from Alpaca?

Yes, the Unlimited plan of Alpaca Data API v2 provides full volume data.
