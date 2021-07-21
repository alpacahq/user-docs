---
title: Alpaca Works With
weight: 90
---

# Alpaca Works With
Alpaca works with many ways to trade! We pride ourselves in enabling many partnerships and allow our users to customize and automate their trade flow. We are constantly working on new ways to enable you to trade. If you are looking to build applications with Alpaca API and OAuth 2.0, please read [OAuth Integration Guide](https://alpaca.markets/docs/build-apps_services-with-alpaca/oauth-guide/) to start building.


## Integrated Partner Apps

**Check out [Integrations page](https://alpaca.markets/integrations) to see the full list of apps that work with Alpaca through OAuth 2.0**

* TradingView ([How it works](./tradingview), [Partner page](https://tradingview.com/broker/Alpaca)): Connect with TradingView and unlock advanced charting to trade commission-free. TradingView is loved by more than 11 million active traders globally.
* Blueshift ([How it works](./blueshift)): Write trading strategies on this web-based backtesting app to backtest them and run them in live trading instantly.
* Streak ([How it works](https://blog.streak.world/2020/06/22/introducing-alpaca-on-streak-with-0-commission-trading/)): No coding but wanna do algorithmic trading? No problem if you use Streak. Streak supports over 300 thousand traders globally doing 27+ million backtests.
* QuantoRocket ([Partner page](https://www.quantrocket.com/alpaca/)): Research and trade automated strategies with this professional-grade platform that caters to data-driven quants, featuring multiple backtesters and easy deployment to the cloud. 
* Tradetron ([How it works](https://www.youtube.com/watch?v=nzLcNJDeP74&feature=youtu.be)): Allows you to build algo strategies without coding and allow others to subscribe your algos in their own linked brokerage accounts automatically.
* Arcade Trader ([How it works](https://medium.com/@arcade_trader/algorithmic-trading-for-beginners-part-1-d6589d4beb05)): Originally designed for hedge funds, Arcade Trader gives you access to advanced algorithmic trading research tools.
* IFTTT ([How it works](./ifttt), [Partner page](https://ifttt.com/alpaca)): Use triggers to build “if this, then that” services that launch trades or respond to them. IFTTT is trusted by over 18 million consumers and 130 thousand developers accross the world.
* Zapier ([Partner page](https://zapier.com/apps/alpaca/integrations)): Connect to thousands of apps to automate your trading flow. Join the community of more than 3 million people that rely on Zapier.
* Slack ([How it works](https://alpaca.markets/docs/alpaca-works-with/alpaca-for-slack/)): Quickly manage your Alpaca brokerage account right inside Slack. With simple commands, you can have AlpacaBot get your orders and positions, and you can even tell it to place orders for you.
* PredictNow.ai ([How it works](./predictnow)): A no-code machine learning SaaS helping traders apply machine learning to their investment strategies in order to predict the profitability of their next trade.
* MetaTrader 5 ([Partner page](https://www.metatrader5.com/en/terminal/help/startworking/acc_open)): Convenient and Functional Trading Platform featuring technical analysis, fundamental analysis, automated trading, and trading from mobile devices.
* Algo Bulls ([Partner page](https://help.algobulls.com/broker/alpaca/)): Make use of powerful features like Live Trading, Paper Trading, Backtesting, Strategy Customizations, Python Build, and much more directly from the website or Android App.
* Trellis ([How it works](https://www.trytrellis.co/)): Build a trading bot without code, connect with social integrations, and compete against other users in both paper and live trading all free of charge.
* PredictNow.ai ([How it works](https://www.predictnow.ai/)): A no-code machine learning SaaS, helping traders apply machine learning to their investment strategies in order to predict the profitability of their next trade.
            


## Open Source Integration Libraries
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
