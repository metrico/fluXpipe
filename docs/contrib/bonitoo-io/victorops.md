## `victorops` package

Package victorops provides functions that send events to [VictorOps](https://victorops.com/).

Import the `victorops` package:

```flux
import "victorops"
```

### Functions

### `alert()`

alert sends an alert to VictorOps.

#### Function type signature

```flux
(
    messageType: A,
    url: string,
    ?entityDisplayName: B,
    ?entityID: C,
    ?monitoringTool: D,
    ?stateMessage: E,
    ?timestamp: F,
) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: VictorOps REST endpoint integration URL. | Yes |
| monitoringTool | monitoringTool: Monitoring agent name. Default is `""`. | No |
| messageType | messageType: VictorOps message type (alert behavior). | Yes |
| entityID | entityID: Incident ID. Default is `""`. | No |
| entityDisplayName | entityDisplayName: Incident display name or summary. Default is `""`. | No |
| stateMessage | stateMessage: Verbose incident message. Default is `""`. | No |
| timestamp | timestamp: Incident start time. Default is `now()`. | No |
### `endpoint()`

endpoint sends events to VictorOps using data from input rows.

#### Function type signature

```flux
(
    url: string,
    ?monitoringTool: A,
) => (
    mapFn: (
        r: B,
    ) => {
        C with
        timestamp: H,
        stateMessage: G,
        messageType: F,
        entityID: E,
        entityDisplayName: D,
    },
) => (<-tables: stream[B]) => stream[{B with _sent: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: VictorOps REST endpoint integration URL. | Yes |
| monitoringTool | monitoringTool: Tool to use for monitoring.
  Default is `InfluxDB`. | No |
