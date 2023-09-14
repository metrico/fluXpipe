## `csv` package

Package csv provides tools for working with data in annotated CSV format.

Import the `csv` package:

```flux
import "csv"
```

### Functions

### `from()`

from retrieves data from a comma separated value (CSV) data source and
returns a stream of tables.

#### Function type signature

```flux
(?csv: string, ?file: string, ?mode: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| csv | csv: CSV data. | No |
| file | file: File path of the CSV file to query. | No |
| mode | mode: is the CSV parsing mode. Default is `annotations`. | No |
