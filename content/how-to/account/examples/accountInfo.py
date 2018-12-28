import alpaca_trade_api as tradeapi

api = tradeapi.REST()

# Get our account information.
account = api.get_account()

# Check if our account is restricted from trading.
can_trade = not account.trading_blocked

# Check how much money we can use to open new positions.
buying_power = account.buying_power