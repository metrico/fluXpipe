## `socket` package

Package socket provides tools for returning data from socket connections.

Import the `socket` package:

```flux
import "socket"
```

### Functions

### `from()`

from returns data from a socket connection and outputs a stream of tables
given a specified decoder.

#### Function type signature

```flux
(url: string, ?decoder: string) => stream[A]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL to return data from. | Yes |
| decoder | decoder: Decoder to use to parse returned data into a stream of tables. | No |
