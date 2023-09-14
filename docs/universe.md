## `universe` package

Package universe provides options and primitive functions that are
loaded into the Flux runtime by default and do not require an
import statement.

Import the `universe` package:

```flux
import "universe"
```

### Functions

### `aggregateWindow()`

aggregateWindow downsamples data by grouping data into fixed windows of time
and applying an aggregate or selector function to each window.

#### Function type signature

```flux
(
    <-tables: stream[D],
    every: duration,
    fn: (<-: stream[B], column: A) => stream[C],
    ?column: A,
    ?createEmpty: bool,
    ?location: {zone: string, offset: duration},
    ?offset: duration,
    ?period: duration,
    ?timeDst: string,
    ?timeSrc: string,
) => stream[E] where B: Record, C: Record, D: Record, E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| every | every: Duration of time between windows. | Yes |
| period | period: Duration of windows. Default is the `every` value. | No |
| offset | offset: Duration to shift the window boundaries by. Default is `0s`. | No |
| fn | fn: Aggregate or selector function to apply to each time window. | Yes |
| location | location: Location used to determine timezone. Default is the `location` option. | No |
| column | column: Column to operate on. | No |
| timeSrc | timeSrc: Column to use as the source of the new time value for aggregate values.
  Default is `_stop`. | No |
| timeDst | timeDst: Column to store time values for aggregate values in.
  Default is `_time`. | No |
| createEmpty | createEmpty: Create empty tables for empty window. Default is `true`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `bool()`

bool converts a value to a boolean type.

#### Function type signature

```flux
(v: A) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to convert. | Yes |
### `bottom()`

bottom sorts each input table by specified columns and keeps the bottom `n`
records in each table.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?columns: [string]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of rows to return from each input table. | Yes |
| columns | columns: List of columns to sort by. Default is `["_value"]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `bytes()`

bytes converts a string value to a bytes type.

#### Function type signature

```flux
(v: A) => bytes
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to convert. | Yes |
### `chandeMomentumOscillator()`

chandeMomentumOscillator applies the technical momentum indicator developed
by Tushar Chande to input data.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?columns: [string]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Period or number of points to use in the calculation. | Yes |
| columns | columns: List of columns to operate on. Default is `["_value"]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `columns()`

columns returns the column labels in each input table.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Name of the output column to store column labels in.
  Default is "_value". | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `contains()`

contains tests if an array contains a specified value and returns `true` or `false`.

#### Function type signature

```flux
(set: [A], value: A) => bool where A: Nullable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| value | value: Value to search for. | Yes |
| set | set: Array to search. | Yes |
### `count()`

count returns the number of records in each input table.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to count values in and store the total count. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `cov()`

cov computes the covariance between two streams of tables.

#### Function type signature

```flux
(on: [string], x: A, y: B, ?pearsonr: bool) => stream[C] where C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: First input stream. | Yes |
| y | y: Second input stream. | Yes |
| on | on: List of columns to join on. | Yes |
| pearsonr | pearsonr: Normalize results to the Pearson R coefficient. Default is `false`. | No |
### `covariance()`

covariance computes the covariance between two columns.

#### Function type signature

```flux
(<-tables: stream[A], columns: [string], ?pearsonr: bool, ?valueDst: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| columns | columns: List of two columns to operate on. | Yes |
| pearsonr | pearsonr: Normalize results to the Pearson R coefficient. Default is `false`. | No |
| valueDst | valueDst: Column to store the result in. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `cumulativeSum()`

cumulativeSum  computes a running sum for non-null records in a table.

#### Function type signature

```flux
(<-tables: stream[A], ?columns: [string]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| columns | columns: List of columns to operate on. Default is `["_value"]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `derivative()`

derivative computes the rate of change per unit of time between subsequent
non-null records.

#### Function type signature

```flux
(
    <-tables: stream[A],
    ?columns: [string],
    ?initialZero: bool,
    ?nonNegative: bool,
    ?timeColumn: string,
    ?unit: duration,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| unit | unit: Time duration used to calculate the derivative. Default is `1s`. | No |
| nonNegative | nonNegative: Disallow negative derivative values. Default is `false`. | No |
| columns | columns: List of columns to operate on. Default is `["_value"]`. | No |
| timeColumn | timeColumn: Column containing time values to use in the calculation.
  Default is `_time`. | No |
| initialZero | initialZero: Use zero (0) as the initial value in the derivative calculation
  when the subsequent value is less than the previous value and `nonNegative` is
  `true`. Default is `false`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `die()`

die stops the Flux script execution and returns an error message.

#### Function type signature

```flux
(msg: string) => A
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| msg | msg: Error message to return. | Yes |
### `difference()`

difference returns the difference between subsequent values.

#### Function type signature

```flux
(
    <-tables: stream[A],
    ?columns: [string],
    ?initialZero: bool,
    ?keepFirst: bool,
    ?nonNegative: bool,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| nonNegative | nonNegative: Disallow negative differences. Default is `false`. | No |
| columns | columns: List of columns to operate on. Default is `["_value"]`. | No |
| keepFirst | keepFirst: Keep the first row in each input table. Default is `false`. | No |
| initialZero | initialZero: Use zero (0) as the initial value in the difference calculation
  when the subsequent value is less than the previous value and `nonNegative` is
  `true`. Default is `false`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `display()`

display returns the Flux literal representation of any value as a string.

#### Function type signature

```flux
(v: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to convert for display. | Yes |
### `distinct()`

distinct returns all unique values in a specified column.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to return unique values from. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `doubleEMA()`

doubleEMA returns the double exponential moving average (DEMA) of values in
the `_value` column grouped into `n` number of points, giving more weight to
recent data.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}], n: int) => stream[C] where B: Numeric, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of points to average. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `drop()`

drop removes specified columns from a table.

#### Function type signature

```flux
(<-tables: stream[A], ?columns: [string], ?fn: (column: string) => bool) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| columns | columns: List of columns to remove from input tables. Mutually exclusive with `fn`. | No |
| fn | fn: Predicate function with a `column` parameter that returns a boolean
  value indicating whether or not the column should be removed from input tables.
  Mutually exclusive with `columns`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `duplicate()`

duplicate duplicates a specified column in a table.

#### Function type signature

```flux
(<-tables: stream[A], as: string, column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to duplicate. | Yes |
| as | as: Name to assign to the duplicate column. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `duration()`

duration converts a value to a duration type.

#### Function type signature

```flux
(v: A) => duration
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to convert. | Yes |
### `elapsed()`

elapsed returns the time between subsequent records.

#### Function type signature

```flux
(<-tables: stream[A], ?columnName: string, ?timeColumn: string, ?unit: duration) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| unit | unit: Unit of time used in the calculation. Default is `1s`. | No |
| timeColumn | timeColumn: Column to use to compute the elapsed time. Default is `_time`. | No |
| columnName | columnName: Column to store elapsed times in. Default is `elapsed`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `exponentialMovingAverage()`

exponentialMovingAverage calculates the exponential moving average of `n`
number of values in the `_value` column giving more weight to more recent data.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}], n: int) => stream[{A with _value: B}] where B: Numeric
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of values to average. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `fill()`

fill replaces all null values in input tables with a non-null value.

#### Function type signature

```flux
(<-tables: stream[B], ?column: string, ?usePrevious: bool, ?value: A) => stream[C] where B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to replace null values in. Default is `_value`. | No |
| value | value: Constant value to replace null values with. | No |
| usePrevious | usePrevious: Replace null values with the previous non-null value.
  Default is `false`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `filter()`

filter filters data based on conditions defined in a predicate function (`fn`).

#### Function type signature

```flux
(<-tables: stream[A], fn: (r: A) => bool, ?onEmpty: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fn | fn: Single argument predicate function that evaluates `true` or `false`. | Yes |
| onEmpty | onEmpty: Action to take with empty tables. Default is `drop`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `findColumn()`

findColumn returns an array of values in a specified column from the first
table in a stream of tables that matches the specified predicate function.

#### Function type signature

```flux
(<-tables: stream[B], column: string, fn: (key: A) => bool) => [C] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to extract. | Yes |
| fn | fn: Predicate function to evaluate input table group keys. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `findRecord()`

findRecord returns a row at a specified index as a record from the first
table in a stream of tables that matches the specified predicate function.

#### Function type signature

```flux
(<-tables: stream[B], fn: (key: A) => bool, idx: int) => B where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| idx | idx: Index of the record to extract. | Yes |
| fn | fn: Predicate function to evaluate input table group keys. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `first()`

first returns the first non-null record from each input table.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to operate on. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `float()`

float converts a value to a float type.

#### Function type signature

```flux
(v: A) => float
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to convert. | Yes |
### `getColumn()`

getColumn extracts a specified column from a table as an array.

#### Function type signature

```flux
(<-table: stream[A], column: string) => [B] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to extract. | Yes |
| table | table: Input table. Default is piped-forward data (`<-`). | No |
### `getRecord()`

getRecord extracts a row at a specified index from a table as a record.

#### Function type signature

```flux
(<-table: stream[A], idx: int) => A where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| idx | idx: Index of the record to extract. | Yes |
| table | table: Input table. Default is piped-forward data (`<-`). | No |
### `group()`

group regroups input data by modifying group key of input tables.

#### Function type signature

```flux
(<-tables: stream[A], ?columns: [string], ?mode: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| columns | columns: List of columns to use in the grouping operation. Default is `[]`. | No |
| mode | mode: Grouping mode. Default is `by`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `highestAverage()`

highestAverage calculates the average of each input table and returns the
highest `n` averages.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?column: string, ?groupColumns: [string]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of records to return. | Yes |
| column | column: Column to evaluate. Default is `_value`. | No |
| groupColumns | groupColumns: List of columns to group by. Default is `[]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `highestCurrent()`

highestCurrent selects the last record from each input table and returns the
highest `n` records.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?column: string, ?groupColumns: [string]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of records to return. | Yes |
| column | column: Column to evaluate. Default is `_value`. | No |
| groupColumns | groupColumns: List of columns to group by. Default is `[]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `highestMax()`

highestMax selects the record with the highest value in the specified `column`
from each input table and returns the highest `n` records.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?column: string, ?groupColumns: [string]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of records to return. | Yes |
| column | column: Column to evaluate. Default is `_value`. | No |
| groupColumns | groupColumns: List of columns to group by. Default is `[]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `histogram()`

histogram approximates the cumulative distribution of a dataset by counting
data frequencies for a list of bins.

#### Function type signature

```flux
(
    <-tables: stream[A],
    bins: [float],
    ?column: string,
    ?countColumn: string,
    ?normalize: bool,
    ?upperBoundColumn: string,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column containing input values. Column must be of type float.
  Default is `_value`. | No |
| upperBoundColumn | upperBoundColumn: Column to store bin upper bounds in. Default is `le`. | No |
| countColumn | countColumn: Column to store bin counts in. Default is `_value`. | No |
| bins | bins: List of upper bounds to use when computing the histogram frequencies. | Yes |
| normalize | normalize: Convert counts into frequency values between 0 and 1.
  Default is `false`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `histogramQuantile()`

histogramQuantile approximates a quantile given a histogram that approximates
the cumulative distribution of the dataset.

#### Function type signature

```flux
(
    <-tables: stream[A],
    ?countColumn: string,
    ?minValue: float,
    ?onNonmonotonic: string,
    ?quantile: float,
    ?upperBoundColumn: string,
    ?valueColumn: string,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| quantile | quantile: Quantile to compute. Value must be between 0 and 1. | No |
| countColumn | countColumn: Column containing histogram bin counts. Default is `_value`. | No |
| upperBoundColumn | upperBoundColumn: Column containing histogram bin upper bounds.
  Default is `le`. | No |
| valueColumn | valueColumn: Column to store the computed quantile in. Default is `_value. | No |
| minValue | minValue: Assumed minimum value of the dataset. Default is `0.0`. | No |
| onNonmonotonic | onNonmonotonic: Describes behavior when counts are not monotonically increasing
  when sorted by upper bound. Default is `error`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `holtWinters()`

holtWinters applies the Holt-Winters forecasting method to input tables.

#### Function type signature

```flux
(
    <-tables: stream[A],
    interval: duration,
    n: int,
    ?column: string,
    ?seasonality: int,
    ?timeColumn: string,
    ?withFit: bool,
    ?withMinSSE: bool,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of values to predict. | Yes |
| interval | interval: Interval between two data points. | Yes |
| withFit | withFit: Return fitted data in results. Default is `false`. | No |
| column | column: Column to operate on. Default is `_value`. | No |
| timeColumn | timeColumn: Column containing time values to use in the calculating.
  Default is `_time`. | No |
| seasonality | seasonality: Number of points in a season. Default is `0`. | No |
| withMinSSE | withMinSSE: Return minSSE data in results. Default is `false`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `hourSelection()`

hourSelection filters rows by time values in a specified hour range.

#### Function type signature

```flux
(
    <-tables: stream[A],
    start: int,
    stop: int,
    ?location: {zone: string, offset: duration},
    ?timeColumn: string,
) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| start | start: First hour of the hour range (inclusive). Hours range from `[0-23]`. | Yes |
| stop | stop: Last hour of the hour range (inclusive). Hours range from `[0-23]`. | Yes |
| location | location: Location used to determine timezone. Default is the `location` option. | No |
| timeColumn | timeColumn: Column that contains the time value. Default is `_time`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `increase()`

increase returns the cumulative sum of non-negative differences between subsequent values.

#### Function type signature

```flux
(<-tables: stream[A], ?columns: [string]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| columns | columns: List of columns to operate on. Default is `["_value"]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `inf()`

inf represents a floating point value of infinity.

#### Function type signature

```flux
duration
```

#### Parameters

No parameters provided.

### `int()`

int converts a value to an integer type.

#### Function type signature

```flux
(v: A) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to convert. | Yes |
### `integral()`

integral computes the area under the curve per unit of time of subsequent non-null records.

#### Function type signature

```flux
(
    <-tables: stream[A],
    ?column: string,
    ?interpolate: string,
    ?timeColumn: string,
    ?unit: duration,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| unit | unit: Unit of time to use to compute the integral. | No |
| column | column: Column to operate on. Default is `_value`. | No |
| timeColumn | timeColumn: Column that contains time values to use in the operation.
  Default is `_time`. | No |
| interpolate | interpolate: Type of interpolation to use. Default is `""`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `join()`

join merges two streams of tables into a single output stream based on columns with equal values.
Null values are not considered equal when comparing column values.
The resulting schema is the union of the input schemas.
The resulting group key is the union of the input group keys.

#### Function type signature

```flux
(<-tables: A, ?method: string, ?on: [string]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Record containing two input streams to join. | No |
| on | on: List of columns to join on. | No |
| method | method: Join method. Default is `inner`. | No |
### `kaufmansAMA()`

kaufmansAMA calculates the Kaufman’s Adaptive Moving Average (KAMA) using
values in input tables.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Period or number of points to use in the calculation. | Yes |
| column | column: Column to operate on. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `kaufmansER()`

kaufmansER computes the Kaufman's Efficiency Ratio (KER) of values in the
`_value` column for each input table.

#### Function type signature

```flux
(<-tables: stream[A], n: int) => stream[{B with _value: float, _value: float}] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Period or number of points to use in the calculation. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `keep()`

keep returns a stream of tables containing only the specified columns.

#### Function type signature

```flux
(<-tables: stream[A], ?columns: [string], ?fn: (column: string) => bool) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| columns | columns: Columns to keep in output tables. Cannot be used with `fn`. | No |
| fn | fn: Predicate function that takes a column name as a parameter (column) and
  returns a boolean indicating whether or not the column should be kept in
  output tables. Cannot be used with `columns`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `keyValues()`

keyValues returns a stream of tables with each input tables' group key and
two columns, _key and _value, that correspond to unique column label and value
pairs for each input table.

#### Function type signature

```flux
(<-tables: stream[A], ?keyColumns: [string]) => stream[{B with _value: C, _key: string}] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| keyColumns | keyColumns: List of columns from which values are extracted. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `keys()`

keys returns the columns that are in the group key of each input table.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to store group key labels in. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `last()`

last returns the last row with a non-null value from each input table.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to use to verify the existence of a value.
  Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `length()`

length returns the number of elements in an array.

#### Function type signature

```flux
(<-arr: [A]) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array to evaluate. Default is the piped-forward array (`<-`). | No |
### `limit()`

limit returns the first `n` rows after the specified `offset` from each input table.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?offset: int) => stream[A]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Maximum number of rows to return. | Yes |
| offset | offset: Number of rows to skip per table before limiting to `n`.
  Default is `0`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `linearBins()`

linearBins generates a list of linearly separated float values.

#### Function type signature

```flux
(count: int, start: float, width: float, ?infinity: bool) => [float]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| start | start: First value to return in the list. | Yes |
| width | width: Distance between subsequent values. | Yes |
| count | count: Number of values to return. | Yes |
| infinity | infinity: Include an infinite value at the end of the list. Default is `true`. | No |
### `logarithmicBins()`

logarithmicBins generates a list of exponentially separated float values.

#### Function type signature

```flux
(count: int, factor: float, start: float, ?infinity: bool) => [float]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| start | start: First value to return in the list. | Yes |
| factor | factor: Multiplier to apply to subsequent values. | Yes |
| count | count: Number of values to return. | Yes |
| infinity | infinity: Include an infinite value at the end of the list. Default is `true`. | No |
### `lowestAverage()`

lowestAverage calculates the average of each input table and returns the lowest
`n` averages.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?column: string, ?groupColumns: [string]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of records to return. | Yes |
| column | column: Column to evaluate. Default is `_value`. | No |
| groupColumns | groupColumns: List of columns to group by. Default is `[]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `lowestCurrent()`

lowestCurrent selects the last record from each input table and returns the
lowest `n` records.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?column: string, ?groupColumns: [string]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of records to return. | Yes |
| column | column: Column to evaluate. Default is `_value`. | No |
| groupColumns | groupColumns: List of columns to group by. Default is `[]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `lowestMin()`

lowestMin selects the record with the lowest value in the specified `column`
from each input table and returns the bottom `n` records.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?column: string, ?groupColumns: [string]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of records to return. | Yes |
| column | column: Column to evaluate. Default is `_value`. | No |
| groupColumns | groupColumns: List of columns to group by. Default is `[]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `map()`

map iterates over and applies a function to input rows.

#### Function type signature

```flux
(<-tables: stream[A], fn: (r: A) => B, ?mergeKey: bool) => stream[B]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fn | fn: Single argument function to apply to each record.
  The return value must be a record. | Yes |
| mergeKey | mergeKey: _(Deprecated)_ Merge group keys of mapped records. Default is `false`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `max()`

max returns the row with the maximum value in a specified column from each
input table.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to return maximum values from. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `mean()`

mean returns the average of non-null values in a specified column from each
input table.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to use to compute means. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `median()`

median returns the median `_value` of an input table or all non-null records
in the input table with values that fall within the 0.5 quantile (50th percentile).

#### Function type signature

```flux
(<-tables: stream[A], ?column: string, ?compression: float, ?method: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to use to compute the median. Default is `_value`. | No |
| method | method: Computation method. Default is `estimate_tdigest`. | No |
| compression | compression: Number of centroids to use when compressing the dataset.
  Default is `0.0`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `min()`

min returns the row with the minimum value in a specified column from each
input table.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to return minimum values from. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `mode()`

mode returns the non-null value or values that occur most often in a
specified column in each input table.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[{B with _value: C}] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to return the mode from. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `movingAverage()`

movingAverage calculates the mean of non-null values using the current value
and `n - 1` previous values in the `_values` column.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}], n: int) => stream[{A with _value: float}] where B: Numeric
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of values to average. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `now()`

now is a function option that, by default, returns the current system time.

#### Function type signature

```flux
() => time
```

#### Parameters

No parameters provided.

### `pearsonr()`

pearsonr returns the covariance of two streams of tables normalized to the
Pearson R coefficient.

#### Function type signature

```flux
(on: [string], x: A, y: B) => stream[C] where C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| x | x: First input stream. | Yes |
| y | y: Second input stream. | Yes |
| on | on: List of columns to join on. | Yes |
### `pivot()`

pivot collects unique values stored vertically (column-wise) and aligns them
horizontally (row-wise) into logical sets.

#### Function type signature

```flux
(<-tables: stream[A], columnKey: [string], rowKey: [string], valueColumn: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| rowKey | rowKey: Columns to use to uniquely identify an output row. | Yes |
| columnKey | columnKey: Columns to use to identify new output columns. | Yes |
| valueColumn | valueColumn: Column to use to populate the value of pivoted `columnKey` columns. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `quantile()`

quantile returns rows from each input table with values that fall within a
specified quantile or returns the row with the value that represents the
specified quantile.

#### Function type signature

```flux
(
    <-tables: stream[A],
    q: float,
    ?column: string,
    ?compression: float,
    ?method: string,
) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to use to compute the quantile. Default is `_value`. | No |
| q | q: Quantile to compute. Must be between `0.0` and `1.0`. | Yes |
| method | method: Computation method. Default is `estimate_tdigest`. | No |
| compression | compression: Number of centroids to use when compressing the dataset.
  Default is `1000.0`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `range()`

range filters rows based on time bounds.

#### Function type signature

```flux
(<-tables: stream[{C with _time: time}], start: A, ?stop: B) => stream[{C with _time: time, _stop: time, _start: time}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| start | start: Earliest time to include in results. | Yes |
| stop | stop: Latest time to include in results. Default is `now()`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `reduce()`

reduce aggregates rows in each input table using a reducer function (`fn`).

#### Function type signature

```flux
(<-tables: stream[B], fn: (accumulator: A, r: B) => A, identity: A) => stream[C] where A: Record, B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fn | fn: Reducer function to apply to each row record (`r`). | Yes |
| identity | identity: Record that defines the reducer record and provides initial values
  for the reducer operation on the first row. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `relativeStrengthIndex()`

relativeStrengthIndex measures the relative speed and change of values in input tables.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?columns: [string]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of values to use to calculate the RSI. | Yes |
| columns | columns: Columns to operate on. Default is `["_value"]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `rename()`

rename renames columns in a table.

#### Function type signature

```flux
(<-tables: stream[B], ?columns: A, ?fn: (column: string) => string) => stream[C] where A: Record, B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| columns | columns: Record that maps old column names to new column names. | No |
| fn | fn: Function that takes the current column name (`column`) and returns a
  new column name. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `sample()`

sample selects a subset of the rows from each input table.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?column: string, ?pos: int) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Sample every Nth element. | Yes |
| pos | pos: Position offset from the start of results where sampling begins.
  Default is -1 (random offset). | No |
| column | column: Column to operate on. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `set()`

set assigns a static column value to each row in the input tables.

#### Function type signature

```flux
(<-tables: stream[A], key: string, value: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| key | key: Label of the column to modify or set. | Yes |
| value | value: String value to set. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `skew()`

skew returns the skew of non-null records in each input table as a float.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to operate on. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `sort()`

sort orders rows in each intput table based on values in specified columns.

#### Function type signature

```flux
(<-tables: stream[A], ?columns: [string], ?desc: bool) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| columns | columns: List of columns to sort by. Default is `["_value"]`. | No |
| desc | desc: Sort results in descending order. Default is `false`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `spread()`

spread returns the difference between the minimum and maximum values in a
specified column.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to operate on. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `stateCount()`

stateCount returns the number of consecutive rows in a given state.

#### Function type signature

```flux
(<-tables: stream[A], fn: (r: A) => bool, ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fn | fn: Predicate function that identifies the state of a record. | Yes |
| column | column: Column to store the state count in. Default is `stateCount`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `stateDuration()`

stateDuration returns the cumulative duration of a given state.

#### Function type signature

```flux
(
    <-tables: stream[A],
    fn: (r: A) => bool,
    ?column: string,
    ?timeColumn: string,
    ?unit: duration,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fn | fn: Predicate function that identifies the state of a record. | Yes |
| column | column: Column to store the state duration in. Default is `stateDuration`. | No |
| timeColumn | timeColumn: Time column to use to calculate elapsed time between rows.
  Default is `_time`. | No |
| unit | unit: Unit of time to use to increment state duration. Default is `1s` (seconds). | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `stateTracking()`

stateTracking returns the cumulative count and duration of consecutive
rows that match a predicate function that defines a state.

#### Function type signature

```flux
(
    <-tables: stream[A],
    fn: (r: A) => bool,
    ?countColumn: string,
    ?durationColumn: string,
    ?durationUnit: duration,
    ?timeColumn: string,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fn | fn: Predicate function to determine state. | Yes |
| countColumn | countColumn: Column to store state count in. | No |
| durationColumn | durationColumn: Column to store state duration in. | No |
| durationUnit | durationUnit: Unit of time to report state duration in. Default is `1s`. | No |
| timeColumn | timeColumn: Column with time values used to calculate state duration.
  Default is `_time`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `stddev()`

stddev returns the standard deviation of non-null values in a specified column.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string, ?mode: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to operate on. Default is `_value`. | No |
| mode | mode: Standard deviation mode or type of standard deviation to calculate.
  Default is `sample`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `string()`

string converts a value to a string type.

#### Function type signature

```flux
(v: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to convert. | Yes |
### `sum()`

sum returns the sum of non-null values in a specified column.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to operate on. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `tableFind()`

tableFind extracts the first table in a stream with group key values that
match a specified predicate.

#### Function type signature

```flux
(<-tables: stream[B], fn: (key: A) => bool) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fn | fn: Predicate function to evaluate input table group keys. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `tail()`

tail limits each output table to the last `n` rows.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?offset: int) => stream[A]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Maximum number of rows to output. | Yes |
| offset | offset: Number of records to skip at the end of a table table before
  limiting to `n`. Default is 0. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `time()`

time converts a value to a time type.

#### Function type signature

```flux
(v: A) => time
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to convert. | Yes |
### `timeShift()`

timeShift adds a fixed duration to time columns.

#### Function type signature

```flux
(<-tables: stream[A], duration: duration, ?columns: [string]) => stream[A]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| duration | duration: Amount of time to add to each time value. May be a negative duration. | Yes |
| columns | columns: List of time columns to operate on. Default is `["_start", "_stop", "_time"]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `timeWeightedAvg()`

timeWeightedAvg returns the time-weighted average of non-null values in
`_value` column as a float for each input table.

#### Function type signature

```flux
(<-tables: stream[A], unit: duration) => stream[{B with _value: float, _value: float, _stop: D, _start: C}] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| unit | unit: Unit of time to use to compute the time-weighted average. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `timedMovingAverage()`

timedMovingAverage returns the mean of values in a defined time range at a
specified frequency.

#### Function type signature

```flux
(<-tables: stream[A], every: duration, period: duration, ?column: string) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| every | every: Frequency of time window. | Yes |
| period | period: Length of each averaged time window. | Yes |
| column | column: Column to operate on. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `toBool()`

toBool converts all values in the `_value` column to boolean types.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B, _value: bool}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `toFloat()`

toFloat converts all values in the `_value` column to float types.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B, _value: float}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `toInt()`

toInt converts all values in the `_value` column to integer types.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B, _value: int}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `toString()`

toString converts all values in the `_value` column to string types.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B, _value: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `toTime()`

toTime converts all values in the `_value` column to time types.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B, _value: time}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `toUInt()`

toUInt converts all values in the `_value` column to unsigned integer types.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}]) => stream[{A with _value: B, _value: uint}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `today()`

today returns the now() timestamp truncated to the day unit.

#### Function type signature

```flux
() => time
```

#### Parameters

No parameters provided.

### `top()`

top sorts each input table by specified columns and keeps the top `n` records
in each table.

#### Function type signature

```flux
(<-tables: stream[A], n: int, ?columns: [string]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of rows to return from each input table. | Yes |
| columns | columns: List of columns to sort by. Default is `["_value"]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `tripleEMA()`

tripleEMA returns the triple exponential moving average (TEMA) of values in
the `_value` column.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}], n: int) => stream[C] where B: Numeric, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of points to use in the calculation. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `tripleExponentialDerivative()`

tripleExponentialDerivative returns the triple exponential derivative (TRIX)
values using `n` points.

#### Function type signature

```flux
(<-tables: stream[{A with _value: B}], n: int) => stream[{A with _value: float}] where A: Record, B: Numeric
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of points to use in the calculation. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `truncateTimeColumn()`

truncateTimeColumn truncates all input time values in the `_time` to a
specified unit.

#### Function type signature

```flux
(<-tables: stream[{B with _time: C}], unit: duration, ?timeColumn: A) => stream[{B with _time: C, _time: time}] where C: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| unit | unit: Unit of time to truncate to. | Yes |
| timeColumn | timeColumn: Time column to truncate. Default is `_time`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `uint()`

uint converts a value to an unsigned integer type.

#### Function type signature

```flux
(v: A) => uint
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to convert. | Yes |
### `union()`

union merges two or more input streams into a single output stream.

#### Function type signature

```flux
(tables: [stream[A]]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: List of two or more streams of tables to union together. | Yes |
### `unique()`

unique returns all records containing unique values in a specified column.

#### Function type signature

```flux
(<-tables: stream[A], ?column: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| column | column: Column to search for unique values. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `window()`

window groups records using regular time intervals.

#### Function type signature

```flux
(
    <-tables: stream[A],
    ?createEmpty: bool,
    ?every: duration,
    ?location: {zone: string, offset: duration},
    ?offset: duration,
    ?period: duration,
    ?startColumn: string,
    ?stopColumn: string,
    ?timeColumn: string,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| every | every: Duration of time between windows. | No |
| period | period: Duration of windows. Default is the `every` value. | No |
| offset | offset: Duration to shift the window boundaries by. Default is `0s`. | No |
| location | location: Location used to determine timezone. Default is the `location` option. | No |
| timeColumn | timeColumn: Column that contains time values. Default is `_time`. | No |
| startColumn | startColumn: Column to store the window start time in. Default is `_start`. | No |
| stopColumn | stopColumn: Column to store the window stop time in. Default is `_stop`. | No |
| createEmpty | createEmpty: Create empty tables for empty window. Default is `false`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `yield()`

yield delivers input data as a result of the query.

#### Function type signature

```flux
(<-tables: stream[A], ?name: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| name | name: Unique name for the yielded results. Default is `_results`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
