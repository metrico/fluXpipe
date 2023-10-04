## `kafka` package

Package kafka provides tools for working with [Apache Kafka](https://kafka.apache.org/).

Import the `kafka` package:

```flux
import "kafka"
```

### Functions

### `to()`

to sends data to [Apache Kafka](https://kafka.apache.org/) brokers.

#### Function type signature

```flux
(
    <-tables: stream[A],
    brokers: [string],
    topic: string,
    ?balancer: string,
    ?name: string,
    ?nameColumn: string,
    ?tagColumns: [string],
    ?timeColumn: string,
    ?valueColumns: [string],
) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| brokers | brokers: List of Kafka brokers to send data to. | Yes |
| topic | topic: Kafka topic to send data to. | Yes |
| balancer | balancer: Kafka load balancing strategy. Default is `hash`. | No |
| name | name: Kafka metric name. Default is the value of the `nameColumn`. | No |
| nameColumn | nameColumn: Column to use as the Kafka metric name.
  Default is `_measurement`. | No |
| timeColumn | timeColumn: Time column. Default is `_time`. | No |
| tagColumns | tagColumns: List of tag columns in input data. | No |
| valueColumns | valueColumns: List of value columns in input data. Default is `["_value"]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
