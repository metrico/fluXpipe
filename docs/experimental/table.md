## `table` package

Package table provides tools working with Flux tables.

Import the `table` package:

```flux
import "table"
```

### Functions

### `fill()`

fill adds a single row to empty tables in a stream of tables.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
