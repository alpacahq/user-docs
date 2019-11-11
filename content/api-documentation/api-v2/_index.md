---
title: API v2
weight: 10
---
# API v2
Alpaca is a modern platform for algorithmic trading.  Alpaca's
API is the interface for your trading algo to communicate with Alpaca's brokerage
service.

The API allows your trading algo to access real-time price, fundamentals,
place orders and manage your portfolio, in either REST (pull) or streaming
(push) style.

In order to start trading with Alpaca API, please sign up
[here](https://alpaca.markets/).

Once you have signed up and have familiarized yourself with our API, please
check out our [python client](https://github.com/alpacahq/alpaca-trade-api-python)
to begin writing your own algo!

## Authentication
Every private API call requires key-based authentication. API keys can
be acquired in the developer web console.  The client must provide a pair of API
key ID and secret key in the HTTP request headers named
`APCA-API-KEY-ID` and `APCA-API-SECRET-KEY` respectively.

Here is an example using curl showing how to authenticate with the API.

```
curl -X GET \
    -H "APCA-API-KEY-ID: {YOUR_API_KEY_ID}" \
    -H "APCA-API-SECRET-KEY: {YOUR_API_SECRET_KEY}"\
    https://{apiserver_domain}/v2/account
```

Alpaca's live API domain is `api.alpaca.markets`.

## Paper Trading
Alpaca's paper trading service uses a different domain and different credentials from
the live API. You'll need to connect to the right domain so that you don't
run your paper trading algo on your live account.

To use the paper trading api, set `APCA-API-KEY-ID` and
`APCA-API-SECRET-KEY` to your paper credentials, and set the domain to `paper-api.alpaca.markets`.

After you have tested your algo in the paper environment and are ready to start running your algo in the live
environment, you can switch the domain to the live domain, and the credentials to your
live credentials. Your algo will then start trading with real money.

To learn more about paper trading, visit the [paper trading page]({{< relref "/trading-on-alpaca/paper-trading.md" >}}).

## Rate Limit
There is a rate limit for the API requests.  When it is exceeded, the API
server returns error response with HTTP status code 429.  **The rate limit is
200 requests per every minute per API key.**

## General Rules
### Time Format and Time Zone
Most date time type inputs and outputs are serialized according to
[ISO8601](https://www.iso.org/iso-8601-date-and-time-format.html)
(more specifically [RFC3339](https://tools.ietf.org/html/rfc3339)).  The
communication does not assume a particular time zone, and this date time
serialization denominates the time offset of each value. An exception to
this rule is the Calendar endpoint, which for example, accepts date
strings of the format "%Y-%m-%d".

### Numbers
Decimal numbers are returned as strings to preserve full precision across
platforms. When making a request, it is recommended that you also convert
your numbers to strings to avoid truncation and precision errors.

### IDs
Object ID in Alpaca system uses UUID v4.  When making requests, the format
with dashes is accepted.

```
904837e3-3b76-47ec-b432-046db621571b
```

### Assets and Symbology
An asset in this API is a tradable or non-tradable financial instrument.
Alpaca maintains our own asset database and assigns an internal
ID for each asset which you can use to identify assets to specify in API
calls.  Assets are also identified by a combination of symbol, exchange,
and asset class.  The symbol of an asset may change over the time, but
the symbol for an asset is always the one at the time API call is made.

When the API accepts a parameter named `symbol`, you can use one of the
following four different forms unless noted otherwise.

    - "{symbol}"
    - "{symbol}:{exchange}"
    - "{symbol}:{exchange}:{asset_class}"
    - "{asset_id}"

Typically the first form is enough, but in the case multiple assets are
found with a symbol (the same symbol may be used in different exchanges or
asset classes), the most commonly-used asset is assumed. To avoid
the ambiguity, you can use the second or third form with suffixes joined
by colons (:)   Alternatively, `asset_id` is guaranteed as unique, in the
form of UUID v4. When the API accepts `symbols` to specify more than one
symbol in one API call, the general rule is to use commas (,) to separate
them.

All of four symbol forms are case-sensitve.

## v2 Release Notes

* Updated Account and Assets entities with new fields for short selling and margin trading
* Updated
* Orders that encounter account permissioning, insufficient buying power, or lack of available borrow will return
a 403 “forbidden” error with an existing or new log message describing the reject.
* The endpoint host name will be the same as v1 (api.alpaca.markets/paper-api.alpaca.markets) but with different path (“/v2”).
* API key between v1 and v2 are interchangeable.
* Updated Python SDK for API v2.0 support, example usage:
  `api = tradeapi.REST('<key_id>', '<secret_key>', api_version=’v2’)`

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
* Market buy orders that would increase your position (e.g. buy order for a stock where your position in that stock is >= 0)
are converted to marketable limit orders with a 2.5% to 4% price collar.
* Market sell orders that would increase your position (e.g. sell order for a stock when your position in that stock is <= 0)
are NOT converted but are subject to a pre-trade buying power risk check of 102.5 to 104% of the order value.
* Short positions will be reported with negative quantities and negative market values.

# API Endpoints
