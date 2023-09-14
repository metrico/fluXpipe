## `servicenow` package

Package servicenow  provides functions for sending events to [ServiceNow](https://www.servicenow.com/).

Import the `servicenow` package:

```flux
import "servicenow"
```

### Functions

### `endpoint()`

endpoint sends events to [ServiceNow](https://servicenow.com/) using data from input rows.

#### Function type signature

```flux
(
    password: string,
    url: string,
    username: string,
    ?source: A,
) => (
    mapFn: (
        r: B,
    ) => {
        C with
        severity: J,
        resource: I,
        node: H,
        metricType: G,
        metricName: F,
        messageKey: E,
        description: D,
    },
) => (<-tables: stream[B]) => stream[{B with _sent: string}] where J: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: ServiceNow web service URL. | Yes |
| username | username: ServiceNow username to use for HTTP BASIC authentication. | Yes |
| password | password: ServiceNow password to use for HTTP BASIC authentication. | Yes |
| source | source: Source name. Default is `"Flux"`. | No |
### `event()`

event sends an event to [ServiceNow](https://servicenow.com/).

#### Function type signature

```flux
(
    description: A,
    password: string,
    severity: B,
    url: string,
    username: string,
    ?additionalInfo: C,
    ?messageKey: D,
    ?metricName: E,
    ?metricType: F,
    ?node: G,
    ?resource: H,
    ?source: I,
) => int where B: Equatable, C: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: ServiceNow web service URL. | Yes |
| username | username: ServiceNow username to use for HTTP BASIC authentication. | Yes |
| password | password: ServiceNow password to use for HTTP BASIC authentication. | Yes |
| description | description: Event description. | Yes |
| severity | severity: Severity of the event. | Yes |
| source | source: Source name. Default is `"Flux"`. | No |
| node | node: Node name or IP address related to the event.
  Default is an empty string (`""`). | No |
| metricType | metricType: Metric type related to the event (for example, `CPU`).
  Default is an empty string (`""`). | No |
| resource | resource: Resource related to the event (for example, `CPU-1`).
  Default is an empty string (`""`). | No |
| metricName | metricName: Metric name related to the event (for example, `usage_idle`).
  Default is an empty string (`""`). | No |
| messageKey | messageKey: Unique identifier of the event (for example, the InfluxDB alert ID).
  Default is an empty string (`""`).
  If an empty string, ServiceNow generates a value. | No |
| additionalInfo | additionalInfo: Additional information to include with the event. | No |
