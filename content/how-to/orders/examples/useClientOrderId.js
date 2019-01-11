const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca()

// Submit a market order and assign it a Client Order ID.
alpaca.createOrder({
    symbol: 'AAPL',
    qty: 1,
    side: 'buy',
    type: 'market',
    time_in_force: 'day',
    client_order_id='my_first_order'
})

// Get our order using its Client Order ID.
alpaca.getOrderByClientOrderId('my_first_order')
    .then((myOrder) => {
        console.log(`Got order #${myOrder.id}.`)
    })