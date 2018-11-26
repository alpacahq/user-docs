---
title: Libraries
weight: 70
---
# Libraries
From client SDK to trading system integration, we provide many open source
libraries and code snippets to fit your unique use case. Here are the
list of our official libraries. If there is not one that fits your needs,
let us know and we will work on that.

## Client SDK
These are the basic SDK for each lanuage program to interact with
our Web API.

- Python: [alpaca-trade-api-python](https://github.com/alpacahq/alpaca-trade-api-python/) /
[PyPI](https://pypi.org/project/alpaca-trade-api/)
- .NET/C#: [alpaca-trade-api-charp](https://github.com/alpacahq/alpaca-trade-api-csharp/) /
[NuGet](https://www.nuget.org/packages/Alpaca.Markets/)
- Go: [alpaca-trade-api-go](https://github.com/alpacahq/alpaca-trade-api-go/)
- Node: [alpaca-trade-api-js](https://github.com/alpacahq/alpaca-trade-api-js/)

## Python

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

## .NET/C\#

### Lean
We have integrated with [Lean](https://github.com/QuantConnect/Lean/)
which is a big player in the .NET algo trading space. If you are
already testing your algorithm in QuantConnect, this is the tool
you can use to run it in live with our API.
