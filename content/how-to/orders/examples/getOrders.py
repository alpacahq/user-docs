import alpaca_trade_api as tradeapi

api = tradeapi.REST()

# Get the last 100 of our closed orders
orders = api.list_orders(
    status='closed',
    limit=100
)

# Get only the orders for a particular stock
aapl_orders = [o for o in orders if o.symbol is 'AAPL']

