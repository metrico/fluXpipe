import "experimental/mqtt"
import "csv"

response = mqtt.publish(
    broker: "tcp://mqtt-server:1883",
    topic: "example-topic",
    message: "Example message",
    qos: 0,
    timeout: 1s,
)

// Use the returned data in a stream of tables
csvData = "#datatype,string,long,string
#group,false,false,false
#default,,,
,result,table,column
,,0,*
"

csv.from(csv: csvData)
    |> map(
        fn: (r) => ({
            httpStatus: response,
        }),
    )
