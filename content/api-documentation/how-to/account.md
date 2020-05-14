---
title: Account Examples
weight: 10
---

## View Account Information
By sending a GET request to our [`/v2/account`]({{< relref "/api-documentation/api-v2/account.md" >}}) endpoint, you can see various information about your account, such as the amount of buying power available or whether or not it has a [PDT flag](https://alpaca.markets/learn/pattern-day-trader/).

{{< code-example exampleId="accountInfo" pathURL="/api-documentation/how-to/account/examples">}}

## View Gain/Loss of Portfolio
You can use the information from the account endpoint to do things like calculating the daily profit or loss of your account.

{{< code-example exampleId="gainloss" pathURL="/api-documentation/how-to/account/examples">}}

