## `telegram` package

Package telegram provides functions for sending messages to [Telegram](https://telegram.org/)
using the [Telegram Bot API](https://core.telegram.org/bots/api).

Import the `telegram` package:

```flux
import "telegram"
```

### Functions

### `defaultDisableWebPagePreview()`

defaultDisableWebPagePreview - Use Telegram web page preview by default. Default is `false`.

#### Function type signature

```flux
bool
```

#### Parameters

No parameters provided.

### `defaultParseMode()`

defaultParseMode is the default [Telegram parse mode](https://core.telegram.org/bots/api#formatting-options). Default is `MarkdownV2`.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `defaultSilent()`

defaultSilent - Send silent Telegram notifications by default. Default is `true`.

#### Function type signature

```flux
bool
```

#### Parameters

No parameters provided.

### `defaultURL()`

defaultURL is the default Telegram bot URL. Default is `https://api.telegram.org/bot`.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `endpoint()`

endpoint sends a message to a Telegram channel using data from table rows.

#### Function type signature

```flux
(
    token: string,
    ?disableWebPagePreview: A,
    ?parseMode: B,
    ?url: string,
) => (
    mapFn: (r: C) => {D with text: G, silent: F, channel: E},
) => (<-tables: stream[C]) => stream[{C with _sent: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL of the Telegram bot endpoint. Default is `https://api.telegram.org/bot`. | No |
| token | token: Telegram bot token. | Yes |
| parseMode | parseMode: [Parse mode](https://core.telegram.org/bots/api#formatting-options)
  of the message text.
  Default is `MarkdownV2`. | No |
| disableWebPagePreview | disableWebPagePreview: Disable preview of web links in the sent message.
  Default is false. | No |
### `message()`

message sends a single message to a Telegram channel
using the [`sendMessage`](https://core.telegram.org/bots/api#sendmessage) method of the Telegram Bot API.

#### Function type signature

```flux
(
    channel: A,
    text: B,
    token: string,
    ?disableWebPagePreview: C,
    ?parseMode: D,
    ?silent: E,
    ?url: string,
) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL of the Telegram bot endpoint. Default is `https://api.telegram.org/bot`. | No |
| token | token: Telegram bot token. | Yes |
| channel | channel: Telegram channel ID. | Yes |
| text | text: Message text. | Yes |
| parseMode | parseMode: [Parse mode](https://core.telegram.org/bots/api#formatting-options)
  of the message text.
  Default is `MarkdownV2`. | No |
| disableWebPagePreview | disableWebPagePreview: Disable preview of web links in the sent message.
  Default is `false`. | No |
| silent | silent: Send message [silently](https://telegram.org/blog/channels-2-0#silent-messages).
  Default is `true`. | No |
