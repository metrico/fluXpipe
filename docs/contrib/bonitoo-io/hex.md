## `hex` package

Package hex provides functions that perform hexadecimal conversion
of `int`, `uint` or `bytes` values to and from `string` values.

Import the `hex` package:

```flux
import "hex"
```

### Functions

### `bytes()`

bytes converts a hexadecimal string to bytes.

#### Function type signature

```flux
(v: string) => bytes
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to convert. | Yes |
### `int()`

int converts a hexadecimal string to an integer.

#### Function type signature

```flux
(v: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to convert. | Yes |
### `string()`

string converts a Flux basic type to a hexadecimal string.

#### Function type signature

```flux
(v: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to convert. | Yes |
### `uint()`

uint converts a hexadecimal string to an unsigned integer.

#### Function type signature

```flux
(v: string) => uint
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to convert. | Yes |
