import "sql" 

// Run service with auto-replace of clickhouse-mysql URL details:
//
// # fluxpipe-server -url "default:password@tcp(127.0.0.1:9004)/system"
//

sql.from(driverName: "clickhouse", query: "SELECT database, total_rows FROM tables") 
|> rename(columns: {database: "_value", total_rows: "_data"}) |> keep(columns: ["_value","_data"])
