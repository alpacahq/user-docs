---
title: Python Tutorial
weight: 100
---

The purpose of this tutorial is to explain the general flow of building your first algorithm by going through an example. It is completely free to run this algorithm as-is or to modify it for your own needs. No guarantees about its performance are made, however.

## Algorithm Idea
The idea behind the sample algorithm is this:

- We trade stocks currently within the S&P 500. (Adjustments to account for changes in the S&P 500 "universe" of stocks are beyond the scope of this tutorial.)
- We build a portfolio with securities that we believe are oversold, trying to capture profit from mean reversion. We use exponential moving average (EMA) to measure how oversold we think something is.
- We rebuild our portfolio once a day in the morning.
- We maintain a reasonable size of positions (~5 stocks) at a time.

In short, this strategy hopes to capitlaize on some of the random-walk fluctuations in stock prices by placing orders to buy stocks when their prices stray too far below the EMA, hoping they will recover quickly. It shuffles the portfolio by ranking the most oversold stocks and keeping only the top-ranked stocks at any point. And due to the liquidity and generally low volatility of the S&P 500's consituents, we do not expect positions to have highly volatile returns. Therefore, the risk should be low and there should be some correlation with the S&P 500.

Now, we are going to go through the steps to write this algorithm. We have a finished version of the algorithm on GitHub [here](https://github.com/alpacahq/samplealgo01).

## Starting point
First thing first, you need to have a main entry point, which we call `main()`.

```py
def main():
    # empty!
    pass

if __name__ == '__main__':
    main()
```

## Setting up our Universe
There are more than 7,000 tradable symbols for US equities today, but
we stay away from smaller companies and ETFs here to avoid taking on
too much risk. We'll only consider S&P 500 stocks. The contents of the
S&P 500 can be obtained from different sources and referenced dynamically,
but in this tutorial, we'll just hard-code it in the source code.

```py
# SP500
Universe = ['MMM', 'ABT', 'ABBV', 'ACN', 'ATVI', 'AYI', 'ADBE', ..., 'ZTS']
```

We will store the security list in a file called `universe.py`.

## Setting up your API client
The Python SDK provides a `REST` class which represents the API client.
The class initializer takes the API key ID and secret key you got from
the dashboard as parameters, or you can set the environment variables
`APCA_API_KEY_ID` and `APCA_API_SECRET_KEY` (which are only used if
parameters are not specified.) Since we'll be using the paper API, we'll
also need to specify the base URL parameter (or the `APCA_API_BASE_URL`
environment variable) as `https://paper-api.alpaca.markets`.


Here we create an `api` object from this class.

```py
import alpaca_trade_api as tradeapi

api = tradeapi.REST(
    key_id='REPLACEME',
    secret_key='REPLACEME',
    base_url='https://paper-api.alpaca.markets'
)
```

With this `api` object, you can call each REST API endpoint
as its corresponding method, such as `get_account()`.

## Infinite loop
Now, let's start coding the logic! Most algorithms watch
market movements and take actions based on specific market
conditions. For this to occur, our algorithm needs to run infinitely.
Since we want to check the daily movement of our universe, we will
check the time (Eastern US, for consistency with market time) and do
something in the morning.

```py
def main():
    done = None
    logging.info('start running')
    while True:
        # clock API returns the server time including
        # the boolean flag for market open
        clock = api.get_clock()
        now = clock.timestamp
        if clock.is_open and done != now.strftime('%Y-%m-%d'):
            # ** do our stuff here! **

            do_stuff()

            # flag it as done so it doesn't work again for the day
            # TODO: this isn't tolerant to process restarts, so this
            # flag should probably be saved on disk
            done = now.strftime('%Y-%m-%d')
            logger.info(f'Done for {done}')

        time.sleep(1)
```

As you can see in the code, this algo checks if the market is open by
using the `get_clock()` API that tells the time at the server side.
This flag becomes `True` when it is market hours. If so, we want it to
do our trading. If you want to check whether the market is open on other
dates as well, Alpaca provides a way to do this with with `get_calendar()`.

The `done` flag is set to the date string so we make sure it does work only
once a day. When developing your own trading idea, you may want to kick off
your logic every minute, every hour, every Monday, or whatever you want.

## Main logic
This is the fun part since this is the core of our algorithm. Remember,
our algorithm calculates EMA to find the most oversold stocks in the S&P 500
universe of stocks.

### Get the price data
First, we want to get price data to calculate the EMA. We'll use the [/bars]({{< ref path="/web-api/market-data/bars/_index.md" >}})
endpoint of the Alpaca data API to do this. Alpaca users with brokerage
accounts also have access to a premium data source called Polygon. Since
we're starting out with paper trading, which can be done without a brokerage
account, we'll only be talking about Alpaca's data API in this tutorial. For
more information about the data sources available to you, reference [this page]({{< ref path="/web-api/market-data/_index.md" >}}).

Price data can be retrieved from the `/bars` endpoint by calling `get_barset`
on the REST object. For example:

```py
api.get_barset('AAPL', 'day', limit=1)
```

This will give you Apple's OHLCV information for the current trading day or,
if the market is closed, the last day the market was open. We'll define a
function called `prices()` which takes a parameter, symbols, to indicate which
price data to get.

```py
def prices(symbols):
    now = pd.Timestamp.now(tz=NY)
    end_dt = now
    if now.time() >= pd.Timestamp('09:30', tz=NY).time():
        end_dt = now - \
            pd.Timedelta(now.strftime('%H:%M:%S')) - pd.Timedelta('1 minute')
    return _get_prices(symbols, end_dt)
```

There are some checks to adjust what to specify for `end_dt` parameter since
we want to make sure this function always returns the prices up to yesterday,
even if you call it during market hours. If you call this like `prices(['AAPL'])`,
you will get a DataFrame object containing Apple's price data.

`_get_prices()` will look like this:

```py
def _get_prices(symbols, end_dt, max_workers=5):
    '''Get the map of DataFrame price data from Alpaca's data API.'''

    start_dt = end_dt - pd.Timedelta('50 days')
    start = start_dt.strftime('%Y-%-m-%-d')
    end = end_dt.strftime('%Y-%-m-%-d')

    def get_barset(symbols):
        return api.get_barset(
            symbols,
            'day',
            limit = 50,
            start=start,
            end=end
        )

    # The maximum number of symbols we can request at once is 200.
    barset = None
    idx = 0
    while idx <= len(symbols) - 1:
        if barset is None:
            barset = get_barset(symbols[idx:idx+200])
        else:
            barset.update(get_barset(symbols[idx:idx+200]))
        idx += 200

    return barset.df
```

### Rank stocks by (price - EMA) difference
It is hard to define which stocks are "the most oversold" among a number of
stocks, but let's assume the difference between the price and EMA indicates
something of value here. Short-term EMA usually converges close to the price,
but if it diverges a lot instead, that means the price changed siginificantly
in a short period of time. We'll need to normalize the value of the difference
so we can compare the significance between stocks in a fair manner.

```py
def calc_scores(price_df, dayindex=-1):
    '''Calculate scores based on the indicator and
    return the sorted result.
    '''
    diffs = {}
    param = 10
    for symbol in price_df.columns.levels[0]:
        df = price_df[symbol]
        if len(df.close.values) <= param:
            continue
        ema = df.close.ewm(span=param).mean()[dayindex]
        last = df.close.values[dayindex]
        diff = (last - ema) / last
        diffs[symbol] = diff

    return sorted(diffs.items(), key=lambda x: x[1])
```

We use pandas's `ewm()` method to calculate the EMA on a DataFrame here,
but if you want to use a different technical indicator to find oversold stocks,
you could use the `ta-lib` package, which supports a wider variety of inicators.
Please note that `diff` is the difference between the last price and 10-day EMA
as a percentage of last price. This value can be negative or positive,
with a negative `diff` indicating the price dropped recently.

### Build orders
Now that we have the ranked list of stocks in hand, it's time to decide
what to buy and what to sell. We want to keep the top 5 oversold stocks
in our portfolio, and it's easy to build 5 buy orders if you start from
zero, but we need to do a bit of work here to check our current holdings.

```py
def get_orders(api, price_df, position_size=100, max_positions=5):
    '''Calculate the scores within the universe to build the optimal
    portfolio as of today, and extract orders to transition from our
    current portfolio to the desired state.
    '''
    # rank the stocks based on the indicators.
    ranked = calc_scores(price_df)
    to_buy = set()
    to_sell = set()
    account = api.get_account()
    # take the top one twentieth out of ranking,
    # excluding stocks too expensive to buy a share
    for symbol, _ in ranked[:len(ranked) // 20]:
        price = float(price_df[symbol].close.values[-1])
        if price > float(account.cash):
            continue
        to_buy.add(symbol)

    # now get the current positions and see what to buy,
    # what to sell to transition to today's desired portfolio.
    positions = api.list_positions()
    logger.info(positions)
    holdings = {p.symbol: p for p in positions}
    holding_symbol = set(holdings.keys())
    to_sell = holding_symbol - to_buy
    to_buy = to_buy - holding_symbol
    orders = []

    # if a stock is in the portfolio, and not in the desired
    # portfolio, sell it
    for symbol in to_sell:
        shares = holdings[symbol].qty
        orders.append({
            'symbol': symbol,
            'qty': shares,
            'side': 'sell',
        })
        logger.info(f'order(sell): {symbol} for {shares}')

    # likewise, if the portfoio is missing stocks from the
    # desired portfolio, buy them. We sent a limit for the total
    # position size so that we don't end up holding too many positions.
    max_to_buy = max_positions - (len(positions) - len(to_sell))
    for symbol in to_buy:
        if max_to_buy <= 0:
            break
        shares = position_size // float(price_df[symbol].close.values[-1])
        if shares == 0.0:
            continue
        orders.append({
            'symbol': symbol,
            'qty': shares,
            'side': 'buy',
        })
        logger.info(f'order(buy): {symbol} for {shares}')
        max_to_buy -= 1
    return orders
```

You can use `calc_scores()` to get the ranked list and `get_account()` and
`list_positions()` to know your current available cash and held positions.
We filter out some stocks whose prices exceed the size of our cash balance,
and we limit the number of positions, even if the ranked list has more stocks
that are oversold.

Note the function has default parameters `position_size` and `max_positions`
that control how many dollars you want to spend at most for each position and
how many positions at most you want to keep in your portfolio. You can change
these to your own liking.

Finally, we return the orders that will transition from our current portfolio
to the desired portfolio.

### Place the orders!
We're finally ready to execute the orders. In this algorithm, we separate
the logic of calculating necessary orders and actual order submissions
so we can easily test the code, but it is possible to mix that logic too.

The main concern you have here is that the buy orders may get rejected
if your cash balance is not enough, so you need to wait for the sell orders
to fill before you submit any buys. This algorithm does not care much
about the precise entry price, for simplicity. It places market orders
so that we don't need to code around limit orders which might not be filled.

```py
def trade(orders, wait=30):
    '''This is where we actually submit the orders and wait for them to fill.
    Waiting is an important step since the orders aren't filled automatically,
    which means if your buys happen to come before your sells have filled,
    the buy orders will be bounced. In order to make the transition smooth,
    we sell first and wait for all the sell orders to fill before submitting
    our buy orders.
    '''

    # process the sell orders first
    sells = [o for o in orders if o['side'] == 'sell']
    for order in sells:
        try:
            logger.info(f'submit(sell): {order}')
            api.submit_order(
                symbol=order['symbol'],
                qty=order['qty'],
                side='sell',
                type='market',
                time_in_force='day',
            )
        except Exception as e:
            logger.error(e)
    count = wait
    while count > 0:
        pending = api.list_orders()
        if len(pending) == 0:
            logger.info(f'all sell orders done')
            break
        logger.info(f'{len(pending)} sell orders pending...')
        time.sleep(1)
        count -= 1

    # process the buy orders next
    buys = [o for o in orders if o['side'] == 'buy']
    for order in buys:
        try:
            logger.info(f'submit(buy): {order}')
            api.submit_order(
                symbol=order['symbol'],
                qty=order['qty'],
                side='buy',
                type='market',
                time_in_force='day',
            )
        except Exception as e:
            logger.error(e)
    count = wait
    while count > 0:
        pending = api.list_orders()
        if len(pending) == 0:
            logger.info(f'all buy orders done')
            break
        logger.info(f'{len(pending)} buy orders pending...')
        time.sleep(1)
        count -= 1
```

It's as simple as that. The order submission (`submit_order()`)
is enclosed by a try-except block in case we get some error related
to the transactions. If an error is encountered, we only log it, as
we'd still rather transition the portfolio as much as possible rather
than bring everything to a halt. This does mean that you should check
the error log occasionally while it runs to make sure something hasn't
gone wrong, as it will continue on if left to its own devices.

## Assemble them
Now that we have all the pieces ready, we just need to execute the main
logic once a day.  Just put the `get_orders()` and `trade()` in the middle
of the main loop.

```py
def main():
    done = None
    logging.info('start running')
    while True:
        # clock API returns the server time including
        # the boolean flag for market open
        clock = api.get_clock()
        now = clock.timestamp
        if clock.is_open and done != now.strftime('%Y-%m-%d'):

            price_df = prices(Universe)
            orders = get_orders(api, price_df)
            trade(orders)

            # flag it as done so it doesn't work again for the day
            # TODO: this isn't tolerant to process restarts, so this
            # flag should probably be saved on disk
            done = now.strftime('%Y-%m-%d')
            logger.info(f'done for {done}')

        time.sleep(1)
```

With the default parameters you saw in the example, it trades with
- S&P 500 stocks
- 5 positions at most
- At most 100 dollars for each position

You can adjust these values based on your needs, but this should be a good
basis. Also, note that this algo will not be given a day trading flag, as long
as it isn't restarted in the middle of a trading day. This means that you won't
have to worry about the $25k account balance requirement for pattern day traders.

It is very hard to try this type of shuffling algorithm without a commission-free
trading platform with this size of cash, since a few dollars of commission will
kill the cash balance and profit margins pretty quickly.

## Set up the enviornment and run the script
Let's get ready to run our script. For your convenience, we've got a repository
set up for this script on GitHub. You should start off by cloning it from [here](https://github.com/alpacahq/samplealgo01).
This repository is set up using `pipenv`, something which sits on top of `pip`,
`virtualenv`, and `pyenv`, and makes them all easier to use. If you haven't
installed `pipenv`, we'd recommend trying it. Once you have `pipenv`, it's as
simple as running these commands:

```
$ pipenv install
$ pipenv shell
```

That will open a shell in the virtual environment with all dependencies installed.

To get the algorithm started, you can do this command from within the pipenv shell:
```
$ python main.py
```
If the market is open, you should see it think for a moment and then place some orders.

## Deployment
This algo has to run live to trade. The question is, where? You
may not want to keep your computer up and running all the time,
and you probably don't want to worry about monitoring whether or
not your computer's online.

You can borrow a machine from AWS or one of the other cloud providiers, which
costs a few dollars a month. The smallest instance should be fine for this
algo, since it doesn't do much CPU work, but you have to set up the environment,
install dependencies, etc.

Alternatively, you can break the infinite loop so it only runs the main logic
and then exits, so that you can run this in AWS Lambda or on some sort of
cron-job service. This is possible since this particular algo does not need to
watch the positions all the time - only once a day. But if you want to add much
more logic to this, it will be a bit harder to keep it running like a cron job.

### Heroku free
Heroku, a cloud platform designed with app hosting in mind,
offers a free tier that can run this simple program for you.
What is important to understand is that you need to set it up
as a "worker" process, since this is a long-running process.

First, set up your Heroku account if you haven't. Then, create an
App in the account.

![Create App](https://cdn.pbrd.co/images/HrJxZNGQ.png)

The next screen asks for the app name, which must be unique. You can
name it however you want.

---
![Name App](https://cdn.pbrd.co/images/HrJzskv.png)

Follow Heroku's instructions to set up the command line tool so you can push your code
to Heroku. You can either set up a GitHub account that links to this Heroku App,
or you can install the `heroku` command line tool and add `heroku` remote to
your local GitHub account.

Make sure you configure your free dyno for the worker.

---
The last thing to do is to configure the environment variables to give
your Alpaca API key info. Go to the `Settings` tab and add the
environment variables.

![Configure Env](https://cdn.pbrd.co/images/HrJAdmb.png)

---
If you run it correctly, you can see the console log from here.

![View Logs](https://cdn.pbrd.co/images/HrJzdqo.png)

If everything works fine, you will see some log output in the
log output, similar to this.

```
2018-06-26T13:30:58.502315+00:00 app[worker.1]: INFO:samplealgo.algo:1 buy orders pending...
2018-06-26T13:30:59.574241+00:00 app[worker.1]: DEBUG:urllib3.connectionpool:https:/iapi.alpaca.markets:443 "GET /v1/orders HTTP/1.1" 200 557
2018-06-26T13:30:59.574927+00:00 app[worker.1]: INFO:samplealgo.algo:1 buy orders pending...
2018-06-26T13:31:00.662481+00:00 app[worker.1]: DEBUG:urllib3.connectionpool:https://api.alpaca.markets:443 "GET /v1/orders HTTP/1.1" 200 557
2018-06-26T13:31:00.663137+00:00 app[worker.1]: INFO:samplealgo.algo:1 buy orders pending...
2018-06-26T13:31:01.740745+00:00 app[worker.1]: DEBUG:urllib3.connectionpool:https://api.alpaca.markets:443 "GET /v1/orders HTTP/1.1" 200 557
2018-06-26T13:31:01.741839+00:00 app[worker.1]: INFO:samplealgo.algo:1 buy orders pending...
2018-06-26T13:31:02.837375+00:00 app[worker.1]: DEBUG:urllib3.connectionpool:https://api.alpaca.markets:443 "GET /v1/orders HTTP/1.1" 200 557
2018-06-26T13:31:02.838408+00:00 app[worker.1]: INFO:samplealgo.algo:1 buy orders pending...
2018-06-26T13:31:03.840173+00:00 app[worker.1]: INFO:samplealgo.algo:done for 2018-06-26
```

Now your script is running! Feel free to play around with its default settings
and try hosting your own algorithm in a similar way when you're ready.