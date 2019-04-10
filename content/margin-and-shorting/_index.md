---
title: Margin & Short Selling
weight: 70
---

# Margin & Short Selling

We've begun providing qualified U.S. residents with early access to the Alpaca Power Account, a new account plan with 
our latest features including margin trading, short selling, support for business entities, and more. We’ll be 
onboarding users over the coming weeks and months, so [reserve your spot today by completing this form](https://goo.gl/forms/D9k1MMNtY9awXcum1), 
or [solve our puzzle](https://alpaca.markets/#power) and gain priority access! 

**The information contained in this section only applies to Power Accounts.**

## Margin Interest Rate

We are pleased to offer a competitive and low annual margin interest rate of **3.75%**.

The rate is charged **only** on your account’s **end of day** (overnight) debit balance using the following calculation:

**daily margin interest charge = (debit balance * 0.0375) / 360**

Interest will accrue daily and post to your account at the end of each month.

As an example, if you deposited $10,000 into your account and bought $15,000 worth of securities that you held at 
the end of the day, you would be borrowing $5,000 overnight and would incur a daily interest expense of 
($5000 * 0.0375) / 360 = **$0.52**. 

On the other hand, if you deposited $10,000 and bought $15,000 worth of stock that you liquidated the same day, 
you would not incur any interest expense. In other words, this allows you to make use of the additional buying 
power for intraday trading without any cost.


## Stock Borrow Rates

Alpaca currently **only** supports **opening** short positions in easy to borrow (“ETB”) securities. 

We are pleased to offer a competitive and low annual ETB stock borrow rate of **0.20%**.

However, please note that stock borrow availability changes daily, and we update our assets table each morning, so 
please use our API to check each stock’s borrow status daily. It is infrequent but names can go from ETB → HTB and 
vice versa.

While we do not currently support opening short positions in hard to borrow (“HTB”) securities, we will not 
force you to close out a position in a stock that has gone from ETB to HTB unless the lender has called the stock. 
**If a stock you hold short has gone from ETB to HTB, you will incur a higher daily stock borrow fee for that stock 
than the 0.20% ETB rate**. We do not currently provide HTB rates via our API, so please contact us in these cases.

Daily stock borrow fees are the fees incurred for all ETB shorts held in your account as of end of day plus any 
HTB shorts held at any point during the day, calculated as:

**Daily stock borrow fee = Daily ETB stock borrow fee + Daily HTB stock borrow fee**

Where,

**Daily ETB stock borrow fee = (end of day total ETB short $ market value * 0.002) / 360**

And

**Daily HTB stock borrow fee = <span style="font-size:18px">&#931;</span>((each stock’s HTB short $ market value * that stock’s HTB rate) / 360)**
