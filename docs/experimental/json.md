## `json` package

Package json provides tools for working with JSON.

Import the `json` package:

```flux
import "json"
```

### Functions

### `parse()`

parse takes JSON data as bytes and returns a value.

#### Function type signature

```flux
(data: bytes) => A
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| data | data: JSON data (as bytes) to parse. | Yes |
