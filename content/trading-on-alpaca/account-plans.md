---
title: Account Plans
weight: 50
aliases:
    - /account-plans.md
---

## Alpaca Paper Only Account (Paper-Trading)

Anyone globally can create an Alpaca Paper Only Account! All you need to do is [sign up](https://app.alpaca.markets/signup)
with your email address.

An Alpaca Paper Only Account is for **paper trading only**. It allows you to fully utilize the Alpaca API and
run your algorithm in our paper trading environment only. You won't be trading real money, but you will be able
to track your simulated activity and balance in the Alpaca web dashboard. As an Alpaca Paper Only Account
holder, you are only entitled to receive and make use of IEX market data. For more information about our paper
trading environment, please refer to [Paper Trading Specification]({{< relref "/trading-on-alpaca/paper-trading.md" >}}).

## Alpaca Brokerage Account (Live-Trading)

After creating an Alpaca Paper Only Account, you can enable live trading by becoming an Alpaca Brokerage Account
holder. This requires you to go through an account on-boarding process with Alpaca Securities LLC, a FINRA
member and SEC registered broker-dealer. Currently, we support brokerage accounts for
qualified U.S. residents and are currently running an invite-only beta program for non-U.S. residents and business entities globally.

With a brokerage account, you will be able to fully utilize Alpaca for your automated trading and investing needs.
Using the Alpaca API, you'll be able to buy and sell stocks in your brokerage account, and you'll receive
real-time consolidated market data. In addition, you will continue to be able to test your strategies and
simulate your trades in our paper trading environment. And with the Alpaca web dashboard, it's easy to monitor
both your paper trading and your real money brokerage account. All accounts are opened as margin accounts. Accounts with $2,000
or more equity will have access to margin trading and short selling.

**For U.S. Residents: Individual Taxable Account**

Alpaca Securities LLC supports individual taxable brokerage accounts. At this time, we do not support retirement accounts.


**For Non-U.S. Residents: Invite Only Beta**

We started an invite-only beta program for non-U.S. residents to open live-trading accounts. Due to the nature of this beta program, there are several restrictions including that the Alpaca Securities LLC brokerage is available internationally on a limited basis for "non-solicited" customers* only (* Any customer onboarding for live brokerage must electronically sign a Non-Solicitation Agreement certifying that the customer was not directly solicited to open a live brokerage account and must indicate how they became aware of Alpaca).

Although we are working to remove certain restrictions that we have with today’s beta program, we start sending invites to the users who are okay with **$30,000 initial funding** and **using international-wire transfer to deposit funds** (This beta program does not support TransferWise or Revolut. ACH does not work with a bank account outside the US). Please read [this announcement](https://alpaca.markets/blog/non-us-live-trading-beta/) to learn more about the beta program.

Please fill out [this form](https://forms.gle/vV96nn5zFBtHSrfU9) to be on the invite list.


{{< hash-link business>}}
**For Business Entities Globally: Business Trading Account**

We are currently running a beta program for the business trading account. This business trading account is prepared for business entities that intend to use Alpaca for trading purposes, but not for building apps/services. If you are looking to integrate with Alpaca API for your apps, please read [Build Apps/Services with Alpaca](https://alpaca.markets/docs/build-apps_services-with-alpaca/).

Today’s beta program for the business trading account accepts select entities that are mainly Corporations, LLCs, Partnerships in the U.S. and we continue to expand the coverage of eligible business entity types. There's no charge to open a business trading account, but it does have a **$30,000 account minimum**. Additionally, please note that business accounts are considered professional data subscribers for which we currently offer two plans: **Free** and **Unlimited**. Please read [this announcement](https://alpaca.markets/blog/business-brokerage-account-beta/) to learn more about the beta program.

You can [sign up](https://app.alpaca.markets/signup) with your email address and follow steps on the dashboard. When our system considers that your business entity is eligible for this beta program, you should be able to go on to upload necessary information through the process to submit the application, otherwise you will be placed on the waitlist. Please note that we are updating the beta eligibility criteria and business trading account opening flow with ongoing basis.


## Paper Only Account vs Brokerage Account Features

|<span style="font-size:14px">Feature</span>|<span style="font-size:14px">Alpaca Paper Only Account</span>|<span style="font-size:14px">Alpaca Brokerage Account</span>|
|---|---|---|
|Eligibility|All|Qualified U.S. Residents|
|Paper Trading|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px"> &#10003; </span>|
|Commission-Free Live Trading|N/A|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|API Access|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Free IEX Real-Time Data|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Multi-Factor Authentication|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Margin Trading|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Short Selling|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|
|Premarket/Afterhours Trading|<span style="color:#27e272;font-size:18px">&#10003;</span>|<span style="color:#27e272;font-size:18px">&#10003;</span>|

## Markets Supported

Currently, Alpaca only supports trading of listed U.S. stocks.
