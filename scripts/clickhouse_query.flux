import "csv"
import "experimental/http/requests"
import "strings"

url = "https://clickhouse.server"
query = "SELECT 1 FORMAT CSVWithNames"

response = requests.get(url:url, params: ["query": [query]])

csv.from(csv: string(v: response.body), mode: "raw")
