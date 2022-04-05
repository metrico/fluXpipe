import "sql" 

sql.from(driverName: "mysql",dataSourceName: "default:password@tcp(127.0.0.1:9004)/system", query: "SELECT database, total_rows FROM tables") 
|> rename(columns: {database: "_value", total_rows: "_data"}) |> keep(columns: ["_value","_data"])
