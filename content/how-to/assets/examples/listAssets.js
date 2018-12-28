const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca({
    paper: true
})

// Get a list of all active assets.
const activeAssets = alpaca.getAssets({
    status: 'active'
}).then((activeAssets) => {
    // Filter the assets down to just those on NASDAQ.
    const nasdaqAssets = activeAssets.filter(asset => asset.exchange == 'NASDAQ')
    console.log(nasdaqAssets)
})
