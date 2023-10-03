## `pushbullet` package

Package pushbullet provides functions for sending data to Pushbullet.

Import the `pushbullet` package:

```flux
import "pushbullet"
```

### Functions

### `defaultURL()`

defaultURL is the default Pushbullet API URL used by functions in the `pushbullet` package.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `endpoint()`

endpoint creates the endpoint for the Pushbullet API and sends a notification of type note.

#### Function type signature

```flux
(
    ?token: A,
    ?url: string,
) => (mapFn: (r: B) => {C with title: E, text: D}) => (<-tables: stream[B]) => stream[{B with _sent: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: PushBullet API endpoint URL. Default is `"https://api.pushbullet.com/v2/pushes"`. | No |
| token | token: Pushbullet API token string.  Default is `""`. | No |
### `pushData()`

pushData sends a push notification to the Pushbullet API.

#### Function type signature

```flux
(data: A, ?token: B, ?url: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL of the PushBullet endpoint. Default is `"https://api.pushbullet.com/v2/pushes"`. | No |
| token | token: API token string.  Default is `""`. | No |
| data | data: Data to send to the endpoint. Data is JSON-encoded and sent to the Pushbullet's endpoint. | Yes |
### `pushNote()`

pushNote sends a push notification of type "note" to the Pushbullet API.

#### Function type signature

```flux
(text: A, title: B, ?token: C, ?url: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL of the PushBullet endpoint. Default is `"https://api.pushbullet.com/v2/pushes"`. | No |
| token | token: API token string.  Defaults to: `""`. | No |
| title | title: Title of the notification. | Yes |
| text | text: Text to display in the notification. | Yes |
