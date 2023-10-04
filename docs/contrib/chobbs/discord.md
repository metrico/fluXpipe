## `discord` package

Package discord provides functions for sending messages to [Discord](https://discord.com/).

Import the `discord` package:

```flux
import "discord"
```

### Functions

### `discordURL()`

discordURL is the Discord webhook URL.
Default is `https://discordapp.com/api/webhooks/`.

#### Function type signature

```flux
string
```

#### Parameters

No parameters provided.

### `endpoint()`

endpoint sends a single message to a Discord channel using a
[Discord webhook](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks&?page=3)
and data from table rows.

#### Function type signature

```flux
(
    username: A,
    webhookID: string,
    webhookToken: string,
    ?avatar_url: B,
) => (mapFn: (r: C) => {D with content: E}) => (<-tables: stream[C]) => stream[{C with _sent: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| webhookToken | webhookToken: Discord [webhook token](https://discord.com/developers/docs/resources/webhook). | Yes |
| webhookID | webhookID: Discord [webhook ID](https://discord.com/developers/docs/resources/webhook). | Yes |
| username | username: Override the Discord webhook’s default username. | Yes |
| avatar_url | avatar_url: Override the Discord webhook’s default avatar. | No |
### `send()`

send sends a single message to a Discord channel using a Discord webhook.

#### Function type signature

```flux
(
    content: A,
    username: B,
    webhookID: string,
    webhookToken: string,
    ?avatar_url: C,
) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| webhookToken | webhookToken: Discord [webhook token](https://discord.com/developers/docs/resources/webhook). | Yes |
| webhookID | webhookID: Discord [webhook ID](https://discord.com/developers/docs/resources/webhook). | Yes |
| username | username: Override the Discord webhook’s default username. | Yes |
| content | content: Message to send to Discord (2000 character limit). | Yes |
| avatar_url | avatar_url: Override the Discord webhook’s default avatar. | No |
