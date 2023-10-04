## `influxdb` package

Package influxdb provides additional functions for querying data from InfluxDB.

Import the `influxdb` package:

```flux
import "influxdb"
```

### Functions

### `from()`

from retrieves data from an InfluxDB bucket between the `start` and `stop` times.

#### Function type signature

```flux
(
    bucket: string,
    start: A,
    ?host: string,
    ?org: string,
    ?stop: B,
    ?token: string,
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
| bucket | bucket: Name of the bucket to query. | Yes |
| start | start: Earliest time to include in results. | Yes |
| stop | stop: Latest time to include in results. Default is `now()`. | No |
| host | host: URL of the InfluxDB instance to query. | No |
| org | org: Organization name. | No |
| token | token: InfluxDB [API token](https://docs.influxdata.com/influxdb/latest/security/tokens/). | No |
### `select()`

select is an alternate implementation of `from()`,
`range()`, `filter()` and `pivot()` that returns pivoted query results and masks
the `_measurement`, `_start`, and `_stop` columns. Results are similar to those
returned by InfluxQL `SELECT` statements.

#### Function type signature

```flux
(
    from: string,
    m: A,
    start: B,
    ?fields: [string],
    ?host: string,
    ?org: string,
    ?stop: C,
    ?token: string,
    ?where: (
        r: {
            D with
            _value: E,
            _time: time,
            _stop: time,
            _start: time,
            _measurement: string,
            _field: string,
        },
    ) => bool,
) => stream[F] where A: Equatable, F: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| from | from: Name of the bucket to query. | Yes |
| start | start: Earliest time to include in results. | Yes |
| stop | stop: Latest time to include in results. Default is `now()`. | No |
| m | m: Name of the measurement to query. | Yes |
| fields | fields: List of fields to query. Default is`[]`. | No |
| where | where: Single argument predicate function that evaluates `true` or `false`
  and filters results based on tag values.
  Default is `(r) => true`. | No |
| host | host: URL of the InfluxDB instance to query. | No |
| org | org: Organization name. | No |
| token | token: InfluxDB [API token](https://docs.influxdata.com/influxdb/latest/security/tokens/). | No |
