## `pagerduty` package

Package pagerduty provides functions for sending data to PagerDuty.

Import the `pagerduty` package:

```flux
import "pagerduty"
```

### Functions

### `actionFromLevel()`

actionFromLevel converts a monitoring level to a PagerDuty action.

#### Function type signature

```flux
(level: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| level | level: Monitoring level to convert to a PagerDuty action. | Yes |
### `actionFromSeverity()`

actionFromSeverity converts a severity to a PagerDuty action.

#### Function type signature

```flux
(severity: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| severity | severity: Severity to convert to a PagerDuty action. | Yes |
### `dedupKey()`

dedupKey uses the group key of an input table to generate and store a
deduplication key in the `_pagerdutyDedupKey`column.
The function sorts, newline-concatenates, SHA256-hashes, and hex-encodes the
group key to create a unique deduplication key for each input table.

#### Function type signature

```flux
(<-tables: stream[A], ?exclude: [string]) => stream[{A with _pagerdutyDedupKey: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| exclude | exclude: Group key columns to exclude when generating the deduplication key.
  Default is ["_start", "_stop", "_level"]. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `defaultURL()`

defaultURL is the default PagerDuty URL used by functions in the `pagerduty` package.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `endpoint()`

endpoint returns a function that sends a message to PagerDuty that includes output data.

#### Function type signature

```flux
(
    ?url: string,
) => (
    mapFn: (
        r: {A with _pagerdutyDedupKey: string},
    ) => {
        B with
        timestamp: K,
        summary: string,
        source: J,
        severity: I,
        routingKey: H,
        group: G,
        eventAction: F,
        clientURL: E,
        client: D,
        class: C,
    },
) => (<-tables: stream[A]) => stream[{A with _status: string, _sent: string, _pagerdutyDedupKey: string, _body: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: PagerDuty v2 Events API URL. | No |
### `sendEvent()`

sendEvent sends an event to PagerDuty and returns the HTTP response code of the request.

#### Function type signature

```flux
(
    class: A,
    client: B,
    clientURL: C,
    dedupKey: D,
    eventAction: E,
    group: F,
    routingKey: G,
    severity: H,
    source: I,
    summary: string,
    timestamp: J,
    ?component: K,
    ?customDetails: L,
    ?pagerdutyURL: string,
) => int where L: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| pagerdutyURL | pagerdutyURL: PagerDuty endpoint URL. | No |
| routingKey | routingKey: Routing key generated from your PagerDuty integration. | Yes |
| client | client: Name of the client sending the alert. | Yes |
| clientURL | clientURL: URL of the client sending the alert. | Yes |
| dedupKey | dedupKey: Per-alert ID that acts as deduplication key and allows you to
  acknowledge or change the severity of previous messages.
  Supports a maximum of 255 characters. | Yes |
| class | class: Class or type of the event. | Yes |
| group | group: Logical grouping used by PagerDuty. | Yes |
| severity | severity: Severity of the event. | Yes |
| eventAction | eventAction: Event type to send to PagerDuty. | Yes |
| source | source: Unique location of the affected system.
  For example, the hostname or fully qualified domain name (FQDN). | Yes |
| component | component: Component responsible for the event. | No |
| summary | summary: Brief text summary of the event used as the summaries or titles of associated alerts.
  The maximum permitted length is 1024 characters. | Yes |
| timestamp | timestamp: Time the detected event occurred in RFC3339nano format. | Yes |
| customDetails | customDetails: Record with additional details about the event. | No |
### `severityFromLevel()`

severityFromLevel converts an InfluxDB status level to a PagerDuty severity.

#### Function type signature

```flux
(level: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| level | level: InfluxDB status level to convert to a PagerDuty severity. | Yes |
