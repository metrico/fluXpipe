import "strings"
import "regexp"
import "experimental"
import "array"
import "http"

apiKey = "1234567890"
auth = http.basicAuth(u: "api", p: "${apiKey}")
url = "https://api.mailgun.net/v3/sandbox.mailgun.org/messages"
data =
    strings.joinStr(
        arr: [
            "from=mailgun@sandbox.mailgun.org",
            "to=qaservice@example.com",
            "subject=InfluxDB%20Alert",
            "text=This%20is%20a%20test%20notification%20FOO",
        ],
        v: "&",
    )

response = http.post(
    url: url,
    headers: {"Content-Type": "application/x-www-form-urlencoded", "Authorization": "${auth}"},
    data: bytes(v: data),
)

array.from(rows: [{value: data, response: response}])
    |> yield(name: "ignore")
