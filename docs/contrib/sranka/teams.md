## `teams` package

Package teams (Microsoft Teams) provides functions
for sending messages to a [Microsoft Teams](https://www.microsoft.com/microsoft-365/microsoft-teams/group-chat-software)
channel using an [incoming webhook](https://docs.microsoft.com/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook).

Import the `teams` package:

```flux
import "teams"
```

### Functions

### `endpoint()`

endpoint sends a message to a Microsoft Teams channel using data from table rows.

#### Function type signature

```flux
(
    url: string,
) => (
    mapFn: (r: A) => {B with title: C, text: string, summary: string},
) => (<-tables: stream[A]) => stream[{A with _sent: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Incoming webhook URL. | Yes |
### `message()`

message sends a single message to a Microsoft Teams channel using an
[incoming webhook](https://docs.microsoft.com/microsoftteams/platform/webhooks-and-connectors/how-to/add-incoming-webhook).

#### Function type signature

```flux
(text: string, title: A, url: string, ?summary: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: Incoming webhook URL. | Yes |
| title | title: Message card title. | Yes |
| text | text: Message card text. | Yes |
| summary | summary: Message card summary.
  Default is `""`. | No |
### `summaryCutoff()`

summaryCutoff is the limit for message summaries.
Default is `70`.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

