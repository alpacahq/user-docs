const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca()

// Get our position in AAPL.
aaplPosition = alpaca.getPosition('AAPL')

// Get a list of all of our positions.
alpaca.getPositions()
    .then((portfolio) => {
        // Print the quantity of shares for each position.
        portfolio.forEach(function (position) {
            console.log(`${position.qty} shares of ${position.symbol}`)
        })
    })

