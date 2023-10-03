## `iox` package

Package iox provides additional functions for querying data from InfluxDB IOx.

Import the `iox` package:

```flux
import "iox"
```

### Functions

### `from()`

from retrieves data from an IOx bucket between the `start` and `stop` times.

#### Function type signature

```flux
(
    bucket: A,
    ?columns: B,
    ?host: C,
    ?limit: D,
    ?org: E,
    ?secure: F,
    ?start: G,
    ?stop: H,
    ?table: I,
    ?token: J,
) => stream[K] where A: Stringable, B: Stringable, C: Equatable + Stringable, D: Stringable, E: Equatable + Stringable, F: Stringable, G: Timeable, H: Timeable, I: Stringable, J: Equatable + Stringable, K: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: Name of the IOx bucket to query. | Yes |
| start | start: Earliest time to include in results. | No |
| stop | stop: Latest time to include in results. Default is `now()`. | No |
| host | host: URL of the IOx instance to query. | No |
| org | org: Organization name. | No |
| token | token: [API token](https://docs.influxdata.com/influxdb/latest/security/tokens/). | No |
| table | table: Table used in the FlightSQL query. | No |
| limit | limit: Limit for the FlightSQL query. Default is `1000`. | No |
| columns | columns: Columns selected by the FlightSQL query. Default is `*`. | No |
| secure | secure: Secure connection to IOx instance. Default is `true`. | No |
