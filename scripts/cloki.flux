import "experimental/json"
import "experimental/http/requests"
import "experimental/array"

query = "{type=\"clickhouse\"}"
limit = "10"
api = "http://cloki-server:3100/loki/api/v1"
url = "${api}/query_range?limit=${limit}&query=${query}"

response = requests.get(url: url)
json_raw = string(v: response.body)
jsonData = json.parse(data: bytes(v: json_raw))

array.from(rows: jsonData.data.result[0].values
|> array.map(
    fn: (x) => ({ 
        _time: time(v: uint(v: x[0] )),
        _value: x[1]
})))
