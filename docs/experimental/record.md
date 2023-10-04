## `record` package

Package record provides tools for working with Flux records.

Import the `record` package:

```flux
import "record"
```

### Functions

### `any()`

any is a polymorphic record value that can be used as a default record value
when input record property types are not known.

#### Function type signature

```flux
A where A: Record
```

#### Parameters

No parameters provided.

### `get()`

get returns a value from a record by key name or a default value if the key
doesn’t exist in the record.

#### Function type signature

```flux
(default: A, key: string, r: B) => A where B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| r | r: Record to retrieve the value from. | Yes |
| key | key: Property key to retrieve. | Yes |
| default | default: Default value to return if the specified key does not exist in the record. | Yes |
