## `testing` package

Package testing provides functions for testing Flux operations.

Import the `testing` package:

```flux
import "testing"
```

### Functions

### `assertEmpty()`

assertEmpty tests if an input stream is empty. If not empty, the function returns an error.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[A]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `assertEqualValues()`

assertEqualValues tests whether two values are equal.

#### Function type signature

```flux
(got: A, want: A) => stream[{v: A, _diff: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| got | got: Value to test. | Yes |
| want | want: Expected value to test against. | Yes |
### `assertEquals()`

assertEquals tests whether two streams of tables are identical.

#### Function type signature

```flux
(<-got: stream[A], name: string, want: stream[A]) => stream[A]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| name | name: Unique assertion name. | Yes |
| got | got: Data to test. Default is piped-forward data (`<-`). | No |
| want | want: Expected data to test against. | Yes |
### `diff()`

diff produces a diff between two streams.

#### Function type signature

```flux
(
    <-got: stream[A],
    want: stream[A],
    ?epsilon: B,
    ?nansEqual: C,
    ?verbose: D,
) => stream[{A with _diff: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| got | got: Stream containing data to test. Default is piped-forward data (`<-`). | No |
| want | want: Stream that contains data to test against. | Yes |
| epsilon | epsilon: Specify how far apart two float values can be, but still considered equal. Defaults to 0.000000001. | No |
| verbose | verbose: Include detailed differences in output. Default is `false`. | No |
| nansEqual | nansEqual: Consider `NaN` float values equal. Default is `false`. | No |
### `load()`

load loads test data from a stream of tables.

#### Function type signature

```flux
(<-tables: A) => A
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `shouldError()`

shouldError calls a function that catches any error and checks that the error matches the expected value.

#### Function type signature

```flux
(fn: () => A, want: regexp) => stream[{v: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fn | fn: Function to call. | Yes |
| want | want: Regular expression to match the expected error. | Yes |
### `tags()`

tags is a list of tags that will be applied to a test case.

#### Function type signature

```flux
[A]
```

#### Parameters

No parameters provided.

