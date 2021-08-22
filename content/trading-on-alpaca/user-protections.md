---
title: User Protections
weight: 200
aliases:
    - /user-protections.md
---

We have enabled several types of protections to enhance your trading experience.

* [Pattern Day Trader (PDT) Protection]({{<relref "#pattern-day-trader-pdt-protection-at-alpaca">}})
* [Day Trade Margin Call (DTMC) Protection]({{<relref "#day-trade-margin-call-dtmc-protection-at-alpaca">}})

Please note that these do not apply to crypto trading as cryptocurrencies are not marginable. Pattern Day Trading rule does not apply to crypto trading either.

## Pattern Day Trader (PDT) Protection at Alpaca
In order to prevent Alpaca Brokerage Account customers from unintentionally being
designated as a Pattern Day Trader (PDT), the Alpaca Trading platform checks the PDT
rule condition every time an order is submitted from a customer. If the order could potentially
result in the account being flagged as a PDT, the order is rejected, and API
returns error with HTTP status code 403 (Forbidden).

### The Rule
A day trade is defined as a round-trip pair of trades within the same day (including extended hours). A
buy must occur first and then a sell of the same security must come later
in the day. The inverse does not make a day trade. Selling short and
covering the short on the same day is also considered a day trade.

An account is designated as a Pattern Day Trader if it makes four (4) day
trades within five (5) business days. Day trades less than this criteria
will not flag the account for PDT.

Cryptocurrency trading is not subject to the PDT rule. As a result, crypto orders are not evaluated by PDT protection logic and round-trip crypto trades on the same day do not contribute to the day trade count.

### Alpaca’s Order Rejection
Alpaca Trading platform monitors the number of day trades for the account
for the past 5 business days and rejects a newly submitted orders **on exit** of a position if it
could potentially result in the account being flagged for PDT. This
protection triggers only when the previous day's closing account equity is less than $25,000 at
the time of order submission.

In addition to the filled orders, the system also takes into
consideration pending orders in the account. In this case, regardless of
the order of pending orders, a pair of buy and sell orders is counted as
a potential day trade. This is because orders that are active (pending)
in the marketplace may fill in random orders. Therefore, even if your
sell limit order is submitted first (without being filled yet) and
another buy order on the same security is submitted later, this buy
order will be blocked if your account already has 3 day trades in
the last 5 business days.

### Paper Trading
The same protection triggers in your paper trading account. It is
advised to test your algorithm with the realistic balance amount you
would manage when going live, to make sure your assumption works under
this PDT protection as well.

For more details of Pattern Day Trader rule, please read
[FINRA website](http://www.finra.org/investors/day-trading-margin-requirements-know-rules).

---

## Day Trade Margin Call (DTMC) Protection at Alpaca
In order to prevent Alpaca Brokerage Account customers from unintentionally receiving day
trading margin calls, Alpaca implements two forms of DTMC protection.


### The Rule
Day traders are required to have a minimum of $25,000 OR 25% of the total market value of securities (whichever is higher) maintained in their account.

The buying power of a pattern day trader is 4x the excess of the maintenance margin from the closing of the previous day. If you exceed this amount, you will receive a day trading margin call.

### How Alpaca's DTMC Protection Settings Work
Users only receive day trading buying power when marked as a pattern day trader. If the user is designated a
pattern day trader, the `account.multiplier` is equal to 4.

Daytrading buying power cannot increase beyond its start of day value. In other words, closing an overnight position will not add to your daytrading buying power.

The following scenarios and protections are applicable only for accounts that are designated as pattern day traders.
Please check your Account API result for the `multiplier` field.

Every trading day, you start with the new `daytrading_buying_power`. This beginning value is
calculated as `4 * (last_equity - last_maintenance_margin)`. The `last_equity` and `last_maintenance_margin` values
can be accessed through Account API. These values are stored from the end of the previous trading day.

Throughout the day, each time you enter a new position, your `daytrading_buying_power` is reduced by that amount. When you exit
that position within the same day, that same amount is credited back, regardless of position's P/L.

At the end of the trading day, on close, the maximum exposure of your day trading position is checked. A Day Trade Margin Call (DTMC) is
issued the next day if the maximum exposure of day trades exceeded your day trading buying power from the beginning of that day.

The `buying_power` value is the larger of `regt_buying_power` and `daytrading_buying_power`. Since
the basic buying power check runs on this `buying_power` value, you could be exceeding your `daytrading_buying_power`
when you enter the position if `regt_buying_power` is larger than your `daytrading_buying_power` at one point in the day.

The following is an example scenario:

- Your equity is $50k
- You hold overnight positions up to $100k
- Your maintenance margin is $30k (~30%), therefore your day trading buying power at the beginning of day is $80k using the calculation of `4 * ($50k - $30k)`
- You sell all of the overnight positions ($100k value) in the morning, which brings your `regt_buying_power` up to $100k
- You now buy and sell the same security up to $100k
- At the end of the day, you have a $20k Day Trade Margin Call ($100k - $80k)

By default, Alpaca users have DTMC protections **on entry** of a position. This means that if
your entering order would exceed `daytrading_buying_power` at the moment, it will be blocked, even if
`regt_buying_power` still has room for it. This is based on the assumption that any entering position
could be day trades later in the day.
This option is the more conservative of the two DTMC protections that our users have.

The second DTMC protection option is protection **on exit** of a position. This means that Alpaca will block
the exit of positions that would cause a Day Trading Margin Call. This may cause users to be unable to liquidate a position until the next day.

Neither of the DTMC protection options evaluate crypto orders since crypto cannot be purchased using margin.

One of the two protections will be enabled for all users (you cannot have both protections disabled). If you would like to switch your protection option, please contact [our support](https://support.alpaca.markets/hc/en-us/requests/new).

We are working towards features to allow users to change their DTMC protection setting on their own without support help.


### Equity/Order Ratio Validation Check  
In order to help Alpaca Brokerage Account customers from placing orders larger than the calculated buying power, Alpaca has instituted a control on the account independent of the buying power for the account.  Alpaca will restrict the account to closing transactions when an account has a position that is 600% larger than the equity in the account.  The account will remain restricted for closing transactions until a member of Alpaca’s trading team reviews the account.  The trading team will either clear the alert by allowing opening transactions or will notify the client of the restriction and take corrective actions as necessary. 


### Paper Trading
The same protection triggers in your paper trading account. It is
advised to test your algorithm with the realistic balance amount you
would manage when going live, to make sure your assumption works under
this DTMC protection as well.

For more details of Pattern Day Trader rule, please read
[FINRA's margin requirements](http://www.finra.org/investors/day-trading-margin-requirements-know-rules).
For more details on day trade margins, please read [FINRA's Mind Your Margin article](https://www.finra.org/investors/day-trading).
