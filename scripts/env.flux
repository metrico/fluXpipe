import "array"
import "runtime"
import "influxdata/influxdb/secrets"
array.from(rows: [{version: runtime.version(), region: secrets.get(key: "FLY_REGION"), app: secrets.get(key: "FLY_APP_NAME")}])
