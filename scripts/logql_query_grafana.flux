import "contrib/qxip/logql"

option logql.defaultURL = "http://qryn:3100"

logql.query_range(
     query: "rate({job=\"dummy-server\"}[1m])",
     start: v.timeRangeStart, 
     end: v.timeRangeStop
)
|> map(fn: (r) => ({r with _time: time(v: uint(v: r.timestamp_ns)), _value: float(v: r.value) }))
|> drop(columns: ["timestamp_ns", "value"])
|> sort(columns: ["_time"])
|> group(columns: ["labels"])
