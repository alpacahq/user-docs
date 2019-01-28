import alpaca_trade_api as tradeapi

api = tradeapi.REST()

# Submit a market order to buy 1 share of Apple at market price
api.submit_order(
    symbol='AAPL',
    qty=1,
    side='buy',
    type='market',
    time_in_force='gtc'
)

# Submit a limit order to attempt to sell 1 share of AMD at a
# particular price ($20.50) when the market opens
api.submit_order(
    symbol='AMD',
    qty=1,
    side='sell',
    type='limit',
    time_in_force='opg',
    limit_price=20.50
)