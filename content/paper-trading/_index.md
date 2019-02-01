---
title: Paper Trading Specification
weight: 50
---

# About Paper Trading

When you run your algorithm with the live market, there are many
things that can happen that you may not see in backtesting.  Orders
may not be filled, prices may spike, or your network may get
disconnected and retry may be needed. During the software development
process, it is important to test your algorithm to catch these things
in advance.

Paper trading is a real-time simulation environment where you can test
your code. Everything on the broker side behaves the same way as the
live account except the orders aren't routed to the real exchanges.
Instead, the system simulates the order filling based on the real-time
quotes. It calculates the account balance and performance numbers with
virtual money so the real money in your live account isn't affected by
the activity in the paper trading account. You may create one paper
trading account per live account.

## Simulation Limitations

We are thrilled that many users have found paper trading useful, and we continue
to work on improving our paper trading assumptions so that users may have a 
superior experience. However, please note that paper trading is only a simulation. 
It provides a good approximation for what one might expect in real trading, but it is
not a substitute for real trading and performance may differ. Specifically,
paper trading does **not** account for:

* Market impact of your orders
* Information leakage of your orders
* Price slippage due to latency
* Order queue position (for non-marketable limit orders)
* Price improvement received
* Regulatory fees
* Dividends

If you are interested to incorporate these issues into your testing, you may
do so by trading a live brokerage account. Even small amounts of real
money can often provide insight into issues not seen in a simulation environment.

## Create A Paper Trading Account
Your initial paper trading account is created with $100k balance as
a default setting. You can reset the paper trading account at any
time later with arbitrary amount as you configure. After it is
created, the account works independently from your live account.
You have to have your brokerage account with any amount of funding
to create a paper trading account.

### Example:

You are creating a paper trading account with initial balance amount of $600.
You place orders in the paper trading account using this $600, and let's say
the net worth increases to $650. In the meantime, if you didn't trade anything
in the live account, your live account net worth wouldn't change. In other
words, the paper trading account is initialized with the net worth of the
live account, but they are independent of one another thereafter.

## How to Start Using Your Paper Trading Account
Your paper trading account will have a different API key from your live
account, and all you need to do to start using your paper trading account
is to replace your API key and API endpoint with ones for the paper trading.
The API spec is the same between the paper trading and live accounts.
The API endpoint (base URL) is displayed in your paper trading dashboard,
and please follow the instruction about how to set it depending on the
library you are using. In most cases, you need to set an environment variable
`APCA_API_BASE_URL=https://paper-api.alpaca.markets`

## Reset Paper Trading Account
You can reset your paper trading account at any time from the dashboard.
Go to the paper trading view and push the reset button.  Once you reset
the paper trading account, the trading history of the paper trading account
is wiped out, and a new paper trading account is created based on your
requested initial balance. Please make sure to re-generate the API key
as old key is associated with your old paper trading account and is not
valid any longer after you reset it.

## Specification
- Paper trading account simulates our Pattern Day Trader checks. Orders
  that would generate a 4th Day Trade within 5 business days will be rejected
  if the real-time net worth is below $25,000. Please read the
  [Pattern Day Trader Protection page]({{< relref "/orders/pdt-protection" >}}) for more details.
- Paper trading account does NOT simulate dividends.
- Paper trading account does NOT send order fill emails.
- Market Data API works identically.
- You cannot change the account balance after it is created, unless you reset it.
- Orders are filled only when they become marketable. This means that a non-marketable buy limit order
will not be filled until its limit price is equal to or greater than the best ask price,
and a non-marketable sell limit order will not be filled until its limit price is equal to or less than
the best bid.
- Your order quantity is not checked against the NBBO quantities. In other words, you can submit and receive
a fill for an order that is much larger than the actual available liquidity.
- When orders are eligible to be filled, they will receive partial fills for a random size 10% of the time. If
the order price is still marketable, the remaining quantity would be re-evaluated for a subsequent fill.
- Regardless of whether you have a Paper Trading Only account or a real money brokerage account,
all orders submitted in paper trading will be matched against the best available current market price (NBBO).  
