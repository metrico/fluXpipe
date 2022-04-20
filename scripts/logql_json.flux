/* Unwrap the raw log string from a json wrapper
    We assume a log string was asserted with
    tag: label and value: key.

    We assume values are JSON key:value pairs such as:
    {
      "test": "value",
      "number": "12"
    }
 */

/* Require necessary imports */
import "experimental/json"
import "experimental/http/requests"
import "experimental/array"
import "strings"
import "dict"

/* Build time range from external input (v) */
tstart = string( v: uint(v: time(v: v.timeRangeStart)))
tend = string(v: uint(v: time(v: v.timeRangeStop)))

/* build query from inputs */
range = "&start${tstart}&end=${tend}"
query = "{label=\"key\"}"
limit = "100"
server = "http://localhost:3100/loki/api/v1/query_range?limit=${limit}&query=${query}${range}"
/* Call REST Api with built query */
response = requests.get(url: server)
/* Process initial response */
jsonData = json.parse(data: response.body)

/* Build a table from an array of values */
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
