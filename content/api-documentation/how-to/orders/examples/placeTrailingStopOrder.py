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

# Submit a trailing stop order to sell 1 share of Apple at a
# trailing stop of
api.submit_order(
    symbol='AAPL',
    qty=1,
    side='sell',
    type='trailing_stop',
    trail_price=1.00,  # stop price will be hwm - 1.00$
    time_in_force='gtc',
)

# Alternatively, you could use trail_percent:
api.submit_order(
    symbol='AAPL',
    qty=1,
    side='sell',
    type='trailing_stop',
    trail_percent=1.0,  # stop price will be hwm*0.99
    time_in_force='gtc',
)
