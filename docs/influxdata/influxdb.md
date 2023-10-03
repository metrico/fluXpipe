## `influxdb` package

Package influxdb provides functions designed for working with InfluxDB and
analyzing InfluxDB metadata.

Import the `influxdb` package:

```flux
import "influxdb"
```

### Functions

### `buckets()`

buckets returns a list of buckets in the specified organization.

#### Function type signature

```flux
(
    ?host: string,
    ?org: string,
    ?orgID: string,
    ?token: string,
) => stream[{
    retentionPolicy: string,
    retentionPeriod: int,
    organizationID: string,
    name: string,
    id: string,
}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| org | org: Organization name. Default is the current organization. | No |
| orgID | orgID: Organization ID. Default is the ID of the current organization. | No |
| host | host: URL of the InfluxDB instance. | No |
| token | token: InfluxDB API token. | No |
### `cardinality()`

cardinality returns the series cardinality of data stored in InfluxDB.

#### Function type signature

```flux
(
    start: A,
    ?bucket: string,
    ?bucketID: string,
    ?host: string,
    ?org: string,
    ?orgID: string,
    ?predicate: (r: {B with _value: C, _measurement: string, _field: string}) => bool,
    ?stop: D,
    ?token: string,
) => stream[{_value: int, _stop: time, _start: time}] where A: Timeable, D: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Bucket to query cardinality from. | No |
| bucketID | bucketID: String-encoded bucket ID to query cardinality from. | No |
| org | org: Organization name. | No |
| orgID | orgID: String-encoded organization ID. | No |
| host | host: URL of the InfluxDB instance to query. | No |
| token | token: InfluxDB API token. | No |
| start | start: Earliest time to include when calculating cardinality. | Yes |
| stop | stop: Latest time to include when calculating cardinality. | No |
| predicate | predicate: Predicate function that filters records.
     Default is `(r) => true`. | No |
### `from()`

from queries data from an InfluxDB data source.

#### Function type signature

```flux
(
    ?bucket: string,
    ?bucketID: string,
    ?host: string,
    ?org: string,
    ?orgID: string,
    ?token: string,
) => stream[{A with _value: B, _time: time, _measurement: string, _field: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Name of the bucket to query.
  _`bucket` and `bucketID` are mutually exclusive_. | No |
| bucketID | bucketID: String-encoded bucket ID to query.
  _`bucket` and `bucketID` are mutually exclusive_. | No |
| host | host: URL of the InfluxDB instance to query. | No |
| org | org: Organization name.
  _`org` and `orgID` are mutually exclusive_. | No |
| orgID | orgID: String-encoded organization ID to query.
  _`org` and `orgID` are mutually exclusive_. | No |
| token | token: InfluxDB API token. | No |
### `to()`

to writes data to an InfluxDB Cloud or 2.x bucket and returns the written data.

#### Function type signature

```flux
(
    <-tables: stream[A],
    ?bucket: string,
    ?bucketID: string,
    ?fieldFn: (r: A) => B,
    ?host: string,
    ?measurementColumn: string,
    ?org: string,
    ?orgID: string,
    ?tagColumns: [string],
    ?timeColumn: string,
    ?token: string,
) => stream[A] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Name of the bucket to write to.
  _`bucket` and `bucketID` are mutually exclusive_. | No |
| bucketID | bucketID: String-encoded bucket ID to to write to.
  _`bucket` and `bucketID` are mutually exclusive_. | No |
| host | host: URL of the InfluxDB instance to write to. | No |
| org | org: Organization name.
  _`org` and `orgID` are mutually exclusive_. | No |
| orgID | orgID: String-encoded organization ID to query.
  _`org` and `orgID` are mutually exclusive_. | No |
| token | token: InfluxDB API token. | No |
| timeColumn | timeColumn: Time column of the output. Default is `"_time"`. | No |
| measurementColumn | measurementColumn: Measurement column of the output. Default is `"_measurement"`. | No |
| tagColumns | tagColumns: Tag columns in the output. Defaults to all columns with type
  `string`, excluding all value columns and columns identified by `fieldFn`. | No |
| fieldFn | fieldFn: Function that maps a field key to a field value and returns a record.
  Default is `(r) => ({ [r._field]: r._value })`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `wideTo()`

wideTo writes wide data to an InfluxDB 2.x or InfluxDB Cloud bucket.
Wide data is _pivoted_ in that its fields are represented as columns making the table wider.

#### Function type signature

```flux
(
    <-tables: stream[A],
    ?bucket: string,
    ?bucketID: string,
    ?host: string,
    ?org: string,
    ?orgID: string,
    ?token: string,
) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Name of the bucket to write to.
  _`bucket` and `bucketID` are mutually exclusive_. | No |
| bucketID | bucketID: String-encoded bucket ID to to write to.
  _`bucket` and `bucketID` are mutually exclusive_. | No |
| host | host: URL of the InfluxDB instance to write to. | No |
| org | org: Organization name.
  _`org` and `orgID` are mutually exclusive_. | No |
| orgID | orgID: String-encoded organization ID to query.
  _`org` and `orgID` are mutually exclusive_. | No |
| token | token: InfluxDB API token. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
