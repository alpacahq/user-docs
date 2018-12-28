const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca({
    paper: true
})

// Check if AAPL is tradable on the Alpaca platform.
alpaca.getAsset('AAPL')
    .then((aaplAsset) => {
        if (aaplAsset.tradable) {
            console.log('We can trade AAPL.')
        }
    })