## `experimental` package

Package experimental includes experimental functions and packages.

Import the `experimental` package:

```flux
import "experimental"
```

### Functions

### `addDuration()`

addDuration adds a duration to a time value and returns the resulting time value.

#### Function type signature

```flux
(d: duration, to: A, ?location: {zone: string, offset: duration}) => time where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| d | d: Duration to add. | Yes |
| to | to: Time to add the duration to. | Yes |
| location | location: Location to use for the time value. | No |
### `alignTime()`

alignTime shifts time values in input tables to all start at a common start time.

#### Function type signature

```flux
(<-tables: stream[B], ?alignTo: A) => stream[C] where B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| alignTo | alignTo: Time to align tables to. Default is `1970-01-01T00:00:00Z`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `catch()`

catch calls a function and returns any error as a string value.
If the function does not error the returned value is made into a string and returned.

#### Function type signature

```flux
(fn: () => A) => {value: A, msg: string, code: uint}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fn | fn: Function to call. | Yes |
### `chain()`

chain runs two queries in a single Flux script sequentially and outputs the
results of the second query.

#### Function type signature

```flux
(first: stream[A], second: stream[B]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| first | first: First query to execute. | Yes |
| second | second: Second query to execute. | Yes |
### `count()`

count returns the number of records in each input table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: int}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `diff()`

diff takes two table streams as input and produces a diff.

#### Function type signature

```flux
(<-got: stream[A], want: stream[A]) => stream[{A with _diff: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| want | want: Input stream for the `-` side of the diff. | Yes |
| got | got: Input stream for the `+` side of the diff. | No |
### `distinct()`

distinct returns unique values from the `_value` column.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `fill()`

fill replaces all null values in the `_value` column with a non-null value.

#### Function type signature

```flux
(<-tables: stream[{B with _value: A}], ?usePrevious: bool, ?value: A) => stream[{B with _value: A}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| value | value: Value to replace null values with.
  Data type must match the type of the `_value` column. | No |
| usePrevious | usePrevious: Replace null values with the value of the previous non-null row. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `first()`

first returns the first record with a non-null value in the `_value` column
for each input table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `group()`

group introduces an `extend` mode to the existing `group()` function.

#### Function type signature

```flux
(<-tables: stream[A], columns: [string], mode: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| columns | columns: List of columns to use in the grouping operation. Default is `[]`. | Yes |
| mode | mode: Grouping mode. `extend` is the only mode available to `experimental.group()`. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `histogram()`

histogram approximates the cumulative distribution of a dataset by counting
data frequencies for a list of bins.

#### Function type signature

```flux
(<-tables: stream[{A with _value: float}], bins: [float], ?normalize: bool) => stream[{A with le: float, _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bins | bins: List of upper bounds to use when computing histogram frequencies,
  including the maximum value of the data set. | Yes |
| normalize | normalize: Convert count values into frequency values between 0 and 1.
  Default is `false`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `histogramQuantile()`

histogramQuantile approximates a quantile given a histogram with the
cumulative distribution of the dataset.

#### Function type signature

```flux
(
    <-tables: stream[{A with le: float, _value: float}],
    ?minValue: float,
    ?quantile: float,
) => stream[{A with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| quantile | quantile: Quantile to compute (`[0.0 - 1.0]`). | No |
| minValue | minValue: Assumed minimum value of the dataset. Default is `0.0`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `integral()`

integral computes the area under the curve per unit of time of subsequent non-null records.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B, _time: time}], ?interpolate: string, ?unit: duration) => stream[{A with _value: B}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| unit | unit: Time duration used to compute the integral. | No |
| interpolate | interpolate: Type of interpolation to use. Default is `""` (no interpolation). | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `join()`

join joins two streams of tables on the **group key and `_time` column**.

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
### `kaufmansAMA()`

kaufmansAMA calculates the Kaufman's Adaptive Moving Average (KAMA) of input
tables using the `_value` column in each table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}], n: int) => stream[{A with _value: float}] where B: Numeric
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Period or number of points to use in the calculation. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `last()`

last returns the last record with a non-null value in the `_value` column
for each input table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `max()`

max returns the record with the highest value in the `_value` column for each
input table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `mean()`

mean computes the mean or average of non-null values in the `_value` column
of each input table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: float}]) => stream[{A with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `min()`

min returns the record with the lowest value in the `_value` column for each
input table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `mode()`

mode computes the mode or value that occurs most often in the `_value` column
in each input table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `objectKeys()`

objectKeys returns an array of property keys in a specified record.

#### Function type signature

```flux
(o: A) => [string] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| o | o: Record to return property keys from. | Yes |
### `preview()`

preview limits the number of rows and tables in the stream.

#### Function type signature

```flux
(<-tables: stream[A], ?nrows: int, ?ntables: int) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| nrows | nrows: Maximum number of rows per table to return. Default is `5`. | No |
| ntables | ntables: Maximum number of tables to return.
  Default is `5`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `quantile()`

quantile returns non-null records with values in the `_value` column that
fall within the specified quantile or represent the specified quantile.

#### Function type signature

```flux
(
    <-tables: stream[{A with _value: float}],
    q: float,
    ?compression: float,
    ?method: string,
) => stream[{A with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| q | q: Quantile to compute (`[0 - 1]`). | Yes |
| method | method: Computation method. Default is `estimate_tdigest`. | No |
| compression | compression: Number of centroids to use when compressing the dataset.
  Default is `1000.0`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `set()`

set sets multiple static column values on all records.

#### Function type signature

```flux
(<-tables: stream[B], o: A) => stream[C] where A: Record, B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| o | o: Record that defines the columns and values to set. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `skew()`

skew returns the skew of non-null values in the `_value` column for each
input table as a float.

#### Function type signature

```flux
(<-tables: stream[{A with _value: float}]) => stream[{A with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `spread()`

spread returns the difference between the minimum and maximum values in the
`_value` column for each input table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B}] where B: Numeric
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `stddev()`

stddev returns the standard deviation of non-null values in the `_value`
column for each input table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: float}], ?mode: string) => stream[{A with _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| mode | mode: Standard deviation mode or type of standard deviation to calculate.
  Default is `sample`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `subDuration()`

subDuration subtracts a duration from a time value and returns the resulting time value.

#### Function type signature

```flux
(d: duration, from: A, ?location: {zone: string, offset: duration}) => time where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| from | from: Time to subtract the duration from. | Yes |
| d | d: Duration to subtract. | Yes |
| location | location: Location to use for the time value. | No |
### `sum()`

sum returns the sum of non-null values in the `_value` column for each input table.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B}] where B: Numeric
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `to()`

to writes _pivoted_ data to an InfluxDB 2.x or InfluxDB Cloud bucket.

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
### `unique()`

unique returns all records containing unique values in the `_value` column.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `unpivot()`

unpivot creates `_field` and `_value` columns pairs using all columns (other than `_time`)
_not_ in the group key.
The `_field` column contains the original column label and the `_value` column
contains the original column value.

#### Function type signature

```flux
(<-tables: stream[{A with _time: time}], ?otherColumns: [string]) => stream[{B with _value: C, _field: string}] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
| otherColumns | otherColumns: List of column names that are not in the group key but are also not field columns. Default is `["_time"]`. | No |
### `window()`

window groups records based on time.

#### Function type signature

```flux
(
    <-tables: stream[{A with _time: time, _stop: time, _start: time}],
    ?createEmpty: bool,
    ?every: duration,
    ?location: {zone: string, offset: duration},
    ?offset: duration,
    ?period: duration,
) => stream[{A with _time: time, _stop: time, _start: time}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| every | every: Duration of time between windows. Default is the `0s`. | No |
| period | period: Duration of the window. Default is `0s`. | No |
| offset | offset: Duration to shift the window boundaries by. Default is 0s. | No |
| location | location: Location used to determine timezone. Default is the `location` option. | No |
| createEmpty | createEmpty: Create empty tables for empty windows. Default is `false`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
