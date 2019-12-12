+++
aliases = [
    "/margin-and-shorting.md"
]
+++

---
title: Margin & Short Selling
weight: 70
---

**In order to trade on margin or sell short, you must have $2,000 or more account equity.
Accounts with less than $2,000 will not have access to these features and will be restricted to 1x buying power.**

## How Margin Works

Trading on margin allows you to trade and hold securities with a value that exceeds your account equity. This is made
possible by funds loaned to you by your broker, who uses your account's cash and securities as collateral.
For example, a Reg T Margin Account holding $10,000 cash may purchase and hold up to $20,000 in marginable securities
overnight (Note: some securities may have a higher maintenance margin requirement making the full 2x overnight buying power
effectively unavailable). In addition to the 2x buying power afforded to margin accounts, a Reg T Margin Account
flagged as a Pattern Day Trader(PDT) with $25,000 or greater equity will further be allowed to use up to 4x *intraday*
buying power. As an example, a PDT account holding $50,000 cash may purchase and hold up to $200,000 in securities
intraday; however, to avoid receiving a margin call the next morning, the securities held would need to be reduced
to $100,000 or less depending on the maintenance margin requirement by the end of the day.

### Initial Margin
Initial margin denotes the percentage of the purchase price of a security or basket of securities (purchased on margin)
that an account holder must pay for with available cash in the margin account, additions to cash in the margin account
or other marginable securities.

Alpaca uses the minimum initial margin requirement of 50% (of the purchase price of a security or basket of securities
that can be purchased on margin) set by Regulation T of the Federal Reserve Board.

### Maintenance Margin
Maintenance margin is the amount of cash or marginable securities required to continue holding an open position. FINRA
has set the minimum maintenance requirement to at least 25% of the total market value of the securities, but brokers
are free to set higher requirements as part of their risk management.

Alpaca uses the following table to calculate the overnight maintenance margin applied to each security held in an account:

|<span style="font-size:14px">Position Side</span>|<span style="font-size:14px">Condition</span>|<span style="font-size:14px">Margin Requirement</span>|
|---|---|---|
|Long|share price < $2.50|100% of EOD market value|
|Long|share price >= $2.50|30% of EOD market value|
|Long|2x Leveraged ETF|60% of EOD market value|
|Long|3x Leveraged ETF|90% of EOD market value|
|Short|share price < $5.00|Greater of $2.50/share or 100%|
|Short|share price >= $5.00|Greater of $5.00/share or 30%|

### Margin Calls
If your account does not satisfy its initial and maintenance margin requirements at the end of
the day, you will receive a margin call the following morning. We will contact you and advise you
of the call amount that you will need to satisfy either by depositing new funds or liquidating
some or all of your positions to reduce your margin requirement sufficiently.

We may contact you prior to the end of the day and ask you to liquidate your positions
immediately in the event that your account equity is materially below your maintenance requirement.
Furthermore, although we will make every effort to contact you so that you can determine how to best
resolve your margin call, we reserve the right to liquidate your holdings in the event we cannot
get ahold of you and your account equity is in danger of turning negative.


Calculating and tracking your margin requirement at all times is helpful to avoid receiving a
margin call. We strongly recommend doing so if you plan to aggressively use overnight leverage.
Please use a 50% initial requirement and refer to the maintenance margin table above.
In the future, we will provide real-time estimated initial and maintenance margin values as
part of the Account API to help users better manage their risk.


## Margin Interest Rate

We are pleased to offer a competitive and low annual margin interest rate of **3.75%**.

The rate is charged **only** on your account’s **end of day** (overnight) debit balance using the following calculation:

**daily margin interest charge = (settlement date debit balance * 0.0375) / 360**

Interest will accrue daily and post to your account at the end of each month. Note that if you have a settlement date
debit balance as of the end of day Friday, you will incur interest charges for 3 days (Fri, Sat, Sun).

As an example, if you deposited $10,000 into your account and bought $15,000 worth of securities that you held at
the end of the day, you would be borrowing $5,000 overnight and would incur a daily interest expense of
($5000 * 0.0375) / 360 = **$0.52**.

On the other hand, if you deposited $10,000 and bought $15,000 worth of stock that you liquidated the same day,
you would not incur any interest expense. In other words, this allows you to make use of the additional buying
power for intraday trading without any cost.


## Stock Borrow Rates

Alpaca currently **only** supports **opening** short positions in easy to borrow (“ETB”) securities. **Any open short order
in a stock that changes from ETB to HTB overnight will be automatically cancelled prior to market open**.

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

**Daily ETB stock borrow fee = (settlement date end of day total ETB short $ market value * 0.002) / 360**

And

**Daily HTB stock borrow fee = <span style="font-size:18px">&#931;</span>((each stock’s HTB short $ market value * that stock’s HTB rate) / 360)**

Note that if you hold short positions as of a Friday settlement date, you will incur stock borrow fees for 3 days (Fri, Sat, Sun).
