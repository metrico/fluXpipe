## `query` package

Package query provides functions meant to simplify common InfluxDB queries.

Import the `query` package:

```flux
import "query"
```

### Functions

### `filterFields()`

filterFields filters input data by field.

#### Function type signature

```flux
(<-table: stream[{B with _field: A}], ?fields: [A]) => stream[{B with _field: A}] where A: Nullable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fields | fields: Fields to filter by. Default is `[]`. | No |
| table | table: Input data. Default is piped-forward data (`<-`). | No |
### `filterMeasurement()`

filterMeasurement filters input data by measurement.

#### Function type signature

```flux
(<-table: stream[{B with _measurement: C}], measurement: A) => stream[{B with _measurement: C}] where A: Equatable, C: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| measurement | measurement: InfluxDB measurement name to filter by. | Yes |
| table | table: Input data. Default is piped-forward data (`<-`). | No |
### `fromRange()`

fromRange returns all data from a specified bucket within given time bounds.

#### Function type signature

```flux
(
    bucket: string,
    start: A,
    ?stop: B,
) => stream[{
    C with
    _value: D,
    _time: time,
    _stop: time,
    _start: time,
    _measurement: string,
    _field: string,
}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: InfluxDB bucket name. | Yes |
| start | start: Earliest time to include in results. | Yes |
| stop | stop: Latest time to include in results. Default is `now()`. | No |
### `inBucket()`

inBucket queries data from a specified InfluxDB bucket within given time bounds,
filters data by measurement, field, and optional predicate expressions.

#### Function type signature

```flux
(
    bucket: string,
    measurement: A,
    start: B,
    ?fields: [string],
    ?predicate: (
        r: {
            C with
            _value: D,
            _time: time,
            _stop: time,
            _start: time,
            _measurement: string,
            _field: string,
        },
    ) => bool,
    ?stop: E,
) => stream[{
    C with
    _value: D,
    _time: time,
    _stop: time,
    _start: time,
    _measurement: string,
    _field: string,
}] where A: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: InfluxDB bucket name. | Yes |
| measurement | measurement: InfluxDB measurement name to filter by. | Yes |
| start | start: Earliest time to include in results. | Yes |
| stop | stop: Latest time to include in results. Default is `now()`. | No |
| fields | fields: Fields to filter by. Default is `[]`. | No |
| predicate | predicate: Predicate function that evaluates column values and returns `true` or `false`.
  Default is `(r) => true`. | No |
