---
title: Historical data
weight: 100
summary: Alpaca Data API v2 provides three types of historical data. Trades, quotes and bars. To check the multiple endpoints click on the card.
---

# Historical Data

Alpaca Data API v2 provides three types of historical data: [trades]({{<
 relref "#trades" >}}), [quotes]({{<
 relref "#quotes" >}}) and [bars]({{<
 relref "#bars" >}}).


{{< note >}} Alpaca Data API v2 is currently in public beta. 
Please keep in mind that the public beta version may be less stable. {{< /note >}}


## Public beta status

As the Alpaca Data API v2 is currently in public beta, we are still running backfill processes to have a complete 5+ years of historical data available. We ask for your patience until we have everything up and running in the next week or two.

The table below will be updated on a daily basis to show you our data availability:

|  | 2021 | 2020 | 2019 | 2018 | 2017 | 2016 |
| -------- | -------- | -------- | -------- | -------- | -------- | -------- |
| Trades     | Done     | Done     | In progress      | In progress    | In progress     | In progress      |
| Quotes     | Done     | In progress     | In progress      | In progress      | In progress      | In progress      |
| Bars    | In progress       | In progress      | In progress      | In progress     | In progress      | In progress     |


## Common behavior

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


**Limiting**

Use the `limit` query parameter. The value should be in the range **1 - 10000** (endpoints included) with **1000 being the default** if unspecified.


**Paging**

To support querying long timespans continuously we support paging in our API. If the result you have received contains a `next_page_token` that is **not `null`** there may be more data available in the timeframe you have chosen. Include the token you have received as the `page_token` query parameter for the next request you make while leaving the other parameters unchanged to continue where the previous response left off.


**Ordering**

The results are ordered in ascending order by time.


## Trades

The Trades API provides historcial trade data for a given ticker symbol on a specified date.

{{< rest-endpoint resource="trades-v2" method="GET" path="/v2/stocks/{symbol}/trades" useh3="true" >}}


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

{{< rest-endpoint resource="quotes-v2" method="GET" path="/v2/stocks/{symbol}/quotes" useh3="true" >}}


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

{{< rest-endpoint resource="bars-v2" method="GET" path="/v2/stocks/{symbol}/bars" useh3="true" >}}


### Example of one bar

{{< rest-entity-example name="bars-v2-item" >}}


### Properties

{{< rest-entity-desc name="bars-v2-item" >}}


### Example of multiple bars

{{< rest-entity-example name="bars-v2" >}}


### Properties

{{< rest-entity-desc name="bars-v2" >}}
