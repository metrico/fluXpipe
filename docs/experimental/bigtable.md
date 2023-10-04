## `bigtable` package

Package bigtable provides tools for working with data in
[Google Cloud Bigtable](https://cloud.google.com/bigtable/) databases.

Import the `bigtable` package:

```flux
import "bigtable"
```

### Functions

### `from()`

from retrieves data from a [Google Cloud Bigtable](https://cloud.google.com/bigtable/) data source.

#### Function type signature

```flux
(instance: string, project: string, table: string, token: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| token | token: Google Cloud IAM token to use to access the Cloud Bigtable database. | Yes |
| project | project: Cloud Bigtable project ID. | Yes |
| instance | instance: Cloud Bigtable instance ID. | Yes |
| table | table: Cloud Bigtable table name. | Yes |
