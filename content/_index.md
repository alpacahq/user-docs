---
title: Documentation
type: index
outputs:
 - html
 - rss
 - json
---


# Where Should I Start?

## Sign up to generate your API key

In order to start using Alpaca [Web API]({{< relref "/api-documentation/web-api/_index.md" >}}), you will need to 
obtain your API key. An API key consists of two strings, one called **Key ID** which always appears in your dashboard, 
and **Secret Key**, which appears only once when generating it. Make sure you keep both in your hands.

To obtain your API key, you need to sign up on the [signup page](https://app.alpaca.markets/signup) by inputting 
your email address and password, then confirming your password through receiving a code on your email.

## Read through the Web API Document

Learn how API works by read through our [Web API document]({{< relref "/api-documentation/web-api/_index.md" >}}). It 
describes how our API works in details so you can prepare for your algorithm to interact with it.

You can also learn about [SDK for your language]({{< relref "/api-documentation/client-sdk/_index.md" >}}). Choose 
one of the supported ones and start building your idea.

## Copy, paste, and run code examples

You can learn how the process of algorithmic trading works by actually running a list of sample algorithms written 
in Python in a paper-trading or live-trading environment.

You can also run code examples for each specific function such as getting market data, placing new orders, and 
getting a list of existing orders.

## Start live trading

After signing up on the signup page with your email and password, you land on the dashboard where you can view your 
positions, historical performance, and orders. On the dashboard, you can follow the top-left link to start 
live trading.

For you to start live trading, you need to open an Alpaca Securities brokerage account. Currently, Alpaca Securities 
brokerage account is available only for the US residents, and requires you to complete the account application. 
Once you complete the application, your information is going to be reviewed and approved if everything is good. 
For more information about account review process, please read [Alpaca Securities FAQ](https://support.alpaca.markets/hc/en-us/) 
page.

## Deposit your money

In order to start trading, you need to deposit your money into your account. Go to your dashboard, link your bank
account, and initiate an ACH transfer from your bank to your Alpaca account. Please also read our 
[FAQ page for banking](https://support.alpaca.markets/hc/en-us/sections/360001964091-Banking-and-Transfers), too.

## Have your Quantopian algorithms work at Alpaca

If you are coming from the popular backtesting service Quantopian, there is something you want to read about [how easy 
it is to run your algorithm for live trading with Alpaca]({{< relref "/platform-migration/zipline-to-pylivetrader" >}}).