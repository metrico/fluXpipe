# <img src="https://user-images.githubusercontent.com/1423657/162334409-293641a0-712a-45b7-96f2-30f9ca3bc22e.gif" height=60/> <img src="https://user-images.githubusercontent.com/1423657/162501021-e29869be-d41e-44e9-af2a-94204cf3ad2c.png" height=60 />

**FluxPipe** is a *very experimental* indepenent API **flux** pipeline for *embedded datasources*

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

### 🐛 Usage Examples
#### HTTP API
Fluxpipe serves a simple REST API loosely compatible with existing flux integrations and clients

##### Grafana Flux [^1]
Usage with native **Grafana InfluxDB/Flux datasource** _(url + organization fields are required!)_
```
import g "generate" 
g.from(start:  v.timeRangeStart, stop: v.timeRangeStop, count: 10, fn: (n) => n * 100 )
```
![image](https://user-images.githubusercontent.com/1423657/162274743-b454d3e6-e678-43aa-8ad6-8d612f2857b5.png)

![image](https://user-images.githubusercontent.com/1423657/162428332-77d869a2-d02b-443d-a3ef-3df1fbf899f6.png)

##### Generate CSV
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

##### Generate CSV
```bash
echo 'import g "generate" g.from(start: 2022-04-01T00:00:00Z, stop: 2022-04-01T00:03:00Z, count: 5, fn: (n) => 1)' \
| ./fluxpipe -stdin
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
cat scripts/csv.flux | ./fluxpipe -stdin
```
##### Query SQL
```bash
cat scripts/sql.flux | ./fluxpipe -stdin
```



## Status
- [x] stdin pipeline
- [x] http api
  - [x] plaintext
  - [x] json support
  - [ ] api doc
- [ ] output templates
- [ ] shared secrets
- [ ] bucket emulation


[^1]: Project is not affiliated or endorsed by Influxdata or Grafana Labs. All rights belong to their respective owners.
