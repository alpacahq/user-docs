import alpaca_trade_api as tradeapi

api = tradeapi.REST()

# Check if AAPL is tradable on the Alpaca platform.
aapl_asset = api.get_asset('AAPL')
if aapl_asset.tradable:
    print('We can trade AAPL.')