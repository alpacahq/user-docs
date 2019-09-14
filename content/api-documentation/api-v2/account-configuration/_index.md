---
title: Account Configurations
weight: 15
---

# Account Configurations API

The account configuration API provides custom configurations about the
trading account. These configurations control various behaviors of the
account to meet your trading needs.

{{< rest-endpoint resource="account-configurations" method="GET" path="/v2/account/configurations" >}}
{{< rest-endpoint resource="account-configurations" method="PATCH" path="/v2/account/configurations" >}}

## Account Configurations Entity

### Example
{{< rest-entity-example name="account-configurations-v2">}}

### Properties
{{< rest-entity-desc name="account-configurations-v2">}}

For DTMC protection, see [Day Trade Margin Call Protection]({{< ref path="/user-protections/_index.md#day-trade-margin-call-dtmc-protection-at-alpaca" >}})