import alpaca_trade_api as tradeapi

api = tradeapi.REST()

# Check if the market is open now.
clock = api.get_clock()
print('The market is {}'.format('open.' if clock.is_open else 'closed.'))

# Check when the market was open on Dec. 1, 2018
date = '2018-12-01'
calendar = api.get_calendar(start=date, end=date)
print('The market opened at {} and closed at {} on {}.'.format(
    calendar.open,
    calendar.close,
    date
))