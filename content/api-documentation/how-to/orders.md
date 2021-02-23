---
title: Order Examples
weight: 10
---

This page contains examples of some of the things you can do with order objects through our API. For additional help
understanding different types of orders and how they behave once they're placed, please see the [order page]({{< relref "/trading-on-alpaca/orders.md" >}})

---

## Place New Orders
Orders can be placed with a POST request to our [/v2/orders]({{< relref "/api-documentation/api-v2/orders.md#request-a-new-order" >}}) endpoint.

{{< code-example exampleId="placeOrder" pathURL="/api-documentation/how-to/orders/examples" >}}

## Submit Short Orders
Short orders can also be placed for securities which you do not hold an open long position in.

{{< code-example exampleId="shortOrder" pathURL="/api-documentation/how-to/orders/examples" >}}

## Use Client Order IDs
Client Order IDs can be used to organize and track specific orders in your client program.

{{< code-example exampleId="useClientOrderId" pathURL="/api-documentation/how-to/orders/examples">}}

## Submit Bracket Orders
Bracket orders allow you to create a chain of orders that react to execution
 and stock price.
For more details, go to [Bracket Order Overview]({{< relref"/trading-on-alpaca/orders.md#bracket-orders" >}})

{{< code-example exampleId="placeBracketOrder" pathURL="/api-documentation/how-to/orders/examples">}}

## Submit Trailing Stop Orders
Trailing stop orders allow you to create a stop order that automatically changes the stop price allowing 
you to maximize your profits while still protecting your position with a stop price.
For more details, go to [Trailing Stop Order Overview]({{< relref"/trading-on-alpaca/orders.md#trailing-stop-orders" >}})

{{< code-example exampleId="placeTrailingStopOrder" pathURL="/api-documentation/how-to/orders/examples">}}

## Get a List of Existing Orders
If you'd like to see a list of your existing orders, you can send a get request to our [/v2/orders]({{< relref "/api-documentation/api-v2/orders.md#get-a-list-of-orders" >}}) endpoint.

{{< code-example exampleId="getOrders" pathURL="/api-documentation/how-to/orders/examples">}}

## Listen for Updates to Orders
You can use Websockets to receive real-time updates about the status of your orders as they change. You can see the full documentation [here]({{< relref "/api-documentation/api-v2/streaming.md#order-updates" >}}).

{{< code-example exampleId="streaming" pathURL="/api-documentation/how-to/orders/examples">}}
