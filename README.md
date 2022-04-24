# <img src="https://user-images.githubusercontent.com/1423657/162720189-976cc0cc-7511-4278-a942-9c4e7cc9148a.png" width=350 />

**FluxPipe** is a *very experimental* stand-alone **Flux API** for *ClickHouse* and other *embedded datasources*

> [Flux](https://github.com/influxdata/flux) is a lightweight *scripting language* for querying databases and working with data. [^1]

<br>

### Instructions
Download a [binary release](https://github.com/lmangani/fluxpipe/releases/) or build from source


#### 📦 Download Binary
```bash
curl -fsSL github.com/lmangani/fluxpipe/releases/latest/download/fluxpipe-server -O && chmod +x fluxpipe-server
```
#### 🔌 Start Server w/ Options
```bash
./fluxpipe-server -port 8086
```

Run with `-h` for a full list of parameters

<br>

### 🐛 Usage

💡 _Check out the [scripts](scripts) folder for working examples_

<br>

#### HTTP API
Fluxpipe serves a simple REST API loosely compatible with existing flux integrations and clients

##### Grafana Flux [^1]
Usage with native **Grafana InfluxDB/Flux datasource** _(url + organization fields are required!)_

###### ⭐ ClickHouse SQL
```
import "sql" 

sql.from(
  driverName: "clickhouse",
  dataSourceName: "clickhouse://default:@clickhouse-host:9000/system",
  query: "SELECT database, total_rows FROM tables WHERE total_rows > 0"
) 
|> rename(columns: {database: "_value", total_rows: "_data"})
|> keep(columns: ["_value","_data"])
```
![image](https://user-images.githubusercontent.com/1423657/162625425-15a92f34-562b-4e27-8832-7bc33a90b185.png)

![image](https://user-images.githubusercontent.com/1423657/162428332-77d869a2-d02b-443d-a3ef-3df1fbf899f6.png)

###### ⭐ CURL POST
Usage with curl

```bash
curl -XPOST localhost:8086/api/v2/query -sS \
  -H 'Accept:application/csv' \
  -H 'Content-type:application/vnd.flux' \
  -d 'import g "generate" g.from(start: 2022-04-01T00:00:00Z, stop: 2022-04-01T00:03:00Z, count: 3, fn: (n) => n)'
```
```flux
#datatype,string,long,dateTime:RFC3339,long
#group,false,false,false,false
#default,_result,,,
,result,table,_time,_value
,,0,2022-04-01T00:00:00Z,1
,,0,2022-04-01T00:00:36Z,2
,,0,2022-04-01T00:01:12Z,3
```

------

#### STDIN CMD
Fluxpipe can be used as a command-line tool and stdin pipeline processor

###### Generate CSV
```bash
echo 'import g "generate" g.from(start: 2022-04-01T00:00:00Z, stop: 2022-04-01T00:03:00Z, count: 5, fn: (n) => 1)' \
| ./fluxpipe-server -stdin
```
```csv
#datatype,string,long,dateTime:RFC3339,long
#group,false,false,false,false
#default,_result,,,
,result,table,_time,_value
,,0,2022-04-01T00:00:00Z,1
,,0,2022-04-01T00:00:36Z,1
,,0,2022-04-01T00:01:12Z,1
,,0,2022-04-01T00:01:48Z,1
,,0,2022-04-01T00:02:24Z,1
```
##### Parse CSV
```bash
cat scripts/csv.flux | ./fluxpipe-server -stdin
```
##### Query SQL
```bash
cat scripts/sql.flux | ./fluxpipe-server -stdin
```

#### Status
- [x] Fluxlib
  - [x] parser
  - [x] executor
- [x] ClickHouse driver by @adubovikov
  - [ ] driver tests
- [x] STDIN pipeline
- [x] HTTP api
  - [x] plaintext
  - [x] json support
  - [ ] ENV secrets
  - [ ] api doc





[^1]: Project is not affiliated or endorsed by Influxdata or Grafana Labs. All rights belong to their respective owners.
