const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca()

// Get our account information.
alpaca.getAccount()
    .then((account) => {
        // Check if our account is restricted from trading.
        if (account.trading_blocked) {
            console.log('Account is currently restricted from trading.')
        }

        // Check how much money we can use to open new positions.
        console.log(`$${account.buying_power} is available as buying power.`)
    })