---
title: Web API
weight: 10
---
# Web API
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
    https://{apiserver_domain}/v1/account
```

Alpaca's live API domain is `api.alpaca.markets`.

## Paper Trading
Alpaca's paper trading service uses a different domain and different credentials from
the live api. You'll need to connect to the right domain so that you don't
run your paper trading algo on your live account.

To use the paper trading api, set `APCA-API-KEY-ID` and
`APCA-API-SECRET-KEY` to your paper credentials, and set the domain to `paper-api.alpaca.markets`.

After you have tested your algo on paper and are ready to start running your algo
for real, you can switch the domain to the live domain, and the credentials to your
live credentials. Your algo will then start trading with real money.

To learn more about paper trading, visit the [paper trading page]({{< relref "/paper-trading/_index.md" >}}).

## Rate Limit
There is a rate limit for the API requests.  When it is exceeded, the API
server returns error response with HTTP status code 429.  The rate limit is
200 requests per every minute per API key.

## General Rules
### Time Format and Time Zone
All date time type inputs and outputs are serialized according to
[ISO8601](https://www.iso.org/iso-8601-date-and-time-format.html)
(more specifically [RFC3339](https://tools.ietf.org/html/rfc3339)).  The
communication does not assume a particular time zone, and this date time
serialization denominates the time offset of each value.

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

