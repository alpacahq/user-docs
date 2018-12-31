const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca()

// Prepare the websocket connection and subscribe to account_updates.
alpaca.websocket.onConnect(function () {
    console.log("Connected")
    client.subscribe(['account_updates'])
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

// Handle incoming account updates.
alpaca.websocket.onAccountUpdate(accountDdata => {
    // Track the cash balance in our account.
    const balance = accountDdata['cash']
    console.log(`Update for account. Cash balance: ${balance}`)
})

// Start listening for updates.
alpaca.websocket.connect()