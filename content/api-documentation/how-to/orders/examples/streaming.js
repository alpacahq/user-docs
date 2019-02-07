const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca()

// Prepare the websocket connection and subscribe to trade_updates.
alpaca.websocket.onConnect(function () {
    console.log("Connected")
    client.subscribe(['trade_updates'])
    setTimeout(() => {
        client.disconnect()
    }, 30 * 1000)
})
alpaca.websocket.onDisconnect(() => {
    console.log("Disconnected")
})
alpaca.websocket.onStateChange(newState => {
    console.log(`State changed to ${newState}`)
})

// Handle updates on an order you've given a Client Order ID.
const client_order_id = 'my_client_order_id'
alpaca.websocket.onOrderUpdate(orderData => {
    const event = orderData['event']
    if (orderData['order']['client_order_id'] == client_order_id) {
        console.log(`Update for ${client_order_id}. Event: ${event}.`)
    }
})

// Start listening for updates.
alpaca.websocket.connect()