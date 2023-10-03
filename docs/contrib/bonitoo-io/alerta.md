## `alerta` package

Package alerta provides functions that send alerts to [Alerta](https://alerta.io/).

Import the `alerta` package:

```flux
import "alerta"
```

### Functions

### `alert()`

alert sends an alert to [Alerta](https://alerta.io/).

#### Function type signature

```flux
(
    apiKey: string,
    attributes: A,
    event: B,
    resource: C,
    severity: D,
    url: string,
    ?environment: E,
    ?group: F,
    ?origin: G,
    ?service: H,
    ?tags: I,
    ?text: J,
    ?timestamp: K,
    ?type: L,
    ?value: M,
) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: (Required) Alerta URL. | Yes |
| apiKey | apiKey: (Required) Alerta API key. | Yes |
| resource | resource: (Required) Resource associated with the alert. | Yes |
| event | event: (Required) Event name. | Yes |
| environment | environment: Alerta environment. Valid values: "Production", "Development" or empty string (default). | No |
| severity | severity: (Required) Event severity. See [Alerta severities](https://docs.alerta.io/en/latest/api/alert.html#alert-severities). | Yes |
| service | service: List of affected services. Default is `[]`. | No |
| group | group: Alerta event group. Default is `""`. | No |
| value | value: Event value.  Default is `""`. | No |
| text | text: Alerta text description. Default is `""`. | No |
| tags | tags: List of event tags. Default is `[]`. | No |
| attributes | attributes: (Required) Alert attributes. | Yes |
| origin | origin: monitoring component. | No |
| type | type: Event type. Default is `""`. | No |
| timestamp | timestamp: time alert was generated. Default is `now()`. | No |
### `endpoint()`

endpoint sends alerts to Alerta using data from input rows.

#### Function type signature

```flux
(
    apiKey: string,
    url: string,
    ?environment: A,
    ?origin: B,
) => (
    mapFn: (
        r: C,
    ) => {
        D with
        value: O,
        type: N,
        timestamp: M,
        text: L,
        tags: K,
        severity: J,
        service: I,
        resource: H,
        group: G,
        event: F,
        attributes: E,
    },
) => (<-tables: stream[C]) => stream[{C with _sent: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: (Required) Alerta URL. | Yes |
| apiKey | apiKey: (Required) Alerta API key. | Yes |
| environment | environment: Alert environment. Default is `""`.
  Valid values: "Production", "Development", or empty string (default). | No |
| origin | origin: Alert origin. Default is `"InfluxDB"`. | No |
