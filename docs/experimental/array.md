## `array` package

Package array provides functions for manipulating arrays and for building tables from Flux arrays.

Import the `array` package:

```flux
import "array"
```

### Functions

### `concat()`

concat appends two arrays and returns a new array.

#### Function type signature

```flux
(<-arr: [A], v: [A]) => [A]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: First array. Default is the piped-forward array (`<-`). | No |
| v | v: Array to append to the first array. | Yes |
### `filter()`

filter iterates over an array, evaluates each element with a predicate function, and then returns
a new array with only elements that match the predicate.

#### Function type signature

```flux
(<-arr: [A], fn: (x: A) => bool) => [A]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array to filter. Default is the piped-forward array (`<-`). | No |
| fn | fn: Predicate function to evaluate on each element.
  The element is represented by `x` in the predicate function. | Yes |
### `from()`

from constructs a table from an array of records.

#### Function type signature

```flux
(<-rows: [A]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| rows | rows: Array of records to construct a table with. | No |
### `map()`

map iterates over an array, applies a function to each element to produce a new element,
and then returns a new array.

#### Function type signature

```flux
(<-arr: [A], fn: (x: A) => B) => [B]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array to operate on. Defaults is the piped-forward array (`<-`). | No |
| fn | fn: Function to apply to elements. The element is represented by `x` in the function. | Yes |
### `toBool()`

toBool converts all values in an array to booleans.

#### Function type signature

```flux
(<-arr: [A]) => [bool]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array of values to convert. Default is the piped-forward array (`<-`). | No |
### `toDuration()`

toDuration converts all values in an array to durations.

#### Function type signature

```flux
(<-arr: [A]) => [duration]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array of values to convert. Default is the piped-forward array (`<-`). | No |
### `toFloat()`

toFloat converts all values in an array to floats.

#### Function type signature

```flux
(<-arr: [A]) => [float]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array of values to convert. Default is the piped-forward array (`<-`). | No |
### `toInt()`

toInt converts all values in an array to integers.

#### Function type signature

```flux
(<-arr: [A]) => [int]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array of values to convert. Default is the piped-forward array (`<-`). | No |
### `toString()`

toString converts all values in an array to strings.

#### Function type signature

```flux
(<-arr: [A]) => [string]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array of values to convert. Default is the piped-forward array (`<-`). | No |
### `toTime()`

toTime converts all values in an array to times.

#### Function type signature

```flux
(<-arr: [A]) => [time]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array of values to convert. Default is the piped-forward array (`<-`). | No |
### `toUInt()`

toUInt converts all values in an array to unsigned integers.

#### Function type signature

```flux
(<-arr: [A]) => [uint]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| arr | arr: Array of values to convert. Default is the piped-forward array (`<-`). | No |
