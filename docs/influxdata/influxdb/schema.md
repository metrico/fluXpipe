## `schema` package

Package schema provides functions for exploring your InfluxDB data schema.

Import the `schema` package:

```flux
import "schema"
```

### Functions

### `fieldKeys()`

fieldKeys returns field keys in a bucket.

#### Function type signature

```flux
(
    bucket: string,
    ?predicate: (
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
    ?start: C,
    ?stop: D,
) => stream[E] where E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Bucket to list field keys from. | Yes |
| predicate | predicate: Predicate function that filters field keys.
  Default is `(r) => true`. | No |
| start | start: Oldest time to include in results. Default is `-30d`. | No |
| stop | stop: Newest time include in results.
    The stop time is exclusive, meaning values with a time equal to stop time are excluded from the results.
    Default is `now()`. | No |
### `fieldsAsCols()`

fieldsAsCols is a special application of `pivot()` that pivots input data
on `_field` and `_time` columns to align fields within each input table that
have the same timestamp.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `measurementFieldKeys()`

measurementFieldKeys returns a list of fields in a measurement.

#### Function type signature

```flux
(bucket: string, measurement: A, ?start: B, ?stop: C) => stream[D] where A: Equatable, D: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Bucket to retrieve field keys from. | Yes |
| measurement | measurement: Measurement to list field keys from. | Yes |
| start | start: Oldest time to include in results. Default is `-30d`. | No |
| stop | stop: Newest time include in results.
    The stop time is exclusive, meaning values with a time equal to stop time are excluded from the results.
    Default is `now()`. | No |
### `measurementTagKeys()`

measurementTagKeys returns the list of tag keys for a specific measurement.

#### Function type signature

```flux
(bucket: string, measurement: A, ?start: B, ?stop: C) => stream[D] where A: Equatable, D: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Bucket to return tag keys from for a specific measurement. | Yes |
| measurement | measurement: Measurement to return tag keys from. | Yes |
| start | start: Oldest time to include in results. Default is `-30d`. | No |
| stop | stop: Newest time include in results.
    The stop time is exclusive, meaning values with a time equal to stop time are excluded from the results.
    Default is `now()`. | No |
### `measurementTagValues()`

measurementTagValues returns a list of tag values for a specific measurement.

#### Function type signature

```flux
(
    bucket: string,
    measurement: A,
    tag: string,
    ?start: B,
    ?stop: C,
) => stream[D] where A: Equatable, D: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Bucket to return tag values from for a specific measurement. | Yes |
| measurement | measurement: Measurement to return tag values from. | Yes |
| tag | tag: Tag to return all unique values from. | Yes |
| start | start: Oldest time to include in results. Default is `-30d`. | No |
| stop | stop: Newest time include in results.
    The stop time is exclusive, meaning values with a time equal to stop time are excluded from the results.
    Default is `now()`. | No |
### `measurements()`

measurements returns a list of measurements in a specific bucket.

#### Function type signature

```flux
(bucket: string, ?start: A, ?stop: B) => stream[C] where C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Bucket to retrieve measurements from. | Yes |
| start | start: Oldest time to include in results. Default is `-30d`. | No |
| stop | stop: Newest time include in results.
    The stop time is exclusive, meaning values with a time equal to stop time are excluded from the results.
    Default is `now()`. | No |
### `tagKeys()`

tagKeys returns a list of tag keys for all series that match the `predicate`.

#### Function type signature

```flux
(
    bucket: string,
    ?predicate: (
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
    ?start: C,
    ?stop: D,
) => stream[E] where E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Bucket to return tag keys from. | Yes |
| predicate | predicate: Predicate function that filters tag keys.
  Default is `(r) => true`. | No |
| start | start: Oldest time to include in results. Default is `-30d`. | No |
| stop | stop: Newest time include in results.
    The stop time is exclusive, meaning values with a time equal to stop time are excluded from the results.
    Default is `now()`. | No |
### `tagValues()`

tagValues returns a list of unique values for a given tag.

#### Function type signature

```flux
(
    bucket: string,
    tag: string,
    ?predicate: (
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
    ?start: C,
    ?stop: D,
) => stream[E] where E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Bucket to return unique tag values from. | Yes |
| tag | tag: Tag to return unique values from. | Yes |
| predicate | predicate: Predicate function that filters tag values.
  Default is `(r) => true`. | No |
| start | start: Oldest time to include in results. Default is `-30d`. | No |
| stop | stop: Newest time include in results.
    The stop time is exclusive, meaning values with a time equal to stop time are excluded from the results.
    Default is `now()`. | No |
