## `iox` package

Package iox provides functions for querying data from IOx.

Import the `iox` package:

```flux
import "iox"
```

### Functions

### `from()`

from reads from the selected bucket and measurement in an IOx storage node.

#### Function type signature

```flux
(bucket: string, measurement: string) => stream[{A with _time: time}] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: IOx bucket to read data from. | Yes |
| measurement | measurement: Measurement to read data from. | Yes |
### `sql()`

sql executes an SQL query against a bucket in an IOx storage node.

#### Function type signature

```flux
(bucket: string, query: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| bucket | bucket: IOx bucket to read data from. | Yes |
| query | query: SQL query to execute. | Yes |
### `sqlInterval()`

sqlInterval converts a duration value to a SQL interval string.

#### Function type signature

```flux
(d: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| d | d: Duration value to convert to SQL interval string. | Yes |
