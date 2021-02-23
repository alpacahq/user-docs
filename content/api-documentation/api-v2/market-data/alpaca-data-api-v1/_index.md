---
title: Alpaca Data API v1
weight: 120
summary: Please note that Alpaca Data API v1 will be deprecated from Feb 28, 2021. Alpaca Data API v1 provides the market data available to the client user through the REST and websocket streaming interfaces and it consolidates data sources from five different exchanges.
---

# Market Data

## Overview

Alpaca Data API v1 provides the market data available to the client user through
the REST and websocket streaming interfaces. Alpaca Data API consolidates
data sources from five different exchanges.

- IEX (Investors Exchange LLC)
- NYSE National, Inc.
- Nasdaq BX, Inc.
- Nasdaq PSX
- NYSE Chicago, Inc.

{{< warning >}} Please note that Alpaca Data API v1 will be deprecated from Feb 28, 2021. {{< /warning >}}


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
the streaming, plase see [the reference page]({{<
 relref "/api-documentation/api-v2/market-data/alpaca-data-api-v1/streaming.md" >}})


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
