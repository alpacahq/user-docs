const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca()

// Prepare the websocket connection and subscribe to trade_updates.
let websocket = alpaca.websocket
client.onConnect(function () {
    console.log("Connected")
    client.subscribe(['trade_updates'])
    setTimeout(() => {
        client.disconnect()
    }, 30 * 1000)
})
client.onDisconnect(() => {
    console.log("Disconnected")
})
client.onStateChange(newState => {
    console.log(`State changed to ${newState}`)
})

// Handle updates on an order you've given a Client Order ID.
const client_order_id = 'my_client_order_id'
client.onOrderUpdate(data => {
    const event = data['event']
    if (data['order']['client_order_id'] == client_order_id) {
        console.log(`Update for ${client_order_id}. Event: ${event}.`)
    }
})

// Start listening for updates.
client.connect()