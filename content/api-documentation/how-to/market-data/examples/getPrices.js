const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca()

// Get daily price data for AAPL over the last 5 trading days.
const barset = alpaca.getBars(
    'day',
    'AAPL',
    {
        limit: 5
    }
).then((barset) => {
    const aapl_bars = barset['AAPL']

    // See how much AAPL moved in that timeframe.
    const week_open = aapl_bars[0].o
    const week_close = aapl_bars.slice(-1)[0].c
    const percent_change = (week_close - week_open) / week_open * 100

    console.log(`AAPL moved ${percent_change}% over the last 5 days`)
})
