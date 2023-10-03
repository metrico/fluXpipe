## `slack` package

Package slack provides functions for sending messages to [Slack](https://slack.com/).

Import the `slack` package:

```flux
import "slack"
```

### Functions

### `defaultURL()`

defaultURL defines the default Slack API URL used by functions in the `slack` package.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `endpoint()`

endpoint returns a function that can be used to send a message to Slack per input row.

#### Function type signature

```flux
(
    ?token: string,
    ?url: string,
) => (
    mapFn: (r: A) => {B with text: D, color: string, channel: C},
) => (<-tables: stream[A]) => stream[{A with _sent: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Slack API URL. Default is  `https://slack.com/api/chat.postMessage`. | No |
| token | token: Slack API token. Default is `""`. | No |
### `message()`

message sends a single message to a Slack channel and returns the HTTP
response code of the request.

#### Function type signature

```flux
(
    channel: A,
    color: string,
    text: B,
    ?token: string,
    ?url: string,
) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Slack API URL.
  Default is `https://slack.com/api/chat.postMessage`. | No |
| token | token: Slack API token. Default is `""`. | No |
| channel | channel: Slack channel or user to send the message to. | Yes |
| text | text: Message text. | Yes |
| color | color: Slack message color. | Yes |
### `validateColorString()`

validateColorString ensures a string contains a valid hex color code.

#### Function type signature

```flux
(color: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| color | color: Hex color code. | Yes |
