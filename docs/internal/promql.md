## `promql` package

Package promql provides an internal API for implementing PromQL via Flux.

Import the `promql` package:

```flux
import "promql"
```

### Functions

### `changes()`

changes implements functionality equivalent to
[PromQL's `changes()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#changes).

#### Function type signature

```flux
(<-tables: stream[{A with _value: float}]) => stream[{B with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `emptyTable()`

emptyTable returns an empty table, which is used as a helper function to implement
PromQL's [`time()`](https://prometheus.io/docs/prometheus/latest/querying/functions/#time) and
[`vector()`](https://prometheus.io/docs/prometheus/latest/querying/functions/#vector) functions.

#### Function type signature

```flux
() => stream[{_value: float, _time: time, _stop: time, _start: time}]
```

#### Parameters

No parameters provided.

### `extrapolatedRate()`

extrapolatedRate is a helper function that calculates extrapolated rates over
counters and is used to implement PromQL's
[`rate()`](https://prometheus.io/docs/prometheus/latest/querying/functions/#rate),
[`delta()`](https://prometheus.io/docs/prometheus/latest/querying/functions/#increase),
and [`increase()`](https://prometheus.io/docs/prometheus/latest/querying/functions/#delta) functions.

#### Function type signature

```flux
(
    <-tables: stream[{A with _value: float, _time: time, _stop: time, _start: time}],
    ?isCounter: bool,
    ?isRate: bool,
) => stream[{B with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
| isCounter | isCounter: Data represents a counter. | No |
| isRate | isRate: Data represents a rate. | No |
### `holtWinters()`

holtWinters implements functionality equivalent to
[PromQL's `holt_winters()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#holt_winters).

#### Function type signature

```flux
(
    <-tables: stream[{A with _value: float, _time: time}],
    ?smoothingFactor: float,
    ?trendFactor: float,
) => stream[{B with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
| smoothingFactor | smoothingFactor: Exponential smoothing factor. | No |
| trendFactor | trendFactor: Trend factor. | No |
### `instantRate()`

instantRate is a helper function that calculates instant rates over
counters and is used to implement PromQL's
[`irate()`](https://prometheus.io/docs/prometheus/latest/querying/functions/#irate) and
[`idelta()`](https://prometheus.io/docs/prometheus/latest/querying/functions/#idelta) functions.

#### Function type signature

```flux
(<-tables: stream[{A with _value: float, _time: time}], ?isRate: bool) => stream[{B with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Defaults is piped-forward data (`<-`). | No |
| isRate | isRate: Data represents a rate. | No |
### `join()`

join joins two streams of tables on the **group key and `_time` column**.
See [`experimental.join`](https://docs.influxdata.com/flux/v0.x/stdlib/experimental/join/).

#### Function type signature

```flux
(fn: (left: A, right: B) => C, left: stream[A], right: stream[B]) => stream[C] where A: Record, B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| left | left: First of two streams of tables to join. | Yes |
| right | right: Second of two streams of tables to join. | Yes |
| fn | fn: Function with left and right arguments that maps a new output record
  using values from the `left` and `right` input records.
  The return value must be a record. | Yes |
### `labelReplace()`

labelReplace implements functionality equivalent to
[PromQL's `label_replace()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#label_replace).

#### Function type signature

```flux
(
    <-tables: stream[{A with _value: float}],
    destination: string,
    regex: string,
    replacement: string,
    source: string,
) => stream[{B with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
| source | source: Input label. | Yes |
| destination | destination: Output label. | Yes |
| regex | regex: Pattern as a regex string. | Yes |
| replacement | replacement: Replacement value. | Yes |
### `linearRegression()`

linearRegression implements linear regression functionality required to implement
PromQL's [`deriv()`](https://prometheus.io/docs/prometheus/latest/querying/functions/#deriv)
and [`predict_linear()`](https://prometheus.io/docs/prometheus/latest/querying/functions/#predict_linear) functions.

#### Function type signature

```flux
(
    <-tables: stream[{A with _value: float, _time: time, _stop: time}],
    ?fromNow: float,
    ?predict: bool,
) => stream[{B with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
| predict | predict: Output should contain a prediction. | No |
| fromNow | fromNow: Time as a floating point value. | No |
### `promHistogramQuantile()`

promHistogramQuantile implements functionality equivalent to
[PromQL's `histogram_quantile()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile).

#### Function type signature

```flux
(
    <-tables: stream[A],
    ?countColumn: string,
    ?quantile: float,
    ?upperBoundColumn: string,
    ?valueColumn: string,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
| quantile | quantile: Quantile to compute (`[0.0 - 1.0]`). | No |
| countColumn | countColumn: Count column name. | No |
| upperBoundColumn | upperBoundColumn: Upper bound column name. | No |
| valueColumn | valueColumn: Output value column name. | No |
### `promqlDayOfMonth()`

promqlDayOfMonth implements functionality equivalent to
[PromQL's `day_of_month()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month).

#### Function type signature

```flux
(timestamp: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| timestamp | timestamp: Time as a floating point value. | Yes |
### `promqlDayOfWeek()`

promqlDayOfWeek implements functionality equivalent to
[PromQL's `day_of_week()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week).

#### Function type signature

```flux
(timestamp: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| timestamp | timestamp: Time as a floating point value. | Yes |
### `promqlDaysInMonth()`

promqlDaysInMonth implements functionality equivalent to
[PromQL's `days_in_month()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month).

#### Function type signature

```flux
(timestamp: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| timestamp | timestamp: Time as a floating point value. | Yes |
### `promqlHour()`

promqlHour implements functionality equivalent to
[PromQL's `hour()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#hour).

#### Function type signature

```flux
(timestamp: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| timestamp | timestamp: Time as a floating point value. | Yes |
### `promqlMinute()`

promqlMinute implements functionality equivalent to
[PromQL's `minute()` function]( https://prometheus.io/docs/prometheus/latest/querying/functions/#minute).

#### Function type signature

```flux
(timestamp: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| timestamp | timestamp: Time as a floating point value. | Yes |
### `promqlMonth()`

promqlMonth implements functionality equivalent to
[PromQL's `month()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#month).

#### Function type signature

```flux
(timestamp: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| timestamp | timestamp: Time as a floating point value. | Yes |
### `promqlYear()`

promqlYear implements functionality equivalent to
[PromQL's `year()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#year).

#### Function type signature

```flux
(timestamp: float) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| timestamp | timestamp: Time as a floating point value. | Yes |
### `quantile()`

quantile accounts checks for quantile values that are out of range, above 1.0 or
below 0.0, by either returning positive infinity or negative infinity in the `_value`
column respectively. `q` must be a float.

#### Function type signature

```flux
(<-tables: stream[A], q: float, ?method: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
| q | q: Quantile to compute (`[0.0 - 1.0]`). | Yes |
| method | method: Quantile method to use. | No |
### `resets()`

resets implements functionality equivalent to
[PromQL's `resets()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#resets).

#### Function type signature

```flux
(<-tables: stream[{A with _value: float}]) => stream[{B with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Defaults is piped-forward data (`<-`). | No |
### `timestamp()`

timestamp implements functionality equivalent to
[PromQL's `timestamp()` function](https://prometheus.io/docs/prometheus/latest/querying/functions/#timestamp).

#### Function type signature

```flux
(<-tables: stream[{A with _value: float}]) => stream[{A with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Defaults is piped-forward data (`<-`). | No |
