import "csv"
import "experimental/http/requests"
import "strings"

url = "https://clickhouse.server"
query = "SELECT 'hello' as _value, 100 as _data FORMAT CSVWithNames"
response = requests.get(url:url, params: ["query": [query]])

csv.from(csv: string(v: response.body), mode: "raw")
|> keep(columns: ["_value","_data"])
