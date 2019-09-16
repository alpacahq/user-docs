---
title: Positions
weight: 30
---

The positions API provides information about an account's current open positions. The response will include information such as cost basis, shares traded, and market value, which will be updated live as price information is updated. Once a position is closed, it will no longer be queryable through this API.

{{< rest-endpoint resource="positions" method="GET" path="/v2/positions" >}}
{{< rest-endpoint resource="positions" method="GET" path="/v2/positions/{symbol}" >}}
{{< rest-endpoint resource="positions" method="DELETE" path="/v2/positions" >}}
{{< rest-endpoint resource="positions" method="DELETE" path="/v2/positions/{symbol}" >}}

## Position Entity

### Example
{{< rest-entity-example name="position">}}

### Properties
{{< rest-entity-desc name="position">}}
