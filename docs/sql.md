## `sql` package

Package sql provides tools for working with data in SQL databases.

Import the `sql` package:

```flux
import "sql"
```

### Functions

### `from()`

from retrieves data from a SQL data source.

#### Function type signature

```flux
(dataSourceName: string, driverName: string, query: string) => stream[A]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| driverName | driverName: Driver to use to connect to the SQL database. | Yes |
| dataSourceName | dataSourceName: Data source name (DNS) or connection string used to connect
  to the SQL database. | Yes |
| query | query: Query to run against the SQL database. | Yes |
### `to()`

to writes data to an SQL database.

#### Function type signature

```flux
(
    <-tables: stream[A],
    dataSourceName: string,
    driverName: string,
    table: string,
    ?batchSize: int,
) => stream[A]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| driverName | driverName: Driver used to connect to the SQL database. | Yes |
| dataSourceName | dataSourceName: Data source name (DNS) or connection string used
  to connect to the SQL database. | Yes |
| table | table: Destination table. | Yes |
| batchSize | batchSize: Number of parameters or columns that can be queued within each
  call to `Exec`. Default is `10000`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
