# <img src="https://user-images.githubusercontent.com/1423657/231777629-38f9c8f5-7968-4ff3-ae66-2fd2459ffa69.png" width=200 />

[![Build-n-Release](https://github.com/metrico/fluXpipe/actions/workflows/go.yml/badge.svg)](https://github.com/metrico/fluXpipe/actions/workflows/go.yml)

**FluxPipe** is an *experimental* stand-alone **Flux API** for *serverless workers* and *embedded datasources*<br>
Execute your Flux scripts locally, in serverless functions or anywhere else - _decoupled from the data and database._<br>

_Fluxpipe runs at 141,6Km/h_[*](https://en.wikipedia.org/wiki/DeLorean_time_machine#:~:text=Various%20proposals%20have%20been%20brought,speedometer%2C%20modified%20for%20the%20movie.) and is compatible with **InfluxDB 3.0 / IOx, ClickHouse, Grafana** and _beyond!_

<br>

> InfluxDB [Flux](https://github.com/InfluxCommunity/flux) is a lightweight *scripting language* for querying databases and working with data. [^1]<br>
> Need a Flux introduction? Check out the [official page](https://www.influxdata.com/products/flux/), [documentation](https://metrico.github.io/fluXpipe/) or ["3 Minutes to Flux"](flux.md) guide.

### Demo
Try our [serverless demo](https://fluxpipe.fly.dev/) or launch your own instance to instantly fall in love with *flux*

<a href="https://flyctl.sh/shell?repo=metrico/fluxpipe" target="_blank">
  <img src="https://user-images.githubusercontent.com/1423657/236479471-a1cb0484-dfd3-4dc2-8d62-121debd7faa3.png" width=300>
</a>

<br>

### Instructions
Download a [binary release](https://github.com/metrico/fluxpipe/releases/), [docker](https://github.com/metrico/fluXpipe/pkgs/container/fluxpipe) or build from source

#### 📦 Download Binary
```bash
curl -fsSL github.com/metrico/fluxpipe/releases/latest/download/fluxpipe-server -O \
&& chmod +x fluxpipe-server
```
##### 🔌 Start Server w/ Options
```bash
./fluxpipe-server -port 8086
```
Run with `-h` for a full list of parameters

#### 🐋 Using Docker
```bash
docker pull ghcr.io/metrico/fluxpipe:latest
docker run -ti --rm -p 8086:8086 ghcr.io/metrico/fluxpipe:latest
```

<br>

### 🐛 Usage

💡 _Check out the [scripts](scripts) folder for working examples_

### Playground
Fluxpipe embeds a playground interface to instantly execute queries _(borrowed from ClickHouse [^2])_

<a href="https://fluxpipe.fly.dev"><img src="https://user-images.githubusercontent.com/1423657/231785057-c1762c65-7c30-483b-a551-57c316caee88.png"></a>

<br>

#### HTTP API
Fluxpipe serves a simple REST API loosely compatible with existing flux integrations and clients

##### Grafana Flux [^1]
Fluxpipe is compatible with the native **Grafana InfluxDB/Flux datasource** _(url + organization fields are required!)_

<br>

##### ⭐ FlightSQL
###### SQL
You can query InfluxDB 3.0 IOx with raw SQL using the native `sql.from` handler
```
import "sql"

sql.from(
    driverName: "influxdb-iox",
    dataSourceName: "iox://iox-server:443/qryn_logs",
    query: "SELECT level, sender, body FROM logs WHERE body LIKE '%DELETE%' limit 10",
)
```

![image](https://github.com/metrico/fluXpipe/assets/1423657/b6c2dcbe-079b-4329-9fee-a8601a8c853c)


###### Flux IOx
You can query InfluxDB 3.0 IOx with Flux using the `iox.from` handler 
```flux
import "contrib/qxip/iox"
iox.from(
     bucket: "test",
     host: "eu-central-1-1.aws.cloud2.influxdata.com:443",
     token: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
     limit: "10",
     columns: "time, level, sender",
     table: "logs",
     start: -100d,
)
```
```
                    _time:time            level:string           sender:string
------------------------------  ----------------------  ----------------------
2023-08-31T00:00:00.091362490Z                    info                 logtest
2023-08-31T00:00:00.091380034Z                    info                 logtest
2023-08-31T00:00:00.091381374Z                    info                 logtest
2023-08-31T00:00:00.091382470Z                    info                 logtest
2023-08-31T00:00:00.091383354Z                    info                 logtest
...
```

You write data back to InfluxDB 3.0 IOx using the `to` and ` wideTo` functions

```flux
import "contrib/qxip/iox"
import "influxdata/influxdb"
iox.from(
     bucket: "qxip",
     host: "eu-central-1-1.aws.cloud2.influxdata.com:443",
     token: "",
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
     token: "",
     orgID: "6a841c0c08328fb1"
    )

```



###### ⭐ ClickHouse SQL
```
import "contrib/qxip/clickhouse"

clickhouse.query(
  url: "https://play@play.clickhouse.com",
  query: "SELECT database, total_rows FROM tables WHERE total_rows > 0"
)
|> rename(columns: {database: "_value", total_rows: "_data"})
|> keep(columns: ["_value","_data"])
```
![image](https://user-images.githubusercontent.com/1423657/162625425-15a92f34-562b-4e27-8832-7bc33a90b185.png)

![image](https://user-images.githubusercontent.com/1423657/162428332-77d869a2-d02b-443d-a3ef-3df1fbf899f6.png)

###### ⭐ LogQL 
```
import "contrib/qxip/logql"

option logql.defaultURL = "http://qryn:3100"
logql.query_range(
     query: "rate({job=\"dummy-server\"}[5m])",
     start: v.timeRangeStart, 
     end: v.timeRangeStop
)
|> map(fn: (r) => ({r with _time: time(v: uint(v: r.timestamp_ns)), _value: float(v: r.value) }))
|> drop(columns: ["timestamp_ns", "value"])
|> sort(columns: ["_time"])
|> group(columns: ["labels"])
```
![image](https://user-images.githubusercontent.com/1423657/215287132-dc8e18ca-25f8-40cd-a925-cc9f6c090be5.png)


###### ⭐ CURL POST
Usage with curl

```bash
curl -XPOST localhost:8086/api/v2/query -sS \
  -H 'Accept:application/csv' \
  -H 'Content-type:application/vnd.flux' \
  -d 'import g "generate" g.from(start: 2022-04-01T00:00:00Z, stop: 2022-04-01T00:03:00Z, count: 3, fn: (n) => n)'
```

#### Secrets
Flux builds using `EnvironmentSecretService` accessing system environment variables from flux scripts.
```
import "influxdata/influxdb/secrets"
key = secrets.get(key: "ENV_SECRET")
```

------

#### STDIN CMD
Fluxpipe can be used as a command-line tool and stdin pipeline processor

###### Generate CSV
```bash
echo 'import g "generate" g.from(start: 2022-04-01T00:00:00Z, stop: 2022-04-01T00:03:00Z, count: 5, fn: (n) => 1)' \
| ./fluxpipe-server -stdin
```
##### Parse CSV
```bash
cat scripts/csv.flux | ./fluxpipe-server -stdin
```
##### Query SQL
```bash
cat scripts/sql.flux | ./fluxpipe-server -stdin
```

<br>

## Public Demo

### Grafana Datasource
Configure your Grafana instance with our public demo endpoint _(limited resources)_
![image](https://github.com/metrico/fluXpipe/assets/1423657/c100050a-49a6-4788-86af-15eb7e59c5d1)



### Documentation
**Flux**(pipe) is built using the [InfluxCommunity/flux](https://github.com/InfluxCommunity/flux) fork which contains additional features and contributions.<br>
All the standard and additional functions available in Fluxpipe are included in the [Flux community documentation](https://metrico.github.io/fluXpipe/)

<br>


#### Status
- [x] Fluxlib
  - [x] parser
  - [x] executor
- [x] Contribs
  - [x] contrib/qxip/clickhouse
  - [x] contrib/qxip/logql
  - [x] contrib/qxip/hash
  - [x] ENV secrets
- [x] STDIN pipeline
- [x] HTTP api
  - [x] plaintext
  - [x] json support
  - [x] web playground



[^1]: Project is not affiliated or endorsed by Influxdata or Grafana Labs. All rights belong to their respective owners.
[^2]: Used under Apache2 terms. Project is not affiliated or endorsed by ClickHouse Inc. All rights belong to their respective owners.
