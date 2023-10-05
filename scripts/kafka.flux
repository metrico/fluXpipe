import "kafka"
import "sampledata"

sampledata.int()
|> kafka.to(
  brokers: ["broker:9092"],
  topic: "test-topic",
  name: "test-metric-name",
  tagColumns: ["tag"],
)
