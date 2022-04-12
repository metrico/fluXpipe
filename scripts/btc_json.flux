import "contrib/jsternberg/rows"
import "array"
import "experimental/json"
import "experimental/http/requests"
import "strings"
import "regexp"
import "experimental/array"
import "types"

api_key = "8e90abe4-b320-4af7-8591-08d29a6e5f1c"
url = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?CMC_PRO_API_KEY=${api_key}&limit=10"

response = requests.get(url:url)
json_s = string(v: response.body)

s1 = regexp.replaceAllString(r: /"platform":(null,|\{[^\}]+\},)/, v: json_s, t:"")
s2 = strings.replaceAll(v: s1, t: "\"self_reported_circulating_supply\":null,", u: "",)
s3 = strings.replaceAll(v: s2, t: "\"self_reported_market_cap\":null,", u: "")
s4 = strings.replaceAll(v: s3, t: "\"max_supply\":null,", u: "\"max_supply\": 0,")
s5 = strings.replaceAll(v: s4, t: "\"error_message\":null,", u: "")
s6 = strings.replaceAll(v: s5, t: "\"notice\":null,", u: "")

data = json.parse(data: bytes(v: s6))

array.from(rows: data.data |> 
    array.map(
    fn: (x) => ({ price:x.quote.USD.price, name: x.name, symbol: x.symbol, _time: time(v: x.last_updated), has_payment: length(arr: x.tags |> array.filter(fn: (x) => x == "payments"))
})))
