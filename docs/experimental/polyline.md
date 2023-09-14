## `polyline` package

Package polyline provides methods for polyline simplication, an efficient way of downsampling curves while retaining moments of variation throughout the path.

Import the `polyline` package:

```flux
import "polyline"
```

### Functions

### `rdp()`

rdp applies the Ramer Douglas Peucker (RDP) algorithm to input data to downsample curves composed
of line segments into visually indistinguishable curves with fewer points.

#### Function type signature

```flux
(
    <-tables: stream[A],
    ?epsilon: float,
    ?retention: float,
    ?timeColumn: string,
    ?valColumn: string,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| valColumn | valColumn: Column with Y axis values of the given curve. Default is `_value`. | No |
| timeColumn | timeColumn: Column with X axis values of the given curve. Default is `_time`. | No |
| epsilon | epsilon: Maximum tolerance value that determines the amount of compression. | No |
| retention | retention: Percentage of points to retain after downsampling. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
