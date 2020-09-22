---
title: Alpaca Works With
weight: 90
---

# Alpaca Works With
Alpaca works with many ways to trade! We pride ourselves in enabling many partnerships and allow our users to customize and automate their trade flow. We are constantly working on new ways to enable you to trade. If you are looking to build applications with Alpaca API and OAuth 2.0, please read [OAuth Integration Guide](https://alpaca.markets/docs/build-apps_services-with-alpaca/oauth-guide/) to start building.

**Check out [Integrations page](https://alpaca.markets/integrations) to see the list of apps that work with Alpaca through OAuth 2.0**

* [TradingView](https://tradingview.com/broker/Alpaca)
* [IFTTT](https://ifttt.com/alpaca)
* [Blueshift](./blueshift)
* [Zapier (Early Access)](https://zapier.com/apps/alpaca/integrations)
* [Slack](./alpaca-for-slack)


## Third-Party Tools Integration
We provide several open source libraries to aid in the platform integration process and help expedite the time it takes for you to get up and running with Alpaca. Below is a list of our official platform integration libraries. If you would like to see us provide integration support for a platform not listed here, please let us know at our Feature Request and Issues repo
[here](https://github.com/alpacahq/Alpaca-API).


**Other helpful packages for migrating your trading activity to Alpaca:**

* [backtrader integrated with Alpaca SDK](https://github.com/alpacahq/alpaca-backtrader-api/) is a python library for the Alpaca trade API within backtrader framework. It allows rapid trading algo development easily, with support for the both REST and streaming interfaces. For details of each API behavior, please see the online API document.
* [pylivetrader](https://github.com/alpacahq/pylivetrader/) is a live trading framework in
python 3 that is API-compatible with [zipline](https://github.com/quantopian/zipline/). You can run your Quantopian algorithm with the minimum code change with Alpaca API. For more details about how to convert your Quantopian algorithm, please see
[zipline to pylivetrader migration](./zipline-to-pylivetrader/).
* [pipeline-live](https://github.com/alpacahq/pipeline-live/) is another library
that works well with your Quantopian algorithm. If you are using the Pipeline API, this library can do what you tested in Quantopian and it works well with `pylivetrader`. For more details about how to adjust your Quantopian algorithm, please see [the migration document](./quantopian-to-pipeline-live/).
* [backtrader](https://github.com/backtrader/backtrader/) is a simple but powerful backtesting and live trading framework that lots of algo traders have been using. We have integrated our SDK to backtrader
and you can run your backtrader algo with our API. [backtrader integrated with Alpaca SDK](https://github.com/alpacahq/alpaca-backtrader-api/)
