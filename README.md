<img src="https://user-images.githubusercontent.com/1423657/161867564-4b3fc400-95e5-424c-9210-604d5671a85e.png" width=100 />

# FluxPipe
Experimental [flux](https://github.com/influxdata/flux) pipeline for embedded or external datasources

### Instructions
Download a [binary release](https://github.com/lmangani/fluxpipe/releases/) or build from source


#### 📦 Download Binary
```bash
curl -fsSL github.com/lmangani/fluxpipe/releases/latest/download/fluxpipe -O && chmod +x fluxpipe
```

### 📖 Build
```bash
go build -o fluxpipe -ldflags="-s -w" fluxpipe.go
```

### 🐛 Test
#### CSV
```bash
cat script.flux | fluxpipe
```
#### SQL
```bash
echo 'import "sql" sql.from(driverName: "mysql",dataSourceName: "default:password@tcp(127.0.0.1:9004)/system",query: "SELECT database, total_rows FROM tables") |> rename(columns: {database: "_value", total_rows: "_data"})|> keep(columns: ["_value","_data"])' \
 | fluxpipe
```
