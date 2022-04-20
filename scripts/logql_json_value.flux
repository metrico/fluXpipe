/* Json Value Extraction Script
    Assuming a json object was asserted with
    some tag: label with value: key.
    The example uses a json object with a float value
    {
      value: 15.2
    }
*/

/* Require necessary imports */
import "experimental"
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
query = "{label=\"key\"}|json"
limit = "100"
server = "http://localhost:3100/loki/api/v1/query_range?limit=${limit}&query=${query}${range}"
/* Call REST Api with built query */
response = requests.get(url: server)
/* Process initial response */
jsonData = json.parse(data: response.body)

/* Build a table from an array of values */
array.from(rows: jsonData.data.result
|> array.map(
    fn: (x) => ({
        _raw: display(v: x),
        "Value": float(v: x.stream.value),
        "Timestamp": display(v: x.values[0][0])
      })
  )
)
/* calculate the mean over all the values in the Value column */
 |> mean(column: "Value")
