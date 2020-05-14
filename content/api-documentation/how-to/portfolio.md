---
title: Portfolio Examples
weight: 10
---

## View Open Positions in Your Portfolio
You can view the positions in your portfolio by making a GET request to our [`/v2/positions`]({{< relref "/api-documentation/api-v2/positions.md" >}}) endpoint. If you specify a symbol, you'll see only your position for the associated stock.

{{< code-example exampleId="getPositions" pathURL="/api-documentation/how-to/portfolio/examples">}}