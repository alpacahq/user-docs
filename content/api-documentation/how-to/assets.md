---
title: Assets Examples
weight: 10
---

## Get a List of Assets
If you send a GET request to our [`/v1/assets`](https://docs.alpaca.markets/api-documentation/web-api/assets/#get-assets) endpoint, you'll receive a list of US equities.

{{< code-example exampleId="listAssets" pathURL="/api-documentation/how-to/assets/examples">}}

## See If a Particular Asset is Tradable on Alpaca
By sending a symbol along with our request, we can get the information about just one asset. This is useful if we just want to make sure that a particular asset is tradable before we attempt to buy it.

{{< code-example exampleId="checkAsset" pathURL="/api-documentation/how-to/assets/examples">}}
