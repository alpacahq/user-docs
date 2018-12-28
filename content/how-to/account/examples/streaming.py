import alpaca_trade_api as tradeapi

conn = tradeapi.stream2.StreamConn()

# Handle incoming account updates.
# The r indicates that we're listening for a regex pattern.
@conn.on(r'.*')
async def on_msg(conn, channel, data):
    # Track the cash balance in our account.
    print("Update for account. Cash balance: {}".format(data['cash']))

# Start listening for updates.
conn.run(['account_updates'])