---
title: Calendar
weight: 50
---

The calendar API serves the full list of market days from 1970 to 2029.
It can also be queried by specifying a start and/or end time to narrow
down the results. In addition to the dates, the response also contains
the specific open and close times for the market days, taking into
account early closures.

{{< rest-endpoint resource="calendar" method="GET" path="/v2/calendar" >}}

## Calendar Entity

### Example
{{< rest-entity-example name="calendar">}}

### Properties
{{< rest-entity-desc name="calendar" >}}
