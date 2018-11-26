---
title: Paper Trading
weight: 100
---

# Paper Trading
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

## How to Start Using Paper Trading Account
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
- Paper trading account simulates our pattern-day-trader checks. Orders
  that would make 4th day trading within 5 business days will be rejected
  if the real-time net worth is below $25,000. Please read the
  [Pattern Day Trader Protection page]({{< relref "/broker-functions/pdt-protection" >}}) for more details.
- Paper trading account does NOT simulate dividends.
- Paper trading account does NOT send order fill emails.
- Market Data API works identically.
- You cannot change the account balance after it is created, unless you reset it.
- Orders are filled at the NBBO at the moment, although the size of order book is not considered.
- Orders may be partially filled at a predefined random rate.
