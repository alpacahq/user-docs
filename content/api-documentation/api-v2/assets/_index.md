---
title: Assets
weight: 40
---

# Assets API
The assets API serves as the master list of assets available for trade and data
consumption from Alpaca. Assets are sorted by asset class, exchange and symbol.
Some assets are only available for data consumption via Polygon, and are not
tradable with Alpaca. These assets will be marked with the flag
`tradable=false`.

{{< rest-endpoint resource="assets" method="GET" path="/v2/assets" >}}
{{< rest-endpoint resource="assets" method="GET" path="/v2/assets/:id" >}}
{{< rest-endpoint resource="assets" method="GET" path="/v2/assets/{symbol}" >}}

## Asset Entity

### Example
{{< rest-entity-example name="asset-v2">}}

### Properties
{{< rest-entity-desc name="asset-v2" >}}
