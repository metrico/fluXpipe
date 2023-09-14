## `statsmodels` package

Package statsmodels provides functions for calculating statistical models.

Import the `statsmodels` package:

```flux
import "statsmodels"
```

### Functions

### `linearRegression()`

linearRegression performs a linear regression.

#### Function type signature

```flux
(
    <-tables: stream[A],
) => stream[{
    B with
    y_hat: float,
    y: float,
    x: float,
    sy: H,
    sxy: G,
    sxx: F,
    sx: E,
    slope: D,
    errors: float,
    N: C,
}] where A: Record, D: Divisible + Subtractable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
