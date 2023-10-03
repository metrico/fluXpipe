## `anomalydetection` package

Package anomalydetection detects anomalies in time series data.

Import the `anomalydetection` package:

```flux
import "anomalydetection"
```

### Functions

### `mad()`

mad uses the median absolute deviation (MAD) algorithm to detect anomalies in a data set.

#### Function type signature

```flux
(<-table: stream[B], ?threshold: A) => stream[{C with level: string, _value_diff_med: D, _value_diff: D, _value: D}] where A: Comparable + Equatable, B: Record, D: Comparable + Divisible + Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| threshold | threshold: Deviation threshold for anomalies. | No |
| table | table: Input data. Default is piped-forward data (`<-`). | No |
