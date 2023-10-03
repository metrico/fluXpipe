## `tickscript` package

Package tickscript provides functions to help migrate
Kapacitor [TICKscripts](https://docs.influxdata.com/kapacitor/latest/tick/) to Flux tasks.

Import the `tickscript` package:

```flux
import "tickscript"
```

### Functions

### `alert()`

alert identifies events of varying severity levels
and writes them to the `statuses` measurement in the InfluxDB `_monitoring`
system bucket.

#### Function type signature

```flux
(
    <-tables: stream[M],
    check: {A with tags: E, _type: D, _check_name: C, _check_id: B},
    ?crit: (r: {F with _time: H, _measurement: G}) => bool,
    ?details: (r: {I with id: J, _check_name: C, _check_id: B}) => K,
    ?id: (r: {I with _check_name: C, _check_id: B}) => J,
    ?info: (r: {F with _time: H, _measurement: G}) => bool,
    ?message: (
        r: {
            F with
            _type: D,
            _time: H,
            _time: time,
            _source_timestamp: int,
            _source_measurement: G,
            _measurement: G,
            _measurement: string,
            _level: string,
            _check_name: C,
            _check_id: B,
        },
    ) => L,
    ?ok: (r: {F with _time: H, _measurement: G}) => bool,
    ?topic: string,
    ?warn: (r: {F with _time: H, _measurement: G}) => bool,
) => stream[{
    F with
    _type: D,
    _time: H,
    _time: time,
    _source_timestamp: int,
    _source_measurement: G,
    _message: L,
    _measurement: G,
    _measurement: string,
    _level: string,
    _check_name: C,
    _check_id: B,
}] where E: Record, I: Record, M: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| check | check: InfluxDB check data.
  See `tickscript.defineCheck()`. | Yes |
| id | id: Function that returns the InfluxDB check ID provided by the check record.
  Default is `(r) => "${r._check_id}"`. | No |
| details | details: Function to return the InfluxDB check details using data from input rows.
  Default is `(r) => ""`. | No |
| message | message: Function to return the InfluxDB check message using data from input rows.
  Default is `(r) => "Threshold Check: ${r._check_name} is: ${r._level}"`. | No |
| crit | crit: Predicate function to determine `crit` status. Default is `(r) => false`. | No |
| warn | warn: Predicate function to determine `warn` status. Default is `(r) => false`. | No |
| info | info: Predicate function to determine `info` status. Default is `(r) => false`. | No |
| ok | ok: Predicate function to determine `ok` status. Default is `(r) => true`. | No |
| topic | topic: Check topic. Default is `""`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `compute()`

compute is an alias for `tickscript.select()` that changes a column’s name and
optionally applies an aggregate or selector function.

#### Function type signature

```flux
(<-tables: B, as: string, ?column: A, ?fn: (<-: B, column: A) => stream[C]) => stream[D] where A: Equatable, C: Record, D: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| as | as: New column name. | Yes |
| column | column: Column to operate on. Default is `_value`. | No |
| fn | fn: Aggregate or selector function to apply. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `deadman()`

deadman detects low data throughput and writes a point with a critical status to
the InfluxDB `_monitoring` system bucket.

#### Function type signature

```flux
(
    <-tables: stream[M],
    check: {A with tags: E, _type: D, _check_name: C, _check_id: B},
    measurement: string,
    ?id: (r: {F with _check_name: C, _check_id: B}) => G,
    ?message: (
        r: {
            H with
            dead: bool,
            _type: D,
            _time: J,
            _time: time,
            _source_timestamp: int,
            _source_measurement: I,
            _measurement: I,
            _measurement: string,
            _level: string,
            _check_name: C,
            _check_id: B,
        },
    ) => K,
    ?threshold: L,
    ?topic: string,
) => stream[{
    H with
    dead: bool,
    _type: D,
    _time: J,
    _time: time,
    _source_timestamp: int,
    _source_measurement: I,
    _message: K,
    _measurement: I,
    _measurement: string,
    _level: string,
    _check_name: C,
    _check_id: B,
}] where E: Record, F: Record, L: Comparable + Equatable, M: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| check | check: InfluxDB check data. See `tickscript.defineCheck()`. | Yes |
| measurement | measurement: Measurement name. Should match the queried measurement. | Yes |
| threshold | threshold: Count threshold. Default is `0`. | No |
| id | id: Function that returns the InfluxDB check ID provided by the check record.
  Default is `(r) => "${r._check_id}"`. | No |
| message | message: Function that returns the InfluxDB check message using data from input rows.
  Default is `(r) => "Deadman Check: ${r._check_name} is: " + (if r.dead then "dead" else "alive")`. | No |
| topic | topic: Check topic. Default is `""`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `defineCheck()`

defineCheck creates custom check data required by `alert()` and `deadman()`.

#### Function type signature

```flux
(id: A, name: B, ?type: C) => {tags: {}, _type: C, _check_name: B, _check_id: A}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| id | id: InfluxDB check ID. | Yes |
| name | name: InfluxDB check name. | Yes |
| type | type: InfluxDB check type. Default is `custom`. | No |
### `groupBy()`

groupBy groups results by the `_measurement` column and other specified columns.

#### Function type signature

```flux
(<-tables: stream[A], columns: [string]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| columns | columns: List of columns to group by. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `join()`

join merges two input streams into a single output stream
based on specified columns with equal values and appends a new measurement name.

#### Function type signature

```flux
(measurement: A, tables: B, ?on: [string]) => stream[{C with _measurement: A}] where B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Map of two streams to join. | Yes |
| on | on: List of columns to join on. Default is `["_time"]`. | No |
| measurement | measurement: Measurement name to use in results. | Yes |
### `select()`

select changes a column’s name and optionally applies an aggregate or selector
function to values in the column.

#### Function type signature

```flux
(<-tables: B, as: string, ?column: A, ?fn: (<-: B, column: A) => stream[C]) => stream[D] where A: Equatable, C: Record, D: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to operate on. Default is `_value`. | No |
| fn | fn: Aggregate or selector function to apply. | No |
| as | as: New column name. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `selectWindow()`

selectWindow changes a column’s name, windows rows by time, and then applies an
aggregate or selector function the specified column for each window of time.

#### Function type signature

```flux
(
    <-tables: stream[D],
    as: string,
    defaultValue: A,
    every: duration,
    fn: (<-: stream[B], column: string) => stream[C],
    ?column: string,
) => stream[E] where B: Record, C: Record, D: Record, E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to operate on. Default is _value. | No |
| fn | fn: Aggregate or selector function to apply. | Yes |
| as | as: New column name. | Yes |
| every | every: Duration of windows. | Yes |
| defaultValue | defaultValue: Default fill value for null values in column.
  Must be the same data type as column. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
