import "sql" 

sql.from(
  driverName: "clickhouse",
  dataSourceName: "clickhouse://default:@clickhouse-host:9000/system",
  query: "SELECT database, total_rows FROM tables WHERE total_rows > 0"
) 
|> rename(columns: {database: "_value", total_rows: "_data"})
|> keep(columns: ["_value","_data"])
