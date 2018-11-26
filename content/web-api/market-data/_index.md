---
title: Market Data
weight: 120
---

# Market Data

All Alpaca Customers can access various kinds of market data in [Polygon](https://polygon.io/) using
the Alpaca Trade API key ID (but the secret key is not required).  For the RESTful interface, you need
to give the `apiKey` parameter with the same API key ID, as demonstrated below.

```sh
$ curl "https://api.polygon.io/v1/historic/quotes/SPY/2018-06-01?apiKey=$APCA_API_KEY_ID"
```

You can query quotes, bars, and fundamentals data for both historical and real-time.

With the same API key ID, you can subscribe your algorithm to the NATS streaming for market data updates.
In order for your program to be authenticated with the key ID, use the `CONNECT` method with the `auth_token` field.

```
[CONNECT {"auth_token":"<your-key-id>"}]\r\n
```

The NATS server addresses for Alpaca users are not standard ones as shown below.

- nats1.polygon.io:31101
- nats2.polygon.io:31102
- nats3.polygon.io:31103

For the further description and specification of each API endpoint, please find more details [here](https://polygon.io/docs/).
The details about NATS protocol can be found [here](https://nats.io/documentation/internals/nats-protocol/). Also, please
refer to the documentation provided by [each language SDK]({{< relref "/libraries/_index.md" >}}) for its I/O specification.