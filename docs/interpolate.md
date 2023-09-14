## `interpolate` package

Package interpolate provides functions that insert rows for missing data
at regular intervals and estimate values using different interpolation methods.

Import the `interpolate` package:

```flux
import "interpolate"
```

### Functions

### `linear()`

linear inserts rows at regular intervals using linear interpolation to
determine values for inserted rows.

#### Function type signature

```flux
(<-tables: stream[{A with _value: float, _time: time}], every: duration) => stream[{A with _value: float, _time: time}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| every | every: Duration of time between interpolated points. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
