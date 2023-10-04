## `testutil` package

Package testutil provides helper function for writing test cases.

Import the `testutil` package:

```flux
import "testutil"
```

### Functions

### `fail()`

fail causes the current script to fail.

#### Function type signature

```flux
() => bool
```

#### Parameters

No parameters provided.

### `makeAny()`

makeAny constructs any value based on a type description as a string.

#### Function type signature

```flux
(typ: string) => A
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| typ | typ: Description of the type to create. | Yes |
### `makeRecord()`

makeRecord is the identity function, but breaks the type connection from input to output.

#### Function type signature

```flux
(o: A) => B where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| o | o: Record value. | Yes |
### `yield()`

yield is the identity function.

#### Function type signature

```flux
(<-v: A) => A
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Any value. | No |
