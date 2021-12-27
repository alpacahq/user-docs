---
title: Announcements
weight: 50
---

The announcements endpoint contains public information on previous and upcoming dividends, mergers, spinoffs, and stock splits. 

Announcement data is made available through the API as soon as it is ingested by Alpaca, which is typically the following trading day after the declaration date. This provides insight into future account balance changes that will take effect on an announcement’s payable date. Additionally, viewing previous announcement details can improve bookkeeping and reconciling previous account stock position and cash balance changes.

{{< rest-endpoint resource="announcements" method="GET" path="/v2/corporate_actions/announcements" >}}

## Announcement Entity

### Example
{{< rest-entity-example name="announcements">}}

### Properties
{{< rest-entity-desc name="announcements" >}}
