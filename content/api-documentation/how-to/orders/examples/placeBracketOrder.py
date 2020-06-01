import alpaca_trade_api as tradeapi

api = tradeapi.REST()

symbol = 'AAPL'
symbol_bars = api.get_barset(symbol, 'minute', 1).df.iloc[0]
symbol_price = symbol_bars[symbol]['close']

# We could buy a position and add a stop-loss and a take-profit of 5 %
api.submit_order(
    symbol=symbol,
    qty=1,
    side='buy',
    type='market',
    time_in_force='gtc',
    order_class='bracket',
    stop_loss={'stop_price': symbol_price * 0.95,
               'limit_price':  symbol_price * 0.94},
    take_profit={'limit_price': symbol_price * 1.05}
)

# We could buy a position and just add a stop loss of 5 % (OTO Orders)
api.submit_order(
    symbol=symbol,
    qty=1,
    side='buy',
    type='market',
    time_in_force='gtc',
    order_class='oto',
    stop_loss={'stop_price': symbol_price * 0.95}
)

# We could split it tp 2 orders. first buy a stock,
# and then add the stop/profit prices (OCO Orders)
api.submit_order(symbol, 1, 'buy', 'limit', 'day', symbol_price)

# wait for it to buy position and then
api.submit_order(
    symbol=symbol,
    qty=1,
    side='sell',
    type='limit',
    time_in_force='gtc',
    order_class='oco',
    stop_loss={'stop_price': symbol_price * 0.95},
    take_profit={'limit_price': symbol_price * 1.05}
)
