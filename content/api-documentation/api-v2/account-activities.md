---
title: Account Activities
weight: 75
---

The account activities API provides access to a historical record of transaction activities that have impacted your account. Trade execution activities and non-trade activities, such as dividend payments, are both reported through this endpoint. See the bottom of this page for a full list of the types of activities that may be reported.

### Paging of Results
Pagination is handled using the `page_token` and `page_size` parameters. `page_token` represents the ID of the end of your current page of results. If specified with a direction of `desc`, for example, the results will end before the activity with the specified ID. If specified with a direction of `asc`, results will begin with the activity immediately after the one specified. `page_size` is the maximum number of entries to return in the response. If `date` is not specified, the default and maximum value is 100. If `date` is specified, the default behavior is to return all results, and there is no maximum page size.

{{< rest-endpoint resource="account-activities" method="GET" path="/v2/account/activities/{activity_type}" >}}

{{< rest-endpoint resource="account-activities" method="GET" path="/v2/account/activities" >}}

## TradeActivity Entity

### Example
{{< rest-entity-example name="account-trade-activity-v2">}}

### Properties
{{< rest-entity-desc name="account-trade-activity-v2">}}

## NonTradeActivity Entity

### Example
{{< rest-entity-example name="account-nontrade-activity-v2">}}

### Properties
{{< rest-entity-desc name="account-nontrade-activity-v2">}}

## Activity Types

* `FILL`: Order fills (both partial and full fills)
* `TRANS`: Cash transactions (both CSD and CSW)
* `MISC`: Miscellaneous or rarely used activity types (All types except those in TRANS, DIV, or FILL)
* `ACATC`: ACATS IN/OUT (Cash)
* `ACATS`: ACATS IN/OUT (Securities)
* `CSD`: Cash deposit(+)
* `CSW`: Cash withdrawal(-)
* `DIV`: Dividends
* `DIVCGL`: Dividend (capital gain long term)
* `DIVCGS`: Dividend (capital gain short term)
* `DIVFEE`: Dividend fee
* `DIVFT`: Dividend adjusted (Foreign Tax Withheld)
* `DIVNRA`: Dividend adjusted (NRA Withheld)
* `DIVROC`: Dividend return of capital
* `DIVTW`: Dividend adjusted (Tefra Withheld)
* `DIVTXEX`: Dividend (tax exempt)
* `INT`: Interest (credit/margin)
* `INTNRA`: Interest adjusted (NRA Withheld)
* `INTTW`: Interest adjusted (Tefra Withheld)
* `JNL`: Journal entry
* `JNLC`: Journal entry (cash)
* `JNLS`: Journal entry (stock)
* `MA`: Merger/Acquisition
* `NC`: Name change
* `OPASN`: Option assignment
* `OPEXP`: Option expiration
* `OPXRC`: Option exercise
* `PTC`: Pass Thru Charge
* `PTR`: Pass Thru Rebate
* `REORG`: Reorg CA
* `SC`: Symbol change
* `SSO`: Stock spinoff
* `SSP`: Stock split
