## `csv` package

Package csv provides functions for retrieving annotated CSV.

Import the `csv` package:

```flux
import "csv"
```

### Functions

### `from()`

from retrieves [annotated CSV](https://docs.influxdata.com/influxdb/latest/reference/syntax/annotated-csv/) **from a URL**.

#### Function type signature

```flux
(url: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL to retrieve annotated CSV from. | Yes |
