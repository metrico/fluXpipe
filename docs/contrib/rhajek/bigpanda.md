## `bigpanda` package

Package bigpanda provides functions for sending alerts to [BigPanda](https://www.bigpanda.io/).

Import the `bigpanda` package:

```flux
import "bigpanda"
```

### Functions

### `defaultTokenPrefix()`

defaultTokenPrefix is the default HTTP authentication scheme to use when authenticating with BigPanda.
Default is `Bearer`.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `defaultUrl()`

defaultUrl is the default [BigPanda alerts API URL](https://docs.bigpanda.io/reference#alerts-how-it-works)
for functions in the `bigpanda` package.
Default is `https://api.bigpanda.io/data/v2/alerts`.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `endpoint()`

endpoint sends alerts to BigPanda using data from input rows.

#### Function type signature

```flux
(
    appKey: A,
    token: string,
    ?url: string,
) => (mapFn: (r: B) => {C with status: D}) => (<-tables: stream[B]) => stream[{B with _sent: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: BigPanda [alerts API URL](https://docs.bigpanda.io/reference#alerts-how-it-works).
  Default is the value of the `bigpanda.defaultURL` option. | No |
| token | token: BigPanda [API Authorization token (API key)](https://docs.bigpanda.io/docs/api-key-management). | Yes |
| appKey | appKey: BigPanda [App Key](https://docs.bigpanda.io/reference#integrating-monitoring-systems). | Yes |
### `sendAlert()`

sendAlert sends an alert to [BigPanda](https://www.bigpanda.io/).

#### Function type signature

```flux
(
    appKey: A,
    rec: B,
    status: C,
    token: string,
    url: string,
) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: BigPanda [alerts API URL](https://docs.bigpanda.io/reference#alerts-how-it-works).
  Default is the value of the `bigpanda.defaultURL` option. | Yes |
| token | token: BigPanda [API Authorization token (API key)](https://docs.bigpanda.io/docs/api-key-management). | Yes |
| appKey | appKey: BigPanda [App Key](https://docs.bigpanda.io/reference#integrating-monitoring-systems). | Yes |
| status | status: BigPanda [alert status](https://docs.bigpanda.io/reference#alerts). | Yes |
| rec | rec: Additional [alert parameters](https://docs.bigpanda.io/reference#alert-object) to send to the BigPanda alert API. | Yes |
### `statusFromLevel()`

statusFromLevel converts an alert level to a BigPanda status.

#### Function type signature

```flux
(level: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| level | level: Alert level. | Yes |
