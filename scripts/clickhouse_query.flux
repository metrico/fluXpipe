import "contrib/qxip/clickhouse"

option clickhouse.defaultURL = "https://demo:demo@github.demo.trial.altinity.cloud:8443"

clickhouse.query(query: "SELECT version()")
