## `prometheus` package

Package prometheus provides tools for working with
[Prometheus-formatted metrics](https://prometheus.io/docs/instrumenting/exposition_formats/).

Import the `prometheus` package:

```flux
import "prometheus"
```

### Functions

### `histogramQuantile()`

histogramQuantile calculates a quantile on a set of Prometheus histogram values.

#### Function type signature

```flux
(<-tables: stream[{B with le: D, _field: C}], quantile: float, ?metricVersion: A, ?onNonmonotonic: string) => stream[E] where A: Equatable, C: Equatable, E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| quantile | quantile: Quantile to compute. Must be a float value between 0.0 and 1.0. | Yes |
| metricVersion | metricVersion: [Prometheus metric parsing format](https://docs.influxdata.com/influxdb/latest/reference/prometheus-metrics/)
  used to parse queried Prometheus data.
  Available versions are `1` and `2`.
  Default is `2`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
| onNonmonotonic | onNonmonotonic: Describes behavior when counts are not monotonically increasing
  when sorted by upper bound. Default is `error`. | No |
### `scrape()`

scrape scrapes Prometheus metrics from an HTTP-accessible endpoint and returns
them as a stream of tables.

#### Function type signature

```flux
(url: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL to scrape Prometheus metrics from. | Yes |
