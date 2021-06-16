---
title: Alpaca Data API v2
weight: 80
summary: Please note that Alpaca Data API v2 is now in public beta. Alpaca Data API v2 provides market data in 2 two different plans, Free and Unlimited. The Free plan is included in both paper-only and live trading accounts as the default plan for free.
---

# Market Data

## Overview

Alpaca Data API v2 provides market data through an easy to use HTTP API for historical data and through websocket for real-time data.

We provide easy to use SDKs written in Python, Go, NodeJS and C#.

{{< note >}} Alpaca Data API v2 is currently in public beta. 
Please keep in mind that the public beta version may be less stable. {{< /note >}}


## Public beta status

*Last update: Apr 1, 2021*

We finished backfilling all trades, quotes and bars for the 2016-2020 period, these are now available through the APIs.

|  | 2021 | 2020 | 2019 | 2018 | 2017 | 2016 |
| -------- | -------- | -------- | -------- | -------- | -------- | -------- |
| Trades     | Done     | Done     | Done      | Done     | Done      | Done      |
| Quotes     | Done    | Done      | Done     | Done      | Done      | Done       |
| Bars    | Done    | Done       | Done       | Done      | Done       | Done      |



## Subscription Plans

Alpaca Data API v2 provides market data in 2 two different plans: **Free** and **Unlimited**.

The Free plan is included in both paper-only and live trading accounts as the default plan for free.

During the public beta, no user will be charged with the Unlimited plan. To sign up click [here](https://app.alpaca.markets/signup). 

|  | Free | Unlimited |
| -------- | -------- | -------- |
| Pricing    | Free     | $49/mo     |
| Securities coverage    | US Stocks & ETFs     | US Stocks & ETFs    |
| Real-time market coverage    | IEX     | All US Stock Exchanges     |
| Websocket subscriptions   | 30 symbols    | Unlimited    |
| Historical data timeframe   | 5+ years | 5+ years     |
| Historical data delay| 15 minutes | 1 minute
| Historical API calls    | 200/min     | Unlimited |


The **Free plan** consists of data from IEX (Investors Exchange LLC).

For the **Unlimited plan**, we receive direct feeds from the CTA (administered by NYSE) and UTP (administered by Nasdaq) SIPs. These 2 feeds combined offer 100% market volume. 


## Exchanges

List of stock exchanges which are supported by Alpaca.

The tape id of each exchange is returned in all market data requests. You can use this table to map the code to an exchange.

| Exchange Code | Name of Exchange |
| -------- | -------- | -------- | 
| A    | NYSE American (AMEX)     |
| B    | NASDAQ OMX BX     |
| C      | National Stock Exchange     |
| D      | FINRA ADF     |
| E      | Market Independent     |
| H      | MIAX      |
| I      | International Securities Exchange     |
| J      | Cboe EDGA     |
| K      | Cboe EDGX     |
| L      | Long Term Stock Exchange     |
| M      | Chicago Stock Exchange |
| N      | New York Stock Exchange     |
| P      | NYSE Arca     |
| Q      | NASDAQ OMX     |
| S      | NASDAQ Small Cap     |
| T      | NASDAQ Int     |
| U      | Members Exchange     |
| V      | IEX     |
| W      | CBOE     |
| X      | NASDAQ OMX PSX     |
| Y      | Cboe BYX     |
| Z      | Cboe BZX     |


## Conditions

Each feed/exchange uses its own set of codes to identify trade and quote conditions, so the same condition may have a different code depending on the originator of the data. 

### Trade conditions

### CTS

The table below shows codes that denotes a particular condition applicable to the trade from the CTA Plan. 

For more information, see page 64 of the [Consolidated Tape System (CTS) Specification](https://www.ctaplan.com/publicdocs/ctaplan/CTS_Pillar_Output_Specification.pdf).


| Code | Value | Code | Value |
| -------- | -------- | -------- | -------- |
| Space     | Regular Sale    | Q    | Market Center Official Open    |
| B     | Average Price Trade    | R    | Seller     |
| C     | Cash Trade (Same Day Clearing)     | T     | Extended Hours Trade     |
| E     | Automatic Execution    | U     | Extended Hours Sold (Out Of Sequence)     |
| F     | Inter-market Sweep Order     | V     | Contingent Trade     |
| H     | Price Variation Trade     | X     | Cross Trade    |
| I     | Odd Lot Trade     | Z     | Sold (Out Of Sequence)     |
| K     | Rule 127 (NYSE only) or Rule 155 (NYSE MKT only)     | 4     | Derivatively Priced     |
| L     | Sold Last (Late Reporting)     | 5     | Market Center Reopening Trade     |
| M     | Market Center Official Close     | 6     | Market Center Closing Trade     |
| N     | Next Day Trade (Next Day Clearing)    | 7     | Qualified Contingent Trade     |
| O     | Market Center Opening Trade    | 8     | Reserved     |
| P     | Prior Reference Price     | 9     | Corrected Consolidated Close Price as per Listing Market     |


### UTDF

The table below shows condition codes from the UTP Plan. 

For more information, see page 43 of the [UTP Specification](https://www.utpplan.com/DOC/UtpBinaryOutputSpec.pdf#page=43).

| Code | Value | Code | Value |
| -------- | -------- | -------- | -------- |
| @     | Regular Sale    | R    | Seller    |
| A     | Acquisition    | S    | Split Trade    |
| B     | Bunched Trade     | T     | Form T     |
| C     | Cash Sale    | U     | Extended trading hours (Sold Out of Sequence)     |
| D     | Distribution     | V     | Contingent Trade     |
| E     | Placeholder     | W     | Average Price Trade    |
| F     | Intermarket Sweep    | X     | Cross Trade     |
| G     | Bunched Sold Trade     | Y     | Yellow Flag Regular Trade     |
| H     | Price Variation Trade     | Z     | Sold (out of sequence)     |
| I     | Odd Lot Trade     | 1     | Stopped Stock (Regular Trade)     |
| K     | Rule 155 Trade (AMEX)   | 4     | Derivatively priced     |
| L     | Sold Last    | 5     | Re-Opening Prints     |
| M     | Market Center Official Close    | 6     | Closing Prints     |
| N     | Next Day    | 7     | Qualified Contingent Trade(“QCT”)     |
| O     | Opening Prints     | 8     | Placeholder For 611 Exempt     |
| P     | Prior Reference Price   | 9     | Corrected Consolidated Close (per listing market)     |
| Q     | Market Center Official Open     | 



## Quote conditions

### CQS

The table below shows codes that denotes a particular condition applicable to a quote from the CTA Plan. 

For more information, see Appendix G of the [CQS Specification](https://www.ctaplan.com/publicdocs/ctaplan/CQS_Pillar_Output_Specification.pdf).

| Code | Value |
| -------- | -------- |
| A    | Slow Quote Offer Side     |
| B    | Slow Quote Bid Side      |
| E    | Slow Quote LRP Bid Side      |
| F    | Slow Quote LRP Offer Side     |
| H    | Slow Quote Bid And Offer Side    |
| O    | Opening Quote     |
| R    | Regular Market Maker Open     |
| W    | Slow Quote Set Slow List     |
| C    | Closing Quote    |
| L    | Market Maker Quotes Closed     |
| U    | Slow Quote LRP Bid And Offer     |
| N    | Non Firm Quote     |
| 4    | On Demand Intra Day Auction     |

### UQDF

The table below shows condition codes from the UTP Plan.

For more information, see the [UQDF Specification](http://www.utpplan.com/DOC/uqdfspecification.pdf).

| Code | Value |
| -------- | -------- |
| A    | Manual Ask Automated Bid     |
| B    | Manual Bid Automated Ask      |
| F    | Fast Trading      |
| H    | Manual Bid And Ask     |
| I    | Order Imbalance    |
| L    | Closed Quote     |
| N    | Non Firm Quote     |
| O    | Opening Quote Automated     |
| R    | Regular Two Sided Open    |
| U    | Manual Bid And Ask Non Firm     |
| Y    | No Offer No Bid One Sided Open     |
| X    | Order Influx      |
| Z    | No Open No Resume     |
| 4    | On Demand Intra Day Auction     |


## Status messages

Status messages can be used to identify security statuses and trading halts real-time via websocket streaming. Each feed uses its own set of indicators.


## Security status

### CTS

The table below shows security indicators from the CTA Plan.

For more information, see the [CTS Specification](https://www.ctaplan.com/publicdocs/ctaplan/CTS_Pillar_Output_Specification.pdf).


| Code | Value |
| -------- | -------- |
| 2    | Trading Halt     |
| 3    | Resume      |
| 5    | Price Indication      |
| 6    | Trading Range Indication     |
| 7    | Market Imbalance Buy    |
| 8    | Market Imbalance Sell     |
| 9    | Market On Close Imbalance Buy    |
| A    | Market On Close Imbalance Sell     |
| C    | No Market Imbalance    |
| D    | No Market On Close Imbalance     |
| E    | Short Sale Restriction     |
| F    | Limit Up-Limit Down      |


### UTDF

The table below shows security indicators from the UTP Plan.

For more information, see the [UTP Specification](https://www.utpplan.com/DOC/UtpBinaryOutputSpec.pdf#page=43).

| Code | Value |
| -------- | -------- |
| H    | Trading Halt     |
| Q    | Quotation Resumption      |
| T    | Trading Resumption      |
| P    | Volatility Trading Pause     |



## Halt reason

### CTS

The table below shows halt reasons from the CTA Plan.

For more information, see the [CTS Specification](https://www.ctaplan.com/publicdocs/ctaplan/CTS_Pillar_Output_Specification.pdf).


| Code | Value |
| -------- | -------- |
| D    | News Released (formerly News Dissemination)     |
| I    | Order Imbalance      |
| M    | Limit Up-Limit Down (LULD) Trading Pause      |
| P    | News Pending     |
| X    | Operational    |
| Y    | Sub-Penny Trading     |
| 1    | Market-Wide Circuit Breaker Level 1 – Breached   |
| 2    | Market-Wide Circuit Breaker Level 2 – Breached     |
| 3    | Market-Wide Circuit Breaker Level 3 – Breached    |

### UTDF

The table below shows halt reasons from the UTP Plan.

For more information, see the [UTP Specification](https://www.utpplan.com/DOC/UtpBinaryOutputSpec.pdf#page=43).

| Code | Value |
| -------- | -------- |
| T1    | Halt News Pending    |
| T2    | Halt News Dissemination      |
| T5    | Single Stock Trading Pause In Affect      |
| T6    | Regulatory Halt Extraordinary Market Activity     |
| T8    | Halt ETF    |
| T12    | Trading Halted; For information requested by NASDAQ      |
| H4    | Halt Non Compliance     |
| H9    | Halt Filings Not Current     |
| H10    | Halt SEC Trading Suspension    |
| H11    | Halt Regulatory Concern      |
| O1    | Operations Halt, Contact Market Operations      |
| IPO1    | IPO Issue not yet Trading     |
| M1    | Corporate Action      |
| M2    | Quotation Not Available      |
| LUDP    | Volatility Trading Pause     |
| LUDS    | Volatility Trading Pause – Straddle Condition      |
| MWC1    | Market Wide Circuit Breaker Halt – Level 1     |
| MWC2    | Market Wide Circuit Breaker Halt – Level 2     |
| MWC3    | Market Wide Circuit Breaker Halt – Level 3      |
| MWC0    | Market Wide Circuit Breaker Halt – Carry over from previous day      |
| T3    | News and Resumption Times     |
| T7    | Single Stock Trading Pause/Quotation-Only Period     |
| R4    | Qualifications Issues Reviewed/Resolved; Quotations/Trading to Resume     |
| R9    | Filing Requirements Satisfied/Resolved; Quotations/Trading To Resume      |
| C3    | Issuer News Not Forthcoming; Quotations/Trading To Resume      |
| C4    | Qualifications Halt ended; maint. Req. met; Resume     |
| C9    | Qualifications Halt Concluded; Filings Met; Quotes/Trades To Resume     |
| C11    | Trade Halt Concluded By Other Regulatory Auth,; Quotes/Trades Resume     |
| R1    | New Issue Available      |
| R    | Issue Available      |
| IPOQ    | IPO security released for quotation     |
| IPOE    | IPO security – positioning window extension     |
| MWCQ    | Market Wide Circuit Breaker Resumption     |




Please visit our [Support](https://alpaca.markets/support/) page to learn more about our market data products.

