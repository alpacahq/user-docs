import alpaca_trade_api as tradeapi

conn = tradeapi.stream2.StreamConn()

# We'll listen for price updates on AAPL's 1 minute bars.
streams = ['bars/AAPL/1Min']

# Handle incoming market data updates.
# The r indicates that we're listening for a regex pattern.
@conn.on(r'bars\/AAPL\/1Min')
async def on_msg(conn, channel, data):
    if channel in streams:
        print('Got price update. Latest close price: ', data.c)

conn.run(streams)