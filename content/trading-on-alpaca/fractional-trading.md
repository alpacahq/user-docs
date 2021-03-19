---
title: Fractional Trading
weight: 60
aliases:
  - /fractional-trading.md
---

## Fractional Trading

Fractional shares are fractions of a whole share, meaning that you don’t need to buy a whole share to own a portion of a company. You can now buy as little as $1 worth of shares for over 2,000 US equities.

By default all Alpaca accounts are allowed to trade fractional shares in both live and paper environments. Please make sure you reset your paper account if you run into any issues dealing with fractional shares.

### Supported Order Types

Alpaca currently supports fractional trading for `market` and `day` order types (support for other order types are in our roadmap). You can pass either a fractional amount (`qty`), or a notional value (`notional`) in any POST/v2/orders request. Note that entering a value for either parameters, will automatically nullify the other. If both `qty` and `notional` are entered the request will be rejected with an error status `400`.

Both `notional` and `qty` fields can take up to 9 decimal point values.

### Possible Fractional Order Requests

Notional Request

```json
{
  "symbol": "AAPL",
  "notional": 500.75,
  "side": "buy",
  "type": "market",
  "time_in_force": "day"
}
```

Fractional Request

```json
{
  "symbol": "AAPL",
  "qty": 3.654,
  "side": "buy",
  "type": "market",
  "time_in_force": "day"
}
```

### Supported Assets

Not all assets are fractionable yet so please make sure you query assets details to check for the parameter `fractionable = true`.

Supported fractionable assets would return a response that looks like this

```json
{
  "id": "b0b6dd9d-8b9b-48a9-ba46-b9d54906e415",
  "class": "us_equity",
  "exchange": "NASDAQ",
  "symbol": "AAPL",
  "name": "Apple Inc. Common Stock",
  "status": "active",
  "tradable": true,
  "marginable": true,
  "shortable": true,
  "easy_to_borrow": true,
  "fractionable": true
}
```

If you request a fractional share order for a stock that is not yet fractionable, the order will get rejected with an error message that reads `requested asset is not fractionable`.

### Dividends

Dividend payments occur the same way in fractional shares as with whole shares, respecting the proportional value of the share that you own.

For example if the dividend amount is $0.10 per share and you own 0.5 shares of that stock then you will receive $0.05 as dividend. As a general rule of thumb all dividends are rounded to the nearest penny.

Some notes on fractional trading

- We only support long fractional trading positions; short selling fractional shares is not supported.
- The expected price of fill is the NBBO quote at the time the order was submitted. We do not provide price improvement on fractional share orders. If you submit an order for a whole and fraction, the price for the whole share fill will be used to price the fractional portion of the order.
- Day trading fractional shares counts towards your day trade count
- You can cancel a fractional share order that is pending. Since fractional shares are only for `market` order types you would only be cancelling orders you enter when the market is closed, and calendar before the market opens - same way as with whole shares.
- Fees for fractional trading work the same way as with whole shares.

Alpaca does not make recommendations with regard to fractional share trading, whether to use fractional shares at all, or whether to invest in any specific security. A security’s eligibility on the list of fractional shares available for trading is not an endorsement of any of the securities, nor is it intended to convey that such stocks have low risk. Fractional share transactions are executed either on a principal or riskless principal basis, and can only be bought or sold with market orders during normal market hours.
