<img src="https://user-images.githubusercontent.com/1423657/161867564-4b3fc400-95e5-424c-9210-604d5671a85e.png" width=100 />

# FluxPipe
Experimental [flux](https://github.com/influxdata/flux) API pipeline for embedded datasources

### Instructions
Download a [binary release](https://github.com/lmangani/fluxpipe/releases/) or build from source


#### 📦 Download Binary
```bash
curl -fsSL github.com/lmangani/fluxpipe/releases/latest/download/fluxpipe -O && chmod +x fluxpipe
```

#### 📖 Build CMD
```bash
go build -o fluxpipe -ldflags="-s -w" fluxpipe.go
```

#### 📖 Build Server
```bash
go build -o fluxpipe -ldflags="-s -w" fluxpipe-server.go
```

### 🐛 Examples
#### HTTP API
##### Generate
```bash
./fluxpipe-server -port 8888 && \
curl -i -X POST localhost:8888/query --data-binary "@scripts/generate.flux"
```
```flux
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
#### STDIN CMD
##### Generate
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
- [x] http query api
- [ ] output templates
- [ ] shared secrets
