## `types` package

Package types provides functions for working with Flux's types.

Import the `types` package:

```flux
import "types"
```

### Functions

### `isNumeric()`

isNumeric tests if a value is a numeric type (int, uint, or float).

#### Function type signature

```flux
(v: A) => bool where A: Basic
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to test. | Yes |
### `isType()`

isType tests if a value is a specified type.

#### Function type signature

```flux
(type: string, v: A) => bool where A: Basic
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to test. | Yes |
| type | type: String describing the type to check against. | Yes |
