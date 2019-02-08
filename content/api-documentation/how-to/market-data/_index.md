---
title: Market Data Examples
weight: 10
---

Alpaca provides market data from various sources. To learn more about data options and which one is right for you, please see [this page](https://docs.alpaca.markets/api-documentation/web-api/market-data/).

---

## Get Historical Price and Volume Data

By making a GET request to our [`/v1/bars`](https://docs.alpaca.markets/api-documentation/web-api/market-data/bars/) endpoint, you can see what a stock price was at a particular time.

{{< code-example exampleId="getPrices" pathURL="/api-documentation/how-to/market-data/examples">}}
