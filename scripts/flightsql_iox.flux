import "contrib/qxip/iox"
import "influxdata/influxdb"
iox.from(
     bucket: "qxip",
     host: "eu-central-1-1.aws.cloud2.influxdata.com:443",
     token: "==",
     limit: "10",
     table: "machine_data",
     start: -2d,
 )
 |> range(start: -2d)
 |> aggregateWindow(every: 5s, fn: mean, column: "load", createEmpty: false)
 |> set(key: "_measurement", value: "downsampled")
 |> wideTo(
     bucket: "qxip",
     host: "https://eu-central-1-1.aws.cloud2.influxdata.com",
     token: "==",
     orgID: ""
    )
