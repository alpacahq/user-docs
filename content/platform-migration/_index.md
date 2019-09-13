---
title: Third-Party Tools Integration
weight: 90
---
# Third-Party Tools Integration

This page is meant to serve as a guide to existing algo traders and researchers
currently using other platforms or brokers including but not limited to Quantopian,
Interactive Brokers, TD Ameritrade, Robinhood, and Backtrader.

We provide several open source libraries to aid in the platform integration process and
help expedite the time it takes for you to get up and running with Alpaca. Below is a list of
our official platform integration libraries. If you would like to see us provide integration support
for a platform not listed here, please let us know at our Feature Request and Issues repo
[here](https://github.com/alpacahq/Alpaca-API).

### Python

### pylivetrader
[pylivetrader](https://github.com/alpacahq/pylivetrader/) is a live trading framework in
python 3 that is API-compatible with [zipline](https://github.com/quantopian/zipline/).
You can run your Quantopian algorithm with the minimum code change with Alpaca API.

For more details about how to convert your Quantopian algorithm, please see
[zipline to pylivetrader migration](./zipline-to-pylivetrader/).

### pipeline-live
[pipeline-live](https://github.com/alpacahq/pipeline-live/) is another library
that works well with your Quantopian algorithm. If you are using the Pipeline API,
this library can do what you tested in Quantopian and it works well with
`pylivetrader`.

For more details about how to adjust your Quantopian algorithm, please see
[the migration document](./quantopian-to-pipeline-live/).


### backtrader
[backtrader](https://github.com/backtrader/backtrader/) is a simple but
powerful backtesting and live trading framework that lots of algo traders
have been using for real use. We have integrated our SDK to backtrader
today and you can run your backtrader algo with our API.

[backtrader integrated with Alpaca SDK](https://github.com/alpacahq/alpaca-backtrader-api/)

