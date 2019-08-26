import alpaca_trade_api as tradeapi

if __name__ == '__main__':
    """
    With the Alpaca API, you can check on your daily profit or loss by
    comparing your current balance to yesterday's balance.
    """

    # First, open the API connection
    api = tradeapi.REST(
        'your_api_key_here',
        'your_api_secret_here',
        'https://paper-api.alpaca.markets'
    )

    # Get account info
    account = api.get_account()

    # Check our current balance vs. our balance at the last market close
    balance_change = float(account.equity) - float(account.last_equity)
    print(f'Today\'s portfolio balance change: ${balance_change}')