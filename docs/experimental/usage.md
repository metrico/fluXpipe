## `usage` package

Package usage provides tools for collecting usage and usage limit data from
**InfluxDB Cloud**.

Import the `usage` package:

```flux
import "usage"
```

### Functions

### `from()`

from returns usage data from an **InfluxDB Cloud** organization.

#### Function type signature

```flux
(
    start: A,
    stop: B,
    ?host: string,
    ?orgID: string,
    ?raw: C,
    ?token: string,
) => stream[D] where D: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| start | start: Earliest time to include in results. | Yes |
| stop | stop: Latest time to include in results. | Yes |
| host | host: [InfluxDB Cloud region URL](https://docs.influxdata.com/influxdb/cloud/reference/regions/).
  Default is `""`. | No |
| orgID | orgID: InfluxDB Cloud organization ID. Default is `""`. | No |
| token | token: InfluxDB Cloud [API token](https://docs.influxdata.com/influxdb/cloud/security/tokens/).
  Default is `""`. | No |
| raw | raw: Return raw, high resolution usage data instead of downsampled usage data.
  Default is `false`. | No |
### `limits()`

limits returns a record containing usage limits for an **InfluxDB Cloud** organization.

#### Function type signature

```flux
(?host: string, ?orgID: string, ?token: string) => A
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| host | host: [InfluxDB Cloud region URL](https://docs.influxdata.com/influxdb/cloud/reference/regions/).
  Default is `""`. | No |
| orgID | orgID: InfluxDB Cloud organization ID. Default is `""`. | No |
| token | token: InfluxDB Cloud [API token](https://docs.influxdata.com/influxdb/cloud/security/tokens/).
  Default is `""`. | No |
