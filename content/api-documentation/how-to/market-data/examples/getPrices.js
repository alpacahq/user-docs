const moment = require("moment");
const Alpaca = require("@alpacahq/alpaca-trade-api");
const alpaca = new Alpaca();

//Get daily price data for AAPL over the last 7 days.
let bars = alpaca.getBarsV2(
  "AAPL",
  {
    start: moment().subtract(7, "days").format(),
    end: moment().subtract(20, "minutes").format(),
    timeframe: "1Day",
  },
  alpaca.configuration
);
const barset = [];

for await (let b of bars) {
  barset.push(b);
}

const week_open = barset[0].OpenPrice;
const week_close = barset.slice(-1)[0].ClosePrice;
const percent_change = ((week_close - week_open) / week_open) * 100;

console.log(`AAPL moved ${percent_change}% over the last 7 days`);
