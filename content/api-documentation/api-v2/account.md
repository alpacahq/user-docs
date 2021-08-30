---
title: Account
weight: 10
---

The account API serves important information related to an account,
including account status, funds available for trade, funds available for
withdrawal, and various flags relevant to an account's ability to trade.
An account maybe be blocked for just for trades (`trades_blocked` flag) or for both
trades and transfers (`account_blocked` flag) if Alpaca identifies the account to
engaging in any suspicious activity. Also, in accordance with FINRA's pattern day
trading rule, an account may be flagged for pattern day trading
(`pattern_day_trader flag`), which would inhibit an account from placing any
further day-trades. Please note that cryptocurrencies are not eligible assets to be
used as collateral for margin accounts and will require the asset be traded using 
cash only.

{{< rest-endpoint resource="account" method="GET" path="/v2/account" >}}

## Account Entity

### Example
{{< rest-entity-example name="account-v2">}}

### Properties
{{< rest-entity-desc name="account-v2">}}

## Account Status
The following are the possible account status values. Most likely, the
account status is `ACTIVE` unless there is any problem. The account status
may get in `ACCOUNT_UPDATED` when personal information is being updated
from the dashboard, in which case you may not be allowed trading for
a short period of time until the change is approved.

- `ONBOARDING`  
  The account is onboarding.
- `SUBMISSION_FAILED`  
  The account application submission failed for some reason.
- `SUBMITTED`  
  The account application has been submitted for review.
- `ACCOUNT_UPDATED`  
  The account information is being updated.
- `APPROVAL_PENDING`  
  The final account approval is pending.
- `ACTIVE`  
  The account is active for trading.
- `REJECTED`  
  The account application has been rejected.
