## `monitor` package

Package monitor provides tools for monitoring and alerting with InfluxDB.

Import the `monitor` package:

```flux
import "monitor"
```

### Functions

### `bucket()`

bucket is the default bucket to store InfluxDB monitoring data in.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `check()`

check checks input data and assigns a level (`ok`, `info`, `warn`, or `crit`)
to each row based on predicate functions.

#### Function type signature

```flux
(
    <-tables: stream[J],
    data: {A with tags: E, _type: D, _check_name: C, _check_id: B},
    messageFn: (
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
    ) => I,
    ?crit: (r: {F with _time: H, _measurement: G}) => bool,
    ?info: (r: {F with _time: H, _measurement: G}) => bool,
    ?ok: (r: {F with _time: H, _measurement: G}) => bool,
    ?warn: (r: {F with _time: H, _measurement: G}) => bool,
) => stream[{
    F with
    _type: D,
    _time: H,
    _time: time,
    _source_timestamp: int,
    _source_measurement: G,
    _message: I,
    _measurement: G,
    _measurement: string,
    _level: string,
    _check_name: C,
    _check_id: B,
}] where E: Record, J: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| crit | crit: Predicate function that determines `crit` status. Default is `(r) => false`. | No |
| warn | warn: Predicate function that determines `warn` status. Default is `(r) => false`. | No |
| info | info: Predicate function that determines `info` status. Default is `(r) => false`. | No |
| ok | ok: Predicate function that determines `ok` status. `Default is (r) => true`. | No |
| messageFn | messageFn: Predicate function that constructs a message to append to each row. | Yes |
| data | data: Check data to append to output used to identify this check. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `deadman()`

deadman detects when a group stops reporting data.
It takes a stream of tables and reports if groups have been observed since time `t`.

#### Function type signature

```flux
(<-tables: stream[{B with _time: C}], t: A) => stream[{B with dead: bool, _time: C}] where A: Comparable, C: Comparable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time threshold for the deadman check. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `from()`

from retrieves check statuses stored in the `statuses` measurement in the
`_monitoring` bucket.

#### Function type signature

```flux
(
    start: A,
    ?fn: (
        r: {
            B with
            _value: C,
            _time: time,
            _stop: time,
            _start: time,
            _measurement: string,
            _field: string,
        },
    ) => bool,
    ?stop: D,
) => stream[E] where E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| start | start: Earliest time to include in results. | Yes |
| stop | stop: Latest time to include in results. Default is `now()`. | No |
| fn | fn: Predicate function that evaluates true or false. | No |
### `levelCrit()`

levelCrit is the string representation of the "crit" level.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `levelInfo()`

levelInfo is the string representation of the "info" level.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `levelOK()`

levelOK is the string representation of the "ok" level.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `levelUnknown()`

levelUnknown is the string representation of the an unknown level.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `levelWarn()`

levelWarn is the string representation of the "warn" level.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `log()`

log persists notification events to an InfluxDB bucket.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `logs()`

logs retrieves notification events stored in the `notifications` measurement
in the `_monitoring` bucket.

#### Function type signature

```flux
(
    fn: (
        r: {
            A with
            _value: B,
            _time: time,
            _stop: time,
            _start: time,
            _measurement: string,
            _field: string,
        },
    ) => bool,
    start: C,
    ?stop: D,
) => stream[E] where E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| start | start: Earliest time to include in results. | Yes |
| stop | stop: Latest time to include in results. Default is `now()`. | No |
| fn | fn: Predicate function that evaluates true or false. | Yes |
### `notify()`

notify sends a notification to an endpoint and logs it in the `notifications`
measurement in the `_monitoring` bucket.

#### Function type signature

```flux
(
    <-tables: stream[E],
    data: A,
    endpoint: (<-: stream[{B with _time: C, _time: time, _status_timestamp: int, _measurement: string}]) => stream[D],
) => stream[D] where A: Record, D: Record, E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| endpoint | endpoint: A function that constructs and sends the notification to an endpoint. | Yes |
| data | data: Notification data to append to the output. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `stateChanges()`

stateChanges detects state changes in a stream of data with a `_level` column
and outputs records that change from `fromLevel` to `toLevel`.

#### Function type signature

```flux
(<-tables: stream[{C with _level: D}], ?fromLevel: A, ?toLevel: B) => stream[E] where A: Equatable, B: Equatable, D: Equatable, E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fromLevel | fromLevel: Level to detect a change from. Default is `"any"`. | No |
| toLevel | toLevel: Level to detect a change to. Default is `"any"`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `stateChangesOnly()`

stateChangesOnly takes a stream of tables that contains a _level column
and returns a stream of tables grouped by `_level` where each record
represents a state change.

#### Function type signature

```flux
(<-tables: stream[{A with _level: B}]) => stream[C] where B: Equatable, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `write()`

write persists check statuses to an InfluxDB bucket.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
