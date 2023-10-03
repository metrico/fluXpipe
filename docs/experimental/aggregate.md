## `aggregate` package

Package aggregate provides functions to simplify common aggregate operations.

Import the `aggregate` package:

```flux
import "aggregate"
```

### Functions

### `rate()`

rate calculates the average rate of increase per window of time for each
input table.

#### Function type signature

```flux
(<-tables: stream[A], every: duration, ?groupColumns: [string], ?unit: duration) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| every | every: Duration of time windows. | Yes |
| groupColumns | groupColumns: List of columns to group by. Default is `[]`. | No |
| unit | unit: Time duration to use when calculating the rate. Default is `1s`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
