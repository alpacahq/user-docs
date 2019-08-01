---
title: User Protections
weight: 200
---
# User Protections
We have enabled several types of protections to enhance your trading experience.

* [Pattern Day Trader (PDT) Protection] ({{<relref "#pattern-day-trader-protection-pdt-at-alpaca">}})
* [Day Trade Margin Call (DTMC) Protection] ({{<relref "#day-trade-margin-call-dtmc-at-alpaca">}})

## Pattern Day Trader (PDT) Protection at Alpaca
In order to prevent Alpaca Brokerage Account customers from unintentionally being
designated as a Pattern Day Trader (PDT), the Alpaca Trading platform checks the PDT
rule condition every time an order is submitted from a customer. If the order could potentially
result in the account being flagged as a PDT, the order is rejected, and API
returns error with HTTP status code 403 (Forbidden).

### The Rule
A day trade is defined as a round-trip pair of trades within the same day. A
buy must be occur first and then a sell of the same security must come later
in the day. The inverse does not make a day trade. Selling short and
covering the short on the same day is also considered a day trade.

An account is designated as a Pattern Day Trader if it makes four (4) day
trades within five (5) business days. Day trades less than this criteria
will not flag the account for PDT.

### Alpaca’s Order Rejection
Alpaca Trading platform monitors the number of day trades for the account
for the past 5 business days and rejects a newly submitted order if it
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
trading margin calls, Alpaca implements two forms of DTMC protections (We will be enabling these
protections beginning Monday 8/5/2019).

Day trading buying power is produced only when your `account.multiplier` is 4 as pattern day trader.
The following is applicable only for the accounts desginated as pattern day traders.
Please check your Account API result for the `multiplier` field.

Every trading day, you start with the new `day_trading_buying_power`. This beginning value is
calculated as `4 * (last_equity - last_maintenance_margin)`. Throughout day, every time you enter
the new position, this value is reduced and the same amount (regardless of position's P/L) is credited
back when you exit the position that you entered the samy day.

The maximum exposure of your day trading position is checked once as of the day close and DTMC is
issued next day if the maximum exposure of day trades exceeded your day trading
buying power at the beginning of day.

The `buying_power` value is bigger of `regt_buying_power` and `day_trading_buying_power`. Since
the basic buying power check runs on this `buying_power` value, you may exceed your `day_trading_buying_power`
when you enter the position, if `regt_buying_power` is bigger at one point of the day.

A particular example is as follows.

- Your equity is 50k
- You hold overnight positions up to 100k
- Your maintenance margin is 30k, therefore your day trading buying power at the beginning of day is 80k (50k - 30k)
- You sold all of 100k in the morning, which brings up your `regt_buying_power` up to 100k
- You now buy and sell the same security up to 100k
- At the end of day, 20k = 100k - 80k is the amount of DTMC

By default, Alpaca users have DTMC protections **on entry** of a position. This means that if
your entering order would exceed `day_trading_buying_power` at the moment, it will be blocked, even if
`regt_buying_power` still has room for it. This is based on the assumption that any entering position
could be day trades later in the day.
This option is the more conservative of the two DTMC protections that our users have.

The second DTMC protection option is protection **on exit** of a position. This means that Alpaca will block
the exit of positions that would cause a Day Trading Margin Call. This may cause users of this protection
method to be unable to liquidate a position until the next day.

One of the two protections will be enabled for all users (you cannot have neither protection enabled). If you would like to switch your protection option, please contact [our support](https://support.alpaca.markets/hc/en-us/requests/new).
Self-service protection change will be available to your dashboard later.

### The Rule
Day traders are required to have a minimum of $25,000 OR 25% of the total market value of securities (whichever is higher) maintained in their account.

The buying power of a pattern day trader is 4x the excess of the maintenance margin from the closing of the previous day. If you exceed this amount, you will receive a day trading margin call.

### Paper Trading
The same protection triggers in your paper trading account. It is
advised to test your algorithm with the realistic balance amount you
would manage when going live, to make sure your assumption works under
this DTMC protection as well.

For more details of Pattern Day Trader rule, please read
[FINRA's margin requirements](http://www.finra.org/investors/day-trading-margin-requirements-know-rules).
For more details on day trade margins, please read [FINRA's Mind Your Margin article](https://www.finra.org/investors/highlights/day-traders-mind-your-margin).
