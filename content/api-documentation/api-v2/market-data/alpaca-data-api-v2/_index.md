---
title: Alpaca Data API v2
weight: 80
summary: Please note that Alpaca Data API v2 will be live starting from Feb 27, 2021. Alpaca Data API v2 provides market data in 2 two different plans, Basic and Pro. The Basic plan is included in both paper-only and live trading accounts as the default plan for free.
---

# Market Data

## Overview

Alpaca Data API v2 provides market data through an easy to use HTTP API for historical data and through websocket for real-time data.

We provide easy to use SDKs written in Python, Go and NodeJS. The SDK in C# is currently in development, we will let you know once it's ready for you to use.

{{< warning >}} Please note that Alpaca Data API v2 will be live starting from Feb 27, 2021. {{< /warning >}}


## Subscription Plans

Alpaca Data API v2 provides market data in 2 two different plans: **Basic** and **Pro**.

The Basic plan is included in both paper-only and live trading accounts as the default plan for free.
You will be able to subscribe to the Pro plan from the Dashboard after Feb 27, 2021, to sign up click [here](https://app.alpaca.markets/signup). 

|  | Basic | Pro |
| -------- | -------- | -------- |
| Pricing    | Free     | $49/mo     |
| Securities coverage    | US Stocks & ETFs     | US Stocks & ETFs    |
| Real-time market coverage    | IEX     | All US Stock Exchanges     |
| Websocket subscriptions   | 30 symbols    | Unlimited    |
| Historical data timefrane   | 5+ years | 5+ years     |
| Historical data delay| 15 minutes | 1 minute
| Historical API calls    | 200/min     | Unlimited |


The **Basic plan** consists of data from IEX (Investors Exchange LLC).

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

