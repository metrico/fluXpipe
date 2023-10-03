## `dict` package

Package dict provides functions for interacting with dictionary types.

Import the `dict` package:

```flux
import "dict"
```

### Functions

### `fromList()`

fromList creates a dictionary from a list of records with `key` and `value`
properties.

#### Function type signature

```flux
(pairs: [{value: B, key: A}]) => [A:B] where A: Comparable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| pairs | pairs: List of records with `key` and `value` properties. | Yes |
### `get()`

get returns the value of a specified key in a dictionary or a default value
if the key does not exist.

#### Function type signature

```flux
(default: A, dict: [B:A], key: B) => A where B: Comparable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| dict | dict: Dictionary to return a value from. | Yes |
| key | key: Key to return from the dictionary. | Yes |
| default | default: Default value to return if the key does not exist in the
  dictionary. Must be the same type as values in the dictionary. | Yes |
### `insert()`

insert inserts a key-value pair into a dictionary and returns a new,
updated dictionary.

#### Function type signature

```flux
(dict: [A:B], key: A, value: B) => [A:B] where A: Comparable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| dict | dict: Dictionary to update. | Yes |
| key | key: Key to insert into the dictionary.
  Must be the same type as the existing keys in the dictionary. | Yes |
| value | value: Value to insert into the dictionary.
  Must be the same type as the existing values in the dictionary. | Yes |
### `remove()`

remove removes a key value pair from a dictionary and returns an updated
dictionary.

#### Function type signature

```flux
(dict: [A:B], key: A) => [A:B] where A: Comparable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| dict | dict: Dictionary to remove the key-value pair from. | Yes |
| key | key: Key to remove from the dictionary.
  Must be the same type as existing keys in the dictionary. | Yes |
