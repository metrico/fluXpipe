## `sensu` package

Package sensu provides functions for sending events to [Sensu Go](https://docs.sensu.io/sensu-go/latest/).

Import the `sensu` package:

```flux
import "sensu"
```

### Functions

### `endpoint()`

endpoint sends an event
to the [Sensu Events API](https://docs.sensu.io/sensu-go/latest/api/events/#create-a-new-event)
using data from table rows.

#### Function type signature

```flux
(
    apiKey: string,
    url: string,
    ?entityName: string,
    ?handlers: A,
    ?namespace: string,
) => (
    mapFn: (r: B) => {C with text: E, status: D, checkName: string},
) => (<-tables: stream[B]) => stream[{B with _sent: string}] where D: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Base URL of [Sensu API](https://docs.sensu.io/sensu-go/latest/migrate/#architecture)
  *without a trailing slash*.
  Example: `http://localhost:8080`. | Yes |
| apiKey | apiKey: Sensu [API Key](https://docs.sensu.io/sensu-go/latest/operations/control-access/). | Yes |
| handlers | handlers: [Sensu handlers](https://docs.sensu.io/sensu-go/latest/reference/handlers/) to execute.
  Default is `[]`. | No |
| namespace | namespace: [Sensu namespace](https://docs.sensu.io/sensu-go/latest/reference/rbac/).
  Default is `default`. | No |
| entityName | entityName: Event source.
  Default is `influxdb`. | No |
### `event()`

event sends a single event to the [Sensu Events API](https://docs.sensu.io/sensu-go/latest/api/events/#create-a-new-event).

#### Function type signature

```flux
(
    apiKey: string,
    checkName: string,
    text: A,
    url: string,
    ?entityName: string,
    ?handlers: B,
    ?namespace: string,
    ?state: string,
    ?status: C,
) => int where C: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Base URL of [Sensu API](https://docs.sensu.io/sensu-go/latest/migrate/#architecture)
  without a trailing slash. | Yes |
| apiKey | apiKey: Sensu [API Key](https://docs.sensu.io/sensu-go/latest/operations/control-access/). | Yes |
| checkName | checkName: Check name. | Yes |
| text | text: Event text. | Yes |
| handlers | handlers: Sensu handlers to execute. Default is `[]`. | No |
| status | status: Event status code that indicates [state](https://docs.influxdata.com/flux/v0.x/stdlib/contrib/sranka/sensu/event/#state).
  Default is `0`. | No |
| state | state: Event state.
  Default is `"passing"` for `0` [status](https://docs.influxdata.com/flux/v0.x/stdlib/contrib/sranka/sensu/event/#status) and `"failing"` for other statuses. | No |
| namespace | namespace: [Sensu namespace](https://docs.sensu.io/sensu-go/latest/reference/rbac/).
  Default is `"default"`. | No |
| entityName | entityName: Event source.
  Default is `influxdb`. | No |
### `toSensuName()`

toSensuName translates a string value to a Sensu name
by replacing non-alphanumeric characters (`[a-zA-Z0-9_.-]`) with underscores (`_`).

#### Function type signature

```flux
(v: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to operate on. | Yes |
