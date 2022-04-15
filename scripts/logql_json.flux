/* Unwrap the raw log string from a json wrapper */
import "experimental/json"
import "experimental/http/requests"
import "experimental/array"
import "strings"
import "dict"

tstart = string( v: uint(v: time(v: v.timeRangeStart)))
tend = string(v: uint(v: time(v: v.timeRangeStop)))

range = "&start${tstart}&end=${tend}"
query = "{label=\"key\"}"
limit = "100"
server = "http://localhost:3100/loki/api/v1/query_range?limit=${limit}&query=${query}${range}"
response = requests.get(url: server)
jsonData = json.parse(data: response.body)

array.from(rows: jsonData.data.result[0].values
|> array.map(
    fn: (x) => (
      {
        _time: time(v: uint(v: x[0] )),
        _value: x[1],
      }
    )
  )
)
|> map(fn: (r) => (
    {
      _time: display(v:r._time),
      _value:
        display(v: strings.split(
          v: strings.trim(
            v: strings.trim(
              v: r._value,
              cutset:"{"
            ),
            cutset:"}"
          ),
          t: ","
        )
      )
    }
  )
)

/* Json Value Extraction */

import "experimental"
import "experimental/json"
import "experimental/http/requests"
import "experimental/array"
import "strings"
import "dict"


tstart = string( v: uint(v: time(v: v.timeRangeStart)))
tend = string(v: uint(v: time(v: v.timeRangeStop)))

range = "&start${tstart}&end=${tend}"
query = "{label=\"value\"}|json"
limit = "100"
server = "http://localhost:3100/loki/api/v1/query_range?limit=${limit}&query=${query}${range}"
response = requests.get(url: server)
jsonData = json.parse(data: response.body)

array.from(rows: jsonData.data.result
|> array.map(
    fn: (x) => ({
        _raw: display(v: x),
        "Value": float(v: x.stream.value),
        "Timestamp": display(v: x.values[0][0])
      })
  )
)
 |> mean(column: "Value")
