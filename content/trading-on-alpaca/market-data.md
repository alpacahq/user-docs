---
title: Market Data
weight: 45
aliases:
    - /market-data.md
---

Alpaca provides you with different market data depending upon your account type. Below is a summary of the data feeds
available. **Currently, we only provide data for U.S. listed equities.**

## Alpaca Data API
**This is available to all Alpaca users.**

Users that have signed up with Alpaca but have not opened a real money brokerage account are able to receive free
real-time data from five exchanges. Please note that this data feed only includes quotes and trades
occurring on the order books on these exchanges.

- IEX (Investors Exchange LLC)
- NYSE National, Inc.
- Nasdaq BX, Inc.
- Nasdaq PSX
- NYSE Chicago, Inc.

## Consolidated Market Data
**This is available only to Alpaca Brokerage Accounts (Alpaca Securities brokerage account customers,
currently U.S. residents only)**

Alpaca Securities brokerage account customers are provided with a consolidated U.S. equity market data feed.
Currently we provide this data for free.

Consolidated stock market data is an aggregated reporting of all securities exchanges’ and alternative trading venues’
quote and trade data. It is the most relied upon type of market data, providing investors and traders globally with a
unified view of U.S. stock market prices and volumes. It also underpins the National Best Bid and Offer (NBBO), which
provides investors with a continuous view of the best available displayed buy and sell prices, and through Rule 611
ensures that investors receive the best available displayed prices on their trades, with a few exceptions.

### Our Data Provider and Terms of Usage

Consolidated market data is provided by third-party data vendor [Polygon](https://polygon.io/). Please note that this
data is currently only available to customers who affirm their **Non-Professional Subscriber** status. During the
account opening process, customers are directed through the data agreement, which is an agreement between the customer
and Polygon.

## Further Reading

For more details about how to access each data API, please read the [API Documentation]({{< relref "/api-documentation/api-v2/market-data/_index.md" >}}).
For more information about market data feeds, please see our [Medium post](https://medium.com/automation-generation/exploring-the-differences-between-u-s-stock-market-data-feeds-3da26946cbd6).
