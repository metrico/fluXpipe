## `influxdb` package

Package influxdb provides tools for working with the InfluxDB API.

Import the `influxdb` package:

```flux
import "influxdb"
```

### Functions

### `api()`

api submits an HTTP request to the specified InfluxDB API path and returns a
record containing the HTTP status code, response headers, and the response body.

#### Function type signature

```flux
(
    method: string,
    path: string,
    ?body: bytes,
    ?headers: [string:string],
    ?host: string,
    ?query: [string:string],
    ?timeout: duration,
    ?token: string,
) => {statusCode: int, headers: [string:string], body: bytes}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| method | method: HTTP request method. | Yes |
| path | path: InfluxDB API path. | Yes |
| host | host: InfluxDB host URL _(Required when executed outside of InfluxDB)_.
  Default is `""`. | No |
| token | token: [InfluxDB API token](https://docs.influxdata.com/influxdb/cloud/security/tokens/)
  _(Required when executed outside of InfluxDB)_.
  Default is `""`. | No |
| headers | headers: HTTP request headers. | No |
| query | query: URL query parameters. | No |
| timeout | timeout: HTTP request timeout. Default is `30s`. | No |
| body | body: HTTP request body as bytes. | No |
