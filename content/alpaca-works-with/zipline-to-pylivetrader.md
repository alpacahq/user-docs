---
title: Zipline Migration
weight: 50
---
# Migrate your Algorithm from Quantopian

This document describes how you can transform your algorithm source built in
Quantopian to pylivetrader. pylivetrader API is compatible with zipline API,
but there are some changes that require mechanical and manual process to
move from Quantopian/zipline.

## Copy and Paste
First thing first, you need to create a python script. Go to your Quantopian
algorithm page, copy the entire source code, paste and save it as a
python script. Here we name it `algo.py` as an example.

## Python 2 and 3 (if applicable)
Quantopian's environment, as of writing, supports both Python 2 and 3.
pylivetrader currently supports only Python 3, so if your algorithm is
currently Python 2, you need to convert your Python 2 script to Python 3.
(If your script is already Python 3, you may skip to the next section)

While those two versions are not compatible, there are not so many differences.
Fortunately, the Python community has been putting so much effort to make it 
easy to migrate from version 2 to 3.

In short, the `2to3` command will help here.  You should have this program
if you have installed Python 3. For more details on this script, you can
read the [official document](https://docs.python.org/3.6/library/2to3.html).

```sh
$ 2to3 -w algo.py
```

This updates your Python script so it works with Python 3. Remember, you
will have the original file as `algo.py.bak`.

## Fix imports
pylivetrader is API-compatible with Quantopian/zipline, but you still need to
update your import statements so you import pylivetrader modules, instead of
Quantopian/zipline ones. Any imported modules starting from `quantopian`
will not work outside of the Quantopian environment. There are a couple of cases
that need different solutions here.

### Algorithm API
This is the API under the `quantopian.algorithms` or `zipline.algorithm`
package. The package path should be replaced with `pylivetrader.algorithm`.
All the Algorithm API is ported to pylivetrader, but some of the API (such as
futures trading specific API) will throw `NotSupported` exception.

### Pipeline API
This is the API under the `quantopian.pipeline` or `zipline.pipeline` package.
pylivetrader does not provide direct replacement for pipeline, but you can
use [pipeline-live](https://github.com/alpacahq/pipeline-live). You need to
install this package separately as pylivetrader does not automatically
include it.  For more information,
please read [pipeline-live's migration document](https://github.com/alpacahq/pipeline-live/blob/master/migration.md).

Under `attach_pipeline` and `pipeline_output`, pylivetrader uses
the pipeline-live package if it is installed, otherwise throws `RuntimeError`.

### Optimize API
This is the API under the `quantopian.optimize` or `zipline.optimize`
package. Optimize API
is not currently supported by pylivetrader.

### Anything Else that Starts with `quantopian.`
For those other APIs, you will have to change the code so it works as you
intended.  The good news is that there is no limitation or restriction in this
pylivetrader, so you will probably be able to do what you can do.

### Zipline API
The other zipline API is ported to pylivetrader as much as possible, and
should cover most of the use cases. Some of the features such as commission
model is no-operation, since it is not applicable in the live trading
environment, but they still exist for compatibility.

### Implicit imports
The Quantopian environment adds certain functions to the namespace
when it loads the algorithm source code. In pylivetrader, your algorithm
file has to explicitly import those by yourself.  These "auto" functions
are under `pylivetrader.api` package.  You can add this line to your
algorithm source code.

```py
from pylivetrader.api import *
```

The implicitly imported APIs are the following.

- order
- order_value
- order_percent
- order_target
- order_target_value
- order_target_percent
- cancel_order
- get_open_orders
- get_order
- continuous_future (unsupported in pylivetrader)
- fetch_csv (unsupported in pylivetrader)
- get_datetime
- get_environment (unsupported in pylivetrader)
- log (need extra step)
- record
- schedule_function
- set_symbol_lookup_date (no-op in pylivetrader)
- sid
- symbol
- symbols
- set_long_only
- set_max_order_count
- set_max_order_size
- set_max_position_size
- set_max_leverage:w

## The `log` Object
Although this is not an imported module, Quantopian provides
the `log` object in the algorithm. In pylivetrader, you can create
one by the following snippet.

```py
import logbook
log = logbook.Logger('algo')
```

## Deal with Restart
In the live trading environment, you possible need to restart the
algorithm program during the day. For one, your program suddenly dies
for some reason, and for another, you may want to intentionally suspend
and resume it.

Keep in mind that pylivetrader persists the information you stored
in `context` property on disk, every time it runs scheduled functions.
However, `initialize()` and `before_trading_start()` will be also triggered
in the event of restart. It is often observed in Quantopian that an
algorithm assumes `initialize()` to be called only once for the life, and
`before_trading_start()` to be called only once a day, which may not be
true.

Typically to deal with it, you may want to modify `before_trading_start()`
so that it checks the last time it ran.

```py
def before_trading_start(context, data):
    # context.age is to remember some age of positions
    if not hasattr(context, 'age') or not context.age:
        context.age = {}

    today = get_datetime().floor('1D')
    last_date = getattr(context, 'last_date', None)
    # check if it ran today
    if today != last_date:

       ...do actual work...

       context.last_date = today
```

This ensures the "actual work" will be executed only once in a day
with restart, and it is a good idea to not change `context` object
in the `initialize()` function but here in `before_trading_start()`
function.


## Checking and Testing
Once you convert your algorithm code, you may want to check if there is
no easy mistake at the literal level. This is optional, but we recommend
using [`autopep8`](https://pypi.org/project/autopep8/) and
[`flake8`](https://pypi.org/project/flake8/) which you can install via
`pip` and tell there are some syntax errors or uninitialized variables etc.

You are almost there, but remember, good software is always tested before
deployed, and there is no difference in trading algorithm. At the moment,
your option is to write a unit test using `unittest` package and `pytest`
by yourself, but pytrader will soon have a good support for it, too.

It is also recommended to take advantage of paper trading account if your
broker supports it. You should expect some issues such as network problem,
API to return errors, orders not being filled for long. Paper trading will
help you spot these issues by running your algorithm in live.

## Future Plans
pylivetrader is designed to run Quantopian/zipline algorithm with minimum
effort. The procedure explained here will be also automated as much as
possible by pylivetrader in the future. Also, as discussed above, the
unit test helper class is also under development currently.