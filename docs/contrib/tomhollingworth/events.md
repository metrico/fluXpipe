## `events` package

Package events provides tools for analyzing event-based data.

Import the `events` package:

```flux
import "events"
```

### Functions

### `duration()`

duration calculates the duration of events.

#### Function type signature

```flux
(
    <-tables: stream[A],
    ?columnName: string,
    ?stop: time,
    ?stopColumn: string,
    ?timeColumn: string,
    ?unit: duration,
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| unit | unit: Duration unit of the calculated state duration.
  Default is `1ns`. | No |
| columnName | columnName: Name of the result column.
  Default is `"duration"`. | No |
| timeColumn | timeColumn: Name of the time column.
  Default is `"_time"`. | No |
| stopColumn | stopColumn: Name of the stop column.
  Default is `"_stop"`. | No |
| stop | stop: The latest time to use when calculating results. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
