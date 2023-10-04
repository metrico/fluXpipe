## `opsgenie` package

Package opsgenie provides functions that send alerts to
[Atlassian Opsgenie](https://www.atlassian.com/software/opsgenie)
using the [Opsgenie v2 API](https://docs.opsgenie.com/docs/alert-api#create-alert).

Import the `opsgenie` package:

```flux
import "opsgenie"
```

### Functions

### `endpoint()`

endpoint sends an alert message to Opsgenie using data from table rows.

#### Function type signature

```flux
(
    apiKey: string,
    ?entity: string,
    ?url: string,
) => (
    mapFn: (
        r: A,
    ) => {
        B with
        visibleTo: [string],
        tags: E,
        responders: [string],
        priority: string,
        message: string,
        details: D,
        description: string,
        alias: string,
        actions: C,
    },
) => (<-tables: stream[A]) => stream[{A with _sent: string}] where D: Stringable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Opsgenie API URL. Defaults to `https://api.opsgenie.com/v2/alerts`. | No |
| apiKey | apiKey: (Required) Opsgenie API authorization key. | Yes |
| entity | entity: Alert entity used to specify the alert domain. | No |
### `respondersToJSON()`

respondersToJSON converts an array of Opsgenie responder strings
to a string-encoded JSON array that can be embedded in an alert message.

#### Function type signature

```flux
(v: [string]) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: (Required) Array of Opsgenie responder strings.
  Responder strings must begin with
  `user: `, `team: `, `escalation: `, or `schedule: `. | Yes |
### `sendAlert()`

sendAlert sends an alert message to Opsgenie.

#### Function type signature

```flux
(
    apiKey: string,
    message: string,
    ?actions: A,
    ?alias: string,
    ?description: string,
    ?details: B,
    ?entity: string,
    ?priority: string,
    ?responders: [string],
    ?tags: C,
    ?url: string,
    ?visibleTo: [string],
) => int where B: Stringable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Opsgenie API URL. Defaults to `https://api.opsgenie.com/v2/alerts`. | No |
| apiKey | apiKey: (Required) Opsgenie API authorization key. | Yes |
| message | message: (Required) Alert message text.
  130 characters or less. | Yes |
| alias | alias: Opsgenie alias usee to de-deduplicate alerts.
  250 characters or less.
  Defaults to [message](https://docs.influxdata.com/flux/v0.x/stdlib/contrib/sranka/opsgenie/sendalert/#message). | No |
| description | description: Alert description. 15000 characters or less. | No |
| priority | priority: Opsgenie alert priority. | No |
| responders | responders: List of responder teams or users.
  Use the `user: ` prefix for users and `teams: ` prefix for teams. | No |
| tags | tags: Alert tags. | No |
| entity | entity: Alert entity used to specify the alert domain. | No |
| actions | actions: List of actions available for the alert. | No |
| details | details: Additional alert details. Must be a JSON-encoded map of key-value string pairs. | No |
| visibleTo | visibleTo: List of teams and users the alert will be visible to without sending notifications.
  Use the `user: ` prefix for users and `teams: ` prefix for teams. | No |
