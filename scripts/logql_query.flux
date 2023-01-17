import "contrib/qxip/logql"

option logql.defaultURL = "https://qryn:3100"

logql.query_range(
     query: "rate({type=\"syslog\"}[1m])",
     start: -10m,
     end: now(),
)
|> map(fn: (r) => ({r with timestamp_ns: time(v: uint(v: r.timestamp_ns)) }))
|> sort(columns: ["timestamp_ns"])
