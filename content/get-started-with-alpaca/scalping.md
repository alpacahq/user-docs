---
title: Concurrent Scalping Algo Using Async Python
weight: 70
---

*This example only works if you have a funded brokerage account or another means of accessing Polygon data.*

This python script is a working example to execute scalping trading algorithm for Alpaca API. This algorithm uses real time order updates as well as minute level bar streaming from Polygon via Websockets (see the document for Polygon data access). One of the contributions of this example is to demonstrate how to handle multiple stocks concurrently as independent routine using Python's asyncio.

[Detailed Explanation on Medium](https://medium.com/automation-generation/concurrent-scalping-algo-using-async-python-8df9f31e22f1)

[Code on GitHub](https://github.com/alpacahq/example-scalping)
