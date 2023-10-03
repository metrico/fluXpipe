## `dynamic` package

Package dynamic provides tools for working with values of unknown types.

Import the `dynamic` package:

```flux
import "dynamic"
```

### Functions

### `asArray()`

asArray converts a dynamic value into an array of dynamic elements.

#### Function type signature

```flux
(<-v: dynamic) => [dynamic]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Dynamic value to convert. Default is the piped-forward value (`<-`). | No |
### `dynamic()`

dynamic wraps a value so it can be used as a `dynamic` value.

#### Function type signature

```flux
(v: A) => dynamic
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to wrap as dynamic. | Yes |
### `isType()`

isType tests if a dynamic type holds a value of a specified type.

#### Function type signature

```flux
(type: string, v: dynamic) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to test. | Yes |
| type | type: String describing the type to check against. | Yes |
### `jsonEncode()`

jsonEncode converts a dynamic value into JSON bytes.

#### Function type signature

```flux
(v: dynamic) => bytes
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: Value to encode into JSON. | Yes |
### `jsonParse()`

jsonParse takes JSON data as bytes and returns dynamic values.

#### Function type signature

```flux
(data: bytes) => dynamic
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| data | data: JSON data (as bytes) to parse. | Yes |
