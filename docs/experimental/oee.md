## `oee` package

Package oee provides functions for calculating overall equipment effectiveness (OEE).

Import the `oee` package:

```flux
import "oee"
```

### Functions

### `APQ()`

APQ computes availability, performance, quality (APQ) and overall equipment
effectiveness (OEE) in producing parts.

#### Function type signature

```flux
(
    <-tables: stream[D],
    idealCycleTime: A,
    plannedTime: B,
    runningState: C,
) => stream[{
    E with
    runTime: G,
    quality: float,
    performance: float,
    oee: float,
    availability: float,
    _time: F,
    _stop: F,
}] where C: Equatable, D: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| runningState | runningState: State value that represents a running state. | Yes |
| plannedTime | plannedTime: Total time that equipment is expected to produce parts. | Yes |
| idealCycleTime | idealCycleTime: Ideal minimum time to produce one part. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `computeAPQ()`

computeAPQ computes availability, performance, and quality (APQ)
and overall equipment effectiveness (OEE) using two separate input streams:
**production events** and **part events**.

#### Function type signature

```flux
(
    idealCycleTime: A,
    partEvents: stream[B],
    plannedTime: C,
    productionEvents: stream[D],
    runningState: E,
) => stream[{
    F with
    runTime: H,
    quality: float,
    performance: float,
    oee: float,
    availability: float,
    _time: G,
    _stop: G,
}] where B: Record, D: Record, E: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| productionEvents | productionEvents: Production events stream that contains the production
  state or start and stop events. | Yes |
| partEvents | partEvents: Part events that contains the running totals of parts produced and
    parts that do not meet quality standards. | Yes |
| runningState | runningState: State value that represents a running state. | Yes |
| plannedTime | plannedTime: Total time that equipment is expected to produce parts. | Yes |
| idealCycleTime | idealCycleTime: Ideal minimum time to produce one part. | Yes |
