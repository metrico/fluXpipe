import "experimental/http/requests"
import ejson "experimental/json"
import "json"
import "array"

response =
     requests.post(
         url: "https://goolnk.com/api/v1/shorten",
         body: json.encode(v: {url: "http://www.influxdata.com"}),
         headers: ["Content-Type": "application/json"],
     )

data = ejson.parse(data: response.body)
array.from(rows: [data])
