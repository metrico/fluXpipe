## `testing` package

Package testing provides functions for testing Flux operations.

Import the `testing` package:

```flux
import "testing"
```

### Functions

### `assertMatches()`

assertMatches tests whether a string matches a given regex.

#### Function type signature

```flux
(got: string, want: regexp) => stream[{v: string, _diff: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| got | got: Value to test. | Yes |
| want | want: Regex to test against. | Yes |
### `shouldErrorWithCode()`

shouldErrorWithCode calls a function that catches any error and checks that the error matches the expected value.

#### Function type signature

```flux
(code: uint, fn: () => A, want: regexp) => stream[{match: bool, code: uint, _diff: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| fn | fn: Function to call. | Yes |
| want | want: Regular expression to match the expected error. | Yes |
| code | code: Which flux error code to expect | Yes |
