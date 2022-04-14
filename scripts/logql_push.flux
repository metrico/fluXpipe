import "pushbullet"

api = "http://cloki:3100"

timestamp = string( v: uint(v: time(v: v.timeRangeStart)))
metric = 500

pushbullet.pushData(
    url: "${api}/loki/api/v1/push",
    token: "",
    data: { streams:[ { labels: "{\"type\":\"flux\"}", entries: [ {timestamp: "${timestamp}", value:"${metric}"}] }]},
)
