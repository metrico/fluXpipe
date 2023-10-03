## `webexteams` package

Package webexteams provides functions that send messages
to [Webex Teams](https://www.webex.com/team-collaboration.html).

Import the `webexteams` package:

```flux
import "webexteams"
```

### Functions

### `endpoint()`

endpoint returns a function that sends a message that includes data from input rows to a Webex room.

#### Function type signature

```flux
(
    token: string,
    ?url: string,
) => (
    mapFn: (r: A) => {B with text: E, roomId: D, markdown: C},
) => (<-tables: stream[A]) => stream[{A with _sent: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Base URL of Webex API endpoint (without a trailing slash).
  Default is `https://webexapis.com`. | No |
| token | token: [Webex API access token](https://developer.webex.com/docs/api/getting-started). | Yes |
### `message()`

message sends a single message to Webex
using the [Webex messages API](https://developer.webex.com/docs/api/v1/messages/create-a-message).

#### Function type signature

```flux
(
    markdown: A,
    roomId: B,
    text: C,
    token: string,
    ?url: string,
) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Base URL of Webex API endpoint (without a trailing slash).
  Default is `https://webexapis.com`. | No |
| token | token: [Webex API access token](https://developer.webex.com/docs/api/getting-started). | Yes |
| roomId | roomId: Room ID to send the message to. | Yes |
| text | text: Plain text message. | Yes |
| markdown | markdown: [Markdown formatted message](https://developer.webex.com/docs/api/basics#formatting-messages). | Yes |
