import alpaca_trade_api as tradeapi

api = tradeapi.REST()

# Get a list of all active assets.
active_assets = api.list_assets(status='active')

# Filter the assets down to just those on NASDAQ.
nasdaq_assets = [a for a in active_assets if a.exchange == 'NASDAQ']
print(nasdaq_assets)