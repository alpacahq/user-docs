---
title: Watchlist
weight: 45
---

The watchlist API provides CRUD operation for the account's watchlist.
An account can have multiple watchlists and each is uniquely identified by `id`
but can also be addressed by user-defined `name`. Each watchlist is an ordered
list of assets.

{{< rest-endpoint resource="watchlist" method="GET" path="/v2/watchlists" >}}
{{< rest-endpoint resource="watchlist" method="POST" path="/v2/watchlists" >}}

{{< rest-endpoint resource="watchlist" method="GET" path="/v2/watchlists/{watchlist_id}" >}}
{{< rest-endpoint resource="watchlist" method="PUT" path="/v2/watchlists/{watchlist_id}" >}}
{{< rest-endpoint resource="watchlist" method="POST" path="/v2/watchlists/{watchlist_id}" >}}
{{< rest-endpoint resource="watchlist" method="DELETE" path="/v2/watchlists/{watchlist_id}" >}}
{{< rest-endpoint resource="watchlist" method="DELETE" path="/v2/watchlists/{watchlist_id}/{symbol}" >}}

## Endpoints for Watchlist name
You can also call GET, PUT, POST and DELETE with watchlist name with another endpoint `/v2/watchlists:by_name` and query parameter `name=<watchlist_name>`, instead of `/v2/watchlists/{watchlist_id}` endpoints.

## Watchlist Entity

### Example
{{< rest-entity-example name="watchlist">}}

### Properties
{{< rest-entity-desc name="watchlist" >}}
