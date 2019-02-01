---
title: Who is Alpaca?
weight: 10
---

# Who is Alpaca?

## Who We Are

Alpaca Securities is a FINRA and SEC registered broker-dealer that launched in early 2018 to serve the needs 
of the growing community of software developers and technology-minded individuals. We offer a robust, user-friendly
API which can be used to automate your trading and investing of U.S. equities.

Alpaca's team consists of developers, traders, and brokerage business specialists, who collectively have decades of
financial services and technology industry experience at organizations such as FINRA, IBM, EMC, Greenplum, 
Wealthfront, Lehman Brothers, UBS. Alpaca is led by Yoshi Yokokawa(CEO), Hitoshi Harada(CTO), John Torola(COO/CCO), 
and Luke Lonergan(Chief Architect). Our investors include a group of well-capitalized investors including Global 
Brain, the biggest Japanese independent venture capital firm, as well as highly experienced industry angel investors 
such as Josh Levine (former CTO/COO of ETRADE) and Eric Di Benedetto (a fintech angel investor with more than 30 
IPO and M&A exits). Additionally, members of Berkeley Angel Network and two biggest Japanese banks (MUFG and SMBC) 
are investors in Alpaca.

## Who Are Our Customers?

Alpaca Securities provides brokerage services to all types of users with a variety of backgrounds. Many of our 
customers come from one of the largest quantitative research platforms, Quantopian. Others are experienced automated 
or manual traders coming from Interactive Brokers, TD Ameritrade, or Robinhood. Some are crypto traders expanding 
into equities. Some have no background in trading at all and are just learning to code. We even receive inquiries from 
boutique quant funds (institutional accounts are on our roadmap!). We welcome them all and seek to provide
a platform for anyone, novice or experienced, to automate their investing and trading.

## Business Model

Securities brokers traditionally generate revenue in multiple ways including charging trading commissions, marking 
up margin lending rates, keeping interest on cash deposits, receiving payment for order flow by routing orders to 
market makers, and marking up the data feed subscription. 

At Alpaca, we believe in the power of automation in trading and investing, which is notably different from that of 
traditional securities brokers. With this belief, we focused on removing obstacles that would prevent such automation 
from happening. 

Naturally, automation tends to generate a large number of smaller transactions. At a traditional broker, this might
result in an unbearable commission expense that outweighs all other costs. To address that, we made it a point to offer 
our users commission-free automated trading. This breathes new life into many types of strategies, including those that 
trade frequently or use small order sizes. It also supports developers and technology-minded individuals in their
journey by allowing them to start small and gradually scale up. Ultimately, our customers' success leads to our success. 

### How will Alpaca make money without charging commissions? 
Alpaca will launch a premium subscription plan, where premium users will have access to various perks, 
such as a higher quality real-time data feed and additional computing resources. 
The pricing and exact features have not yet been determined. Please stay tuned for updates.

Additionally, Alpaca may generate revenue in many of the same ways as traditional online brokerages. These include:

* Interest on cash deposits
* Payment for order flow ("PFOF") - Alpaca receives remuneration for routing your orders to
market makers and exchanges. PFOF helps us offset the expense that occurs when clearing and executing our
customers' trades. You can read more about PFOF in our Medium post 
[here](https://medium.com/automation-generation/commission-free-trading-is-it-helping-or-hurting-you-dc5fdc22ca6a).
For SEC Rule 606, please visit the SEC site. *It is important to note that our customers are not charged.
* Margin financing - Alpaca may charge interest for margin loans.
* Stock loan - Alpaca may charge stock loan fees for users who want to borrow stock to short sell.

## Regulated Broker-Dealer

Alpaca Securities LLC (Alpaca Securities) obtained a license from FINRA to provide securities brokerage services 
at the beginning of 2018. In order for Alpaca Securities to obtain such license and membership, we had to go through 
rigorous paperwork and system development to become fully compliant with the SEC and FINRA. Alpaca Securities' 
brokerage service and operation is led by John Torola, who is a FINRA veteran for more than 20+ years and most 
recently led the compliance team at one of the biggest robo advisory firms, WealthFront. John oversees the brokerage 
operations team consisting of licensed professionals who came from firms including UBS, Morgan Stanley, Wells Fargo.

You can check the background of Alpaca on FINRA’s [BrokerCheck](https://brokercheck.finra.org/firm/summary/288202).

## SIPC Membership
### Safeguarding your assets
Alpaca is a member of SIPC (Securities Investor Protection Corporation). SIPC protects against the loss of 
cash and securities held by a customer up to $500,000, which includes a $250,000 limit for cash. SIPC does
not protect you for any decline in the value of your securities.

For details, please visit [www.sipc.org](www.sipc.org).

## Systems and Technology
We initially started Alpaca in 2015 as a pure technology company building a database solution for unstructured data, 
initially visual data and ultimately time-series data. One of our achievements, which is being improved on an ongoing 
basis, is [MarketStore](https://github.com/alpacahq/marketstore), a dedicated open source database written in Go designed 
specifically to handle financial markets time series data. This led us to provide our technology to financial 
institutions including Bank of Tokyo Mitsubishi, which handles more than 50% of the USD/JPY trading flow. We also 
managed a deep neural net trading algorithm service, AlpacaAlgo, that had more than 25,000 deep neural net trading 
algorithms simultaneously constantly generating trade signals to the users. This required processing a large volume 
of time series data efficiently. Hitoshi and Luke lead Alpaca's technology team and are responsible for the data 
intensive system architecture. Both of them come from Greenplum (Luke is a co-founder/CTO of Greenplum), a parallel 
distributed database company, that was used globally by major financial institutions and their quant desks.  