---
title: Account Activities
weight: 15
---

The account activities API provides access to a historical record of transaction activities that have impacted your account. Trade execution activities and non-trade activities, such as dividend payments, are both reported through this endpoint. See the bottom of this page for a full list of the types of activities that may be reported.

{{< rest-endpoint resource="account-activities" method="GET" path="/v2/account/activities/{activity_type}" >}}

{{< rest-endpoint resource="account-activities" method="GET" path="/v2/account/activities" >}}

## Trade Account Activity Entity

### Example
{{< rest-entity-example name="account-trade-activity-v2">}}

### Properties
{{< rest-entity-desc name="account-trade-activity-v2">}}

## Nontrade Account Activity Entity

### Example
{{< rest-entity-example name="account-nontrade-activity-v2">}}

### Properties
{{< rest-entity-desc name="account-nontrade-activity-v2">}}

## Activity Types

* `FILL`: Order fills (both partial and full fills)
* `TRANS`: Cash transactions (both CSD and CSR)
* `MISC`: Miscellaneous or rarely used activity types (All types except those in TRANS, DIV, or FILL)
* `ACATC`: ACATS IN/OUT (Cash)
* `ACATS`: ACATS IN/OUT (Securities)
* `CSD`: Cash disbursement(+)
* `CSR`: Cash receipt(-)
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
* `INTNRA` Interest adjusted (NRA Withheld)
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