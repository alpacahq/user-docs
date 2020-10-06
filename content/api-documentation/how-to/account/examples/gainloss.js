const Alpaca = require('@alpacahq/alpaca-trade-api')
const alpaca = new Alpaca()

// Get account information.
alpaca.getAccount().then((account) => {
	// Calculate the difference between current balance and balance at the last market close.
    const balanceChange = account.equity - account.last_equity

    console.log('Today\'s portfolio balance change:', balanceChange)
})
