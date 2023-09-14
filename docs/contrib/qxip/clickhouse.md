## `clickhouse` package

Package clickhouse provides functions to query [ClickHouse](https://clickhouse.com/) using the ClickHouse HTTP API.

Import the `clickhouse` package:

```flux
import "clickhouse"
```

### Functions

### `defaultURL()`

defaultURL is the default ClickHouse HTTP API URL.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `query()`

query queries data from ClickHouse using specified parameters.

#### Function type signature

```flux
(
    query: string,
    ?cors: string,
    ?format: string,
    ?limit: A,
    ?max_bytes: B,
    ?url: string,
) => stream[C] where A: Stringable, B: Stringable, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: ClickHouse HTTP API URL. Default is `http://127.0.0.1:8123`. | No |
| query | query: ClickHouse query to execute. | Yes |
| limit | limit: Query rows limit. Defaults is `100`. | No |
| cors | cors: Request remote CORS headers. Defaults is `1`. | No |
| max_bytes | max_bytes: Query bytes limit. Default is `10000000`. | No |
| format | format: Query format. Default is `CSVWithNames`. | No |
