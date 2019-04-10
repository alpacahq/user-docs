---
title: API v2 Reference
weight: 100
---
# API v2 Reference

**The information contained in this section only applies to Power Accounts and is subject to revision.**
We've begun providing qualified U.S. residents with early access to the Alpaca Power Account, a new account plan with 
our latest features including margin trading, short selling, support for business entities, and more. We’ll be 
onboarding users over the coming weeks and months, so [reserve your spot today by completing this form](https://goo.gl/forms/D9k1MMNtY9awXcum1), 
or [solve our puzzle](https://alpaca.markets/#power) and gain priority access!

## Release Notes

* Updated Account and Assets entities with new fields for short selling and margin trading
* Updated Python SDK for API v2.0 support, example usage: 
  `api = tradeapi.REST('<key_id>', '<secret_key>', api_version=’v2’)`
* Orders that encounter account permissioning, insufficient buying power, or lack of available borrow will return 
a 403 “forbidden” error with an existing or new log message describing the reject. 
* Please also refer to Web API v1.0 documentation for existing API usage

## Endpoint

The endpoint host name will be the same as v1 (api.alpaca.markets/paper-api.alpaca.markets) but with different path (“/v2”).

## Authentication

API key between v1 and v2 are interchangable.

{{< rest-endpoint resource="account" method="GET" path="/v2/account" >}}

## Account Entity
### Example
{{< rest-entity-example name="account-v2">}}

### Properties
{{< rest-entity-desc name="account-v2">}}

## Notes on Orders and Positions
* If your account is set to `shorting_enabled: false`, any attempt to place a sell order in a stock that you have no 
position or with a quantity that exceeds your current position will result in a 403 `Forbidden` error. 
In the future, users who are approved for shorting will be able to set the shorting_enabled flag in their dashboard. 
This will extend to the creation of sub-accounts and help ensure strategies that are not supposed to short are never 
allowed to accidentally short.
* For accounts with `shorting_enabled: true`, if you place a sell order for an asset on which you are flat, or 
one that you have a short position in, it will be treated as a short sell. There are no new “side” parameters 
added for short selling or short covering. The existing values of “sell” and “buy” will be used, respectively.
* At this time, you are only allowed to short easy-to-borrow assets. Assets that are easy to borrow are marked 
with `easy_to_borrow: true`. A short sell order will be rejected with a 403 `Forbidden` status if the asset being 
sold is hard to borrow or not shortable at all. 
* At this time, an order that would flip your position in that stock from long to short or short to long is not 
supported and will return a 403 `Forbidden` error. Your position must first be <= 0 for shorts and must first 
be >= 0 for longs. In the future, we plan to support selling (buying) beyond an existing long (short) position 
directly into a short (long) position.
* At this time, if you have no position in a stock, but you have a pending short sell (buy) order, you will not 
be permitted to submit a buy (sell) order. The order will return a 403 `Forbidden` error.
* In order to allow the use of margin, the /orders endpoint will use the `buying_power` field to determine whether 
an account has enough buying power to buy/short a security. The exception to this rule is that assets marked 
with `marginable:false` cannot use margin lending to purchase the asset. 
* Any order that would increase your position (e.g. buy orders when your position is >= 0, sell orders when your 
position is <= 0) reduce buying power by their 104% of their order value (1.04 * qty * price). 
* Short positions will be reported with negative quantities and negative market values.

{{< rest-endpoint resource="assets" method="GET" path="/v2/assets" >}}
{{< rest-endpoint resource="assets" method="GET" path="/v2/assets/:id" >}}
{{< rest-endpoint resource="assets" method="GET" path="/v2/assets/{symbol}" >}}

## Asset Entity
### Example
{{< rest-entity-example name="asset-v2">}}

### Properties
{{< rest-entity-desc name="asset-v2">}}
