---
title: Account Examples
weight: 10
---

## View Account Information
By sending a GET request to our [`/v1/account`](https://docs.alpaca.markets/api-documentation/web-api/account/) endpoint, you can see various information about your account, such as the amount of buying power available or whether or not it has a [PDT flag](https://support.alpaca.markets/hc/en-us/articles/360012203032-Pattern-Day-Trader).

{{< code-example exampleId="accountInfo" pathURL="/api-documentation/how-to/account/examples">}}

## Listen for Account Updates with Websockets
You can use Websockets to receive real-time updates about the status of your account. You can see the full documentation [here](https://docs.alpaca.markets/web-api/streaming/#account-updates).

{{< code-example exampleId="streaming" pathURL="/api-documentation/how-to/account/examples">}}