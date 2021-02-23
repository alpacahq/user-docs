---
title: Historical data
weight: 100
summary: Write something else here
---

# Historical Data

Alpaca Data API v2 provides three types of historical data: [Trades]({{<
 relref "#trades" >}}), [quotes]({{<
 relref "#quotes" >}}) and [bars]({{<
 relref "#bars" >}}).

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
