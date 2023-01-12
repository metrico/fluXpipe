import "contrib/qxip/logql"

option logql.defaultURL = "http://qryn:3100"

logql.query_range(
     query: "rate({job=\"dummy-server\", method=\"DELETE\"}[5m])",
     start: -1h,
     end: now(),
)
