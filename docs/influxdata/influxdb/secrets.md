## `secrets` package

Package secrets functions for working with sensitive secrets managed by InfluxDB.

Import the `secrets` package:

```flux
import "secrets"
```

### Functions

### `get()`

get retrieves a secret from the InfluxDB secret store.

#### Function type signature

```flux
(key: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| key | key: Secret key to retrieve. | Yes |
