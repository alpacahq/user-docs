import alpaca_trade_api as tradeapi

api = tradeapi.REST()

# Check if the market is open now.
clock = api.get_clock()

if clock.is_open:
    print('The market is open for trading.')

# Check when the market was open on Dec. 1, 2018
calendar = api.get_calendar(start='2018-12-01')
print('The market opened at {} and closed at {}'.format(
    calendar.open,
    calendar.close
))