## `mqtt` package

Package mqtt provides tools for working with Message Queuing Telemetry Transport (MQTT) protocol.

Import the `mqtt` package:

```flux
import "mqtt"
```

### Functions

### `publish()`

publish sends data to an MQTT broker using MQTT protocol.

#### Function type signature

```flux
(
    broker: string,
    message: string,
    topic: string,
    ?clientid: string,
    ?password: string,
    ?qos: int,
    ?retain: bool,
    ?timeout: duration,
    ?username: string,
) => bool
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| broker | broker: MQTT broker connection string. | Yes |
| topic | topic: MQTT topic to send data to. | Yes |
| message | message: Message to send to the MQTT broker. | Yes |
| qos | qos: MQTT Quality of Service (QoS) level. Values range from `[0-2]`.
  Default is `0`. | No |
| retain | retain: MQTT retain flag. Default is `false`. | No |
| clientid | clientid: MQTT client ID. | No |
| username | username: Username to send to the MQTT broker. | No |
| password | password: Password to send to the MQTT broker. | No |
| timeout | timeout: MQTT connection timeout. Default is `1s`. | No |
### `to()`

to outputs data from a stream of tables to an MQTT broker using MQTT protocol.

#### Function type signature

```flux
(
    <-tables: stream[A],
    broker: string,
    ?clientid: string,
    ?name: string,
    ?password: string,
    ?qos: int,
    ?retain: bool,
    ?tagColumns: [string],
    ?timeColumn: string,
    ?timeout: duration,
    ?topic: string,
    ?username: string,
    ?valueColumns: [string],
) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| broker | broker: MQTT broker connection string. | Yes |
| topic | topic: MQTT topic to send data to. | No |
| qos | qos: MQTT Quality of Service (QoS) level. Values range from `[0-2]`. Default is `0`. | No |
| retain | retain: MQTT retain flag. Default is `false`. | No |
| clientid | clientid: MQTT client ID. | No |
| username | username: Username to send to the MQTT broker. | No |
| password | password: Password to send to the MQTT broker.
  Password is only required if the broker requires authentication.
  If you provide a password, you must provide a username. | No |
| name | name: Name for the MQTT message. | No |
| timeout | timeout: MQTT connection timeout. Default is `1s`. | No |
| timeColumn | timeColumn: Column to use as time values in the output line protocol.
  Default is `"_time"`. | No |
| tagColumns | tagColumns: Columns to use as tag sets in the output line protocol.
  Default is `[]`. | No |
| valueColumns | valueColumns: Columns to use as field values in the output line protocol.
  Default is `["_value"]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
