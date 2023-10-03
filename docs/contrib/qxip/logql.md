## `logql` package

Package logql provides functions for using [LogQL](https://grafana.com/docs/loki/latest/logql/) to query a [Loki](https://grafana.com/oss/loki/) data source.

Import the `logql` package:

```flux
import "logql"
```

### Functions

### `defaultAPI()`

defaultAPI is the default LogQL Query Range API Path.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `defaultURL()`

defaultURL is the default LogQL HTTP API URL.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `query_range()`

query_range queries data from a specified LogQL query within given time bounds,
filters data by query, timerange, and optional limit expressions.
All values are returned as string values (using `raw` mode in `csv.from`)

#### Function type signature

```flux
(
    query: string,
    ?end: A,
    ?limit: B,
    ?orgid: string,
    ?path: string,
    ?start: C,
    ?step: D,
    ?url: string,
) => stream[E] where A: Timeable, B: Stringable, C: Timeable, D: Stringable, E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: LogQL/qryn URL and port. Default is `http://qryn:3100`. | No |
| path | path: LogQL query_range API path. | No |
| limit | limit: Query limit. Default is 100. | No |
| query | query: LogQL query to execute. | Yes |
| start | start: Earliest time to include in results. Default is `-1h`. | No |
| end | end: Latest time to include in results. Default is `now()`. | No |
| step | step: Query resolution step width in seconds. Default is 10. | No |
| orgid | orgid: Optional Loki organization ID for partitioning. Default is `""`. | No |
