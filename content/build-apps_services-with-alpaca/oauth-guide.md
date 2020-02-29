---
title: OAuth Integration Guide
weight: 10
---

## <a name="introduction"></a>Introduction

We welcome developers to build applications and products that are powered by Alpaca while also protecting the privacy and security of our users. To build using Alpaca's APIs, please follow the guide below.

Alpaca implements OAuth 2.0 to allow third party applications to access Alpaca Trading API on
behalf of the end-users. This document describes how you can integrate with Alpaca through OAuth.

*Note that these materials are supplementary to our API documentation. Please read the [API documentation](https://docs.alpaca.markets/api-documentation/api-v2/) for topics outside of OAuth integration.*

* [Terms of Access and Use]({{< relref "#terms" >}})
* [Registering Your Application]({{< relref "#registration" >}})
* [Integration]({{< relref "#integration" >}})
* [Making API Calls]({{< relref "#APIcall" >}})
* [How to Use Polygon with OAuth]({{< relref "#how-to-use-polygon-with-oauth" >}})

## <a name="terms"></a>Terms of Access and Use

* You must read the terms and register in order to connect and use Alpaca's APIs
* All API clients must authenticate with OAuth 2.0
* You may not imply that the app was developed by Alpaca.
* If you are building a commercial application that makes money (including ads, in-app purchases, etc), you must disclose it in the registration form and receive written approval.

## <a name="registration"></a>Application Creation and Registration

* Before integrating with Alpaca, you'll first need to create a new OAuth app under your [OAuth Apps](https://app.alpaca.markets/brokerage/apps/manage) page.

* To register your app, click the 'OAuth Apps' icon on the left side menu. From here, you can access the apps you have authorized as well as the apps you have developed.

<center><img src="./apps_page.png" width="80%"></center>

* Use the "Create New App" flow to register your new application. The first step will be to agree and sign our API Agreement.

* Once you add your relevant information and create the app, you will receive your Client ID and Client Secret.

<center><img src="./application_info.png" width="60%"></center>
<center><img src="./client_id_secret.png" width="30%"></center>

## <a name="integration"></a>Integration

Once the application is registered, we will issue a Client ID and Client Secret for it.

To integrate your application with Alpaca, use the following flow:

<center><img src="./Flow.png" width="80%"></center>

### 1. ***End user requests service from application. Application redirect users to request Alpaca access***

```
GET https://app.alpaca.markets/oauth/authorize?response_type=code&client_id=YOUR_CLIENT_ID&redirect_uri=YOUR_REDIRECT_URL&state=SOMETHING_RANDOM&scope=account:write%20trading
```

When redirecting a user to Alpaca to authorize access to your application, you’ll need to construct the authorization URL with the correct parameters and scopes. Here’s a list of parameters you should always specify:

| Parameter | Description |
| --------- | ----------- |
| `response_type` | **Required** Must be `code` to request an authorization code. |
| `client_id` | **Required** The Client ID you received when you registered the application. |
| `redirect_uri` | **Required** The URL where the user will be sent after authorization. It must match one of the whitelisted redirect URIs for your application. |
| `state` | **Optional** An unguessable random string, used to protect against request forgery attacks.  |
| `scope` | **Optional** A space-delimited list of scopes your application requests access to. Read-only endpoint access is assumed by default. |

**Allowed Scopes**:

| Scope           | Description                                             |
| --------------- | ------------------------------------------------------- |
| `account:write` | Write access for account configurations and watchlists. |
| `trading`       | Place, cancel or modify orders.                         |
| `data`          | Access to the Data API.                                 |



Example authorization URL:

```
GET https://app.alpaca.markets/oauth/authorize?response_type=code&client_id=fc9c55efa3924f369d6c1148e668bbe8&redirect_uri=https%3A%2F%2Fexample.com%2Foauth%2Fcallback&state=8e02c9c6a3484fadaaf841fb1df290e1&scope=account:write%20trading
```

### 2. ***End user authorizes API access for the applications***

From the user side, they will see the following authorization screen:

<center><img src="./Authorization_Page.png" width="40%"></center>

### 3. ***Alpaca redirects end user to application with an authorization code.***

If the user approves access, Alpaca will redirect them back to your `redirect_uri` with a temporary `code` parameter. If you specified a state parameter in step 1, it will be returned as well. The parameter will always match the value specified in step 1. If the values don’t match, the request should not be trusted.

Example of the redirect:

```
GET https://example.com/oauth/callback?code=67f74f5a-a2cc-4ebd-88b4-22453fe07994&state=8e02c9c6a3484fadaaf841fb1df290e1
```

### 4. ***Application receives the authorization code***

You can use this code to exchange for an access token.

### 5. ***Application exchanges the authorization code with an access token from Alpaca***

After you have received the temporary `code`, you can exchange it for an access token. This can be done by making a POST call:

```
POST https://api.alpaca.markets/oauth/token
```

With following parameters:

| Parameter | Description |
| --------- | ----------- |
| `grant_type` | **Required** Must be set to `authorization_code` for an access token request. |
| `code` | **Required** The authorization `code` received in step 4 |
| `client_id` | **Required** The Client ID you received when you registered the application. |
| `client_secret` | **Required** The Client Secret you received when you registered the application. |
| `redirect_uri` | **Required** The redirect URI you used for the authorization code request. |

*Note: this request should take place behind-the-scenes from your backend server and shouldn't be visible to the end users for security purposes.*

Example request:

```
curl -X POST https://api.alpaca.markets/oauth/token \
  -d 'grant_type=authorization_code&code=67f74f5a-a2cc-4ebd-88b4-22453fe07994&client_id=fc9c55efa3924f369d6c1148e668bbe8&client_secret=5b8027074d8ab434882c0806833e76508861c366&redirect_uri=https://example.com/oauth/callback'
```

After a successful request, a valid access token will be returned in the response:

```
{
    "access_token": "79500537-5796-4230-9661-7f7108877c60",
    "token_type": "bearer",
    "scope": "account:write trading"
}
```

## <a name="APIcall"></a>API Call

Once you have integrated and have a valid access token you can start make calls to Alpaca Trading API v2 on behalf of the end-user.

Example requests:

```
curl https://api.alpaca.markets/v2/account /
  -H 'Authorization: Bearer 79500537-5796-4230-9661-7f7108877c60'
```

```
curl https://paper-api.alpaca.markets/v2/orders /
  -H 'Authorization: Bearer 79500537-5796-4230-9661-7f7108877c60'
```

## How to Use Polygon with OAuth

You can access Polygon service for the OAuth authorized user. As it is an
external service, you need to take an extra step to allow the access.

To call Polygon API for the users who have a live brokerage account,
the OAuth application needs to make a call to another OAuth endpoint, introspect.

```
curl https://api.alpaca.markets/oauth/token \
  -H 'Authorization: Bearer 79500537-5796-4230-9661-7f7108877c60'
```

The response will look like this.

```
{
    "active": true,
    "aud": "sample-application",
    "client_id": "test-client",
    "iat": 1562984846,
    "id": "TKM441PPNV3O3IX9RLCN",
    "owner_id": "15de0f1b-5a77-4398-97e4-a7b2d00975cb",
    "scope": "general",
    "token_type": "Bearer"
}
```

Then call the Polygon API with the `apiKey` parameter value be the `id`
value of this response.

```
curl 'https://api.polygon.io/v1/last/stocks/AAPL?apiKey=TKM441PPNV3O3IX9RLCN'
```

If the end-user account has the access to Polygon API, this will return the result
with success code.
