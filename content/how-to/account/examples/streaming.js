const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca({
    paper: true
})

// Prepare the websocket connection and subscribe to account_updates.
let websocket = alpaca.websocket
client.onConnect(function () {
    console.log("Connected")
    client.subscribe(['account_updates'])
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

// Handle incoming account updates.
const client_order_id = 'my_client_order_id'
client.onAccountUpdate(data => {
    // Track the cash balance in our account.
    const balance = data['cash']
    console.log(`Update for account. Cash balance: ${balance}`)
})

// Start listening for updates.
client.connect()