import "csv"
import "experimental"
import "experimental/http/requests"

response = requests.get(
    url: "https://example-s3-api-url.com/test/file.csv",
    headers: ["Authorization": "..."],
    body: bytes(v: "example-request-body")
)

csv.from(csv: string(v: response.body), mode: "raw")
