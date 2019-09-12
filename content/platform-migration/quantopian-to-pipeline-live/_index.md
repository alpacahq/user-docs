---
title: Pipeline Migration
weight: 60
---
# Migrate your Pipeline from Quantopian
pipeline-live helps you run your algorithm outside of the Quantopian.
Although this project is an independent effort to provide the Pipeline
API using public/private data, this document is to describe the common
practices around how to migrate your pipeline code from the Quantopian
environment.

Along with these practices, you can migrate your Algorithm API from Quantopian
using [pylivetrader](https://github.com/alpacahq/pylivetrader), and
pylivetrader can run the pipeline object from this package.

## USEquityPricing
The most important class to think about first is the USEquityPricing class
and it is well covered by the
`pipeline_live.data.alpaca.pricing.USEquityPricing` class.
Depending on the requested window length from its upstream pipeline, it
fetches different size of the data range (e.g. 3m, 1y). Again, the volume of
this data is market-wide size, so it's safe to use this with factors such
as AverageDollarVolume.

## Factors
In order to use many of the builtin factors with this price data loader,
you need to use `pipeline_live.data.alpaca.factors` package which has
all the builtin factor classes ported from zipline.

For example, if you have these lines,

```py
from quantopian.pipeline.factors import (
    AverageDollarVolume, SimpleMovingAverage,
)
from quantopian.pipeline.data.builtin import USEquityPricing
```

you can rewrite it to something like this.

```py
from pipeline_live.data.alpaca.factors import (
    AverageDollarVolume, SimpleMovingAverage,
)
from pipeline_live.data.alpaca.pricing import USEquityPricing
```

Of course, the builtin factor classes in the original zipline are mostly
pure functions and take inputs explicitly, so if you give the correct
ones, they also work with this `USEquityPricing`.

```py
from zipline.pipeline.factors import AverageDollarVolume
from pipeline_live.data.alpaca.pricing import USEquityPricing

dollar_volume = AverageDollarVolume(
    inputs=[USEquityPricing.close, USEquityPricing.volume],
    window_length=20,
)
```

The only difference in the factor classes in `pipeline_live.data.alpaca.factors`
is that some of the classes have Alpaca's USEquityPricing as the default
inputs, so you don't need to explicitly specify it.

## Fundamentals
The Quantopian platform allows you to retrieve various proprietary data
sources through pipeline, including Morningstar fundamentals. Previously,
IEX was used by pipeline-live to supply equivalents to these, but recent
changes to the IEX API have made this less possible for most use cases.
The alternative at the moment is the Polygon dataset, which is available
to users with funded Alpaca brokerage accounts and direct subscribers of
Polygon's data feed. If you want to get started with Polygon fundamentals,
please see [the repository's readme file](https://github.com/alpacahq/pipeline-live#polygon-data-source-api)
for more info on what Polygon information is currently available through pipeline-live.

## Primary Share
Many algorithms developed in the Quantopian platform uses the `IsPrimaryShare`
function to perform base filter. While this value is unique to Morningstar,
a similar filter has been created in pipeline-live for users with Polygon data.
Please look at the `pipeline_live.data.polygon.filters.IsPrimaryShareEmulation` class for
the replacement.

## ADR
If you have access to Polygon, you can check out the `PolygonCompany.country`
field to filter out non-US companies.