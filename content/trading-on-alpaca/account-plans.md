---
title: Account Plans
weight: 50
aliases:
    - /account-plans.md
---

## Alpaca Paper Only Account

Anyone globally can create an Alpaca Paper Only Account! All you need to do is [sign up](https://app.alpaca.markets/signup)
with your email address.

An Alpaca Paper Only Account is for **paper trading only**. It allows you to fully utilize the Alpaca API and
run your algorithm in our paper trading environment only. You won't be trading real money, but you will be able
to track your simulated activity and balance in the Alpaca web dashboard. As an Alpaca Paper Only Account
holder, you are only entitled to receive and make use of IEX market data. For more information about our paper
trading environment, please refer to [Paper Trading Specification]({{< relref "/trading-on-alpaca/paper-trading.md" >}}).

## Alpaca Brokerage Account

After creating an Alpaca Paper Only Account, you can enable live trading by becoming an Alpaca Brokerage Account
holder. This requires you to go through an account on-boarding process with Alpaca Securities LLC, a FINRA
member and SEC registered broker-dealer. Currently, we only support brokerage accounts for
qualified U.S. residents.

With a brokerage account, you will be able to fully utilize Alpaca for your automated trading and investing needs.
Using the Alpaca API, you'll be able to buy and sell stocks in your brokerage account, and you'll receive
real-time consolidated market data. In addition, you will continue to be able to test your strategies and
simulate your trades in our paper trading environment. And with the Alpaca web dashboard, it's easy to monitor
both your paper trading and your real money brokerage account. All accounts are opened as margin accounts. Accounts with $2,000
or more equity will have access to margin trading and short selling.

**Individual Taxable Account**

Alpaca Securities LLC supports individual taxable brokerage accounts. At this time, we do not support retirement accounts.


{{< hash-link business>}}
**Business Account**

We support select business accounts for U.S. and non-U.S. residents. There's no charge to open a business account,
but it does have a **$30,000 account minimum**. Additionally, please note that business accounts are considered
professional data subscribers for which we currently are **unable** to provide Polygon's consolidated market data.
You would still have access to our free IEX real-time data, but if you wanted a consolidated market data feed,
you'd need to obtain this through another provider.

Hedge funds, RIAs, app developers, and service providers, please [complete the form](https://docs.google.com/forms/d/e/1FAIpQLScH_5wylQNILGedoS_mAS7-djbU24hDqaOERMLiBVRRn9IaJA/viewform?usp=sf_link)
if you are interested in opening a business account (currently available by invite-only).  You can also [read more](https://medium.com/automation-generation/trading-through-a-business-account-at-alpaca-6ccb79709797)
about some of the entities we support.

**Account for non-U.S. residents?**

We've begun collecting interest from non-U.S. residents seeking commission-free U.S. stock trading.
Please voice your interest by [completing this form](https://forms.gle/umPhEzWtUEuHAuVVA).


## Paper Only Account vs Brokerage Account Features

|<span style="font-size:14px">Feature</span>|<span style="font-size:14px">Alpaca Paper Only Account</span>|<span style="font-size:14px">Alpaca Brokerage Account</span>|
|---|---|---|
|Eligibility|All|Qualified U.S. Residents|
|Paper Trading|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px"> &#10003; </span>|
|Commission-Free Live Trading|N/A|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|API Access|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Free IEX Real-Time Data|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Free Consolidated Real-Time Data|&#10060;|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Multi-Factor Authentication|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Margin Trading|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Short Selling|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Premarket/Afterhours Trading|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|

## Markets Supported

Currently, Alpaca only supports trading of listed U.S. stocks.
