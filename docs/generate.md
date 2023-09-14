## `generate` package

Package generate provides functions for generating data.

Import the `generate` package:

```flux
import "generate"
```

### Functions

### `from()`

from generates data using the provided parameter values.

#### Function type signature

```flux
(count: int, fn: (n: int) => int, start: A, stop: A) => stream[{_value: int, _time: time, _stop: time, _start: time}] where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| count | count: Number of rows to generate. | Yes |
| fn | fn: Function used to generate values. | Yes |
| start | start: Beginning of the time range to generate values in. | Yes |
| stop | stop: End of the time range to generate values in. | Yes |
