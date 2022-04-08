import "http"
import "json"
import g "generate"

g.from(start:  v.timeRangeStart, stop: v.timeRangeStop, count: 10, fn: (n) => n * 100 )
|> mean()
|> map(fn: (r) => ({ r with jsonStr: string(v: json.encode(v: {"mean":r._value}))}))
|> map(fn: (r) => ({r with status_code: http.post(url: "https://example.com", headers: {x:"a", y:"b"}, data: bytes(v: r.jsonStr))}))
