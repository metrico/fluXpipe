## `gen` package

Package gen provides methods for generating data.

Import the `gen` package:

```flux
import "gen"
```

### Functions

### `tables()`

tables generates a stream of table data.

#### Function type signature

```flux
(
    n: int,
    ?nulls: float,
    ?seed: int,
    ?tags: [{name: string, cardinality: int}],
) => stream[{A with _value: float, _time: time}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| n | n: Number of rows to generate. | Yes |
| nulls | nulls: Percentage chance that a null value will be used in the input. Valid value range is `[0.0 - 1.0]`. | No |
| tags | tags: Set of tags with their cardinality to generate. | No |
| seed | seed: Pass seed to tables generator to get the very same sequence each time. | No |
