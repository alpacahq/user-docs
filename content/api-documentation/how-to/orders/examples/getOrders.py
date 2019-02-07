import alpaca_trade_api as tradeapi

api = tradeapi.REST()

# Get the last 100 of our closed orders
closed_orders = api.list_orders(
    status='closed',
    limit=100
)

# Get only the closed orders for a particular stock
closed_aapl_orders = [o for o in closed_orders if o.symbol is 'AAPL']
print(closed_aapl_orders)

