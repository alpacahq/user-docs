---
title: Who is Alpaca?
weight: 10
---

# Who is Alpaca?

## Who We Are

Alpaca started in 2015 as a pure technology company building a database solution for unstructured data, initially
visual data and ultimately time-series data. We then launched an API-first commission-free stock trading platform
in early 2018 to serve the needs of the growing community of software developers and technology-minded individuals.
We offer a robust, user-friendly API which can be used to automate your trading and investing of U.S. equities.

Alpaca's team consists of developers, traders, and brokerage business specialists, who collectively have decades of
financial services and technology industry experience at organizations such as FINRA, Wealthfront, Google, IBM, EMC, Greenplum, 
Lehman Brothers, and UBS. Alpaca is led by [Yoshi Yokokawa](https://www.linkedin.com/in/yoshiyokokawa/) (CEO),
[Hitoshi Harada](https://www.linkedin.com/in/hitoshi-harada-02b01425/) (CTO), [John Torola](https://www.linkedin.com/in/johnttorola/) (COO/CCO),
and [Luke Lonergan](https://www.linkedin.com/in/lukelonergan/) (Chief Architect). Our investors include a group of well-capitalized 
investors including Spark Capital, Y Combinator, Social Leverage, Portag3, as well as highly experienced industry angel investors 
such as Joshua S. Levine (former CTO/COO of ETRADE), Eric Di Benedetto (a fintech angel investor with more than 30 IPO and M&A exits), 
and several YC alumni angels. Additionally, members of the Berkeley Angel Network and two of the largest Japanese banks (MUFG and SMBC) are 
investors in Alpaca.

## Who Is Alpaca For?

Alpaca Securities provides brokerage services to technology-minded users with a variety of experiences
including general usage of APIs or automation of investing and trading. We have users come from two of
the largest quantitative research platforms, Quantopian and QuantConnect. Others are experienced
automated or manual traders coming from Interactive Brokers, TD Ameritrade, or Robinhood. Some are
crypto traders expanding into equities. Some have no background in trading at all and are just learning
to code. We even receive inquiries from boutique quant funds (corporate accounts are on our roadmap!).
We welcome them all and seek to provide a platform for anyone, novice or experienced, to automate their
investing and trading.

Alpaca users trade a variety of strategies ranging from passive to active, long-term to short-term intraday.
As a point of reference, we have some users who execute hundreds of trades worth millions of dollars a day.

With an Alpaca brokerage account, users have access to:

* Commission-free trading
* An official API that is constantly being improved and updated
* Easy to understand API documentation
* Free, real-time Consolidated Market Data
* Paper trading for testing your strategies
* Fully working, example algorithms that can be deployed immediately
* Multiple support channels

Please see the bottom of the [Alpaca homepage](https://alpaca.markets/) for disclosures.

## How Alpaca Works

![Alpaca Trade Flow](/images/trade-flow.png)

## Why Commission-Free?

At Alpaca, we believe in the power of automation in trading and investing, which is notably different from that of
traditional securities brokers. With this belief, we focused on removing obstacles that would prevent such automation
from happening. Naturally, automation tends to generate a large number of smaller transactions. At a traditional broker,
this might result in an unbearable commission expense that outweighs all other costs. To address that, we
made it a point to offer our users commission-free automated trading. This breathes new life into many
types of strategies, including those that trade frequently or use small order sizes. It also supports
developers and technology-minded individuals in their journey by allowing them to start small and
gradually scale up. Ultimately, our customers' success leads to our success.

## How We Make Money?

Securities brokers traditionally generate revenue in multiple ways including charging trading commissions, marking
up margin lending rates and stock loan, keeping interest on cash deposits, receiving payment for order flow by routing orders to
market makers, and marking up the data feed subscription.

Although we do not charge commissions, Alpaca may generate revenue in some of the same ways as traditional online
brokerages. These include:

* Interest on cash deposits
* Payment for order flow ("PFOF") - Alpaca receives remuneration for routing your orders to
market makers and exchanges. PFOF helps us offset the expense that occurs when clearing and executing our
customers' trades. You can read more about PFOF in our Medium post
[here](https://medium.com/automation-generation/commission-free-trading-is-it-helping-or-hurting-you-dc5fdc22ca6a).
For Alpaca Securities SEC Rule 606 disclosures, please click [here](https://alpaca.markets/disclosures). *It is important to note that our customers are not charged.
* Margin financing - Alpaca may charge interest for margin loans.
* Stock loan - Alpaca may charge stock loan fees for users who want to borrow stock to short sell.

## Regulated Financial Institution

At the beginning of 2018 Alpaca Securities LLC (Alpaca Securities) became a member of FINRA and registered
with the SEC to provide securities brokerage services. In order for Alpaca Securities to obtain such
membership and registration, we had to go through rigorous paperwork and procedure development. Alpaca
Securities’ brokerage service and operations are led by John Torola, who is a FINRA veteran with more than
25+ years of financial services experience. He most recently led the compliance team at one of the biggest
robo advisory firms, WealthFront. John oversees a brokerage operations team consisting of licensed
professionals who came from such firms as UBS, Morgan Stanley, and Wells Fargo.

You can check the background of Alpaca on FINRA’s [BrokerCheck](https://brokercheck.finra.org/firm/summary/288202).

## Safety of Your Money
**SIPC Membership**

Member of SIPC, which protects securities customers of its members up to $500,000 (including $250,000 for
claims for cash). Explanatory brochure available upon request or at [www.sipc.org](https://www.sipc.org).

## Systems and Technology
We initially started Alpaca in 2015 as a pure technology company building a database solution for unstructured data,
initially visual data and ultimately time-series data. One of our achievements, which is being improved on an ongoing
basis, is [MarketStore](https://github.com/alpacahq/marketstore), a dedicated open source database written in Go designed
specifically to handle financial markets time series data. This led us to provide our technology to financial
institutions including Bank of Tokyo Mitsubishi, Japan's largest bank. We also managed a deep neural net trading
algorithm service, AlpacaAlgo, that had more than 25,000 deep neural net trading algorithms simultaneously constantly
generating trade signals to the users. This required processing a large volume of time series data efficiently.
Hitoshi and Luke lead Alpaca's technology team and are responsible for the data intensive system architecture.
Both of them come from Greenplum (Luke is a co-founder/CTO of Greenplum), a parallel distributed database company,
that is used globally by major financial institutions and their quant desks.