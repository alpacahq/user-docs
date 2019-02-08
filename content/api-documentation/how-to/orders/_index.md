---
title: Order Examples
weight: 10
---

This page contains examples of some of the things you can do with order objects through our API. For additional help
understanding different types of orders and how they behave once they're placed, please see [this page](https://docs.alpaca.markets/orders/).

---

## Place New Orders
Orders can be placed with a POST request to our [/v1/orders](https://docs.alpaca.markets/api-documentation/web-api/orders/#request-a-new-order) endpoint.

{{< code-example exampleId="placeOrder" pathURL="/api-documentation/how-to/orders/examples/" >}}

## Use Client Order IDs
Client Order IDs can be used to organize and track specific orders in your client program.

{{< code-example exampleId="useClientOrderId" pathURL="/api-documentation/how-to/orders/examples/">}}

## Get a List of Existing Orders
If you'd like to see a list of your existing orders, you can send a get request to our [/v1/orders](https://docs.alpaca.markets/api-documentation/web-api/orders/#get-a-list-of-orders) endpoint.

{{< code-example exampleId="getOrders" pathURL="/api-documentation/how-to/orders/examples">}}

## Listen for Updates to Orders
You can use Websockets to receive real-time updates about the status of your orders as they change. You can see the full documentation [here](https://docs.alpaca.markets/api-documentation/web-api/streaming/#order-updates).

{{< code-example exampleId="streaming" pathURL="/api-documentation/how-to/orders/examples">}}
