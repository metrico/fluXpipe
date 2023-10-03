## `sampledata` package

Package sampledata provides functions that return basic sample datasets.

Import the `sampledata` package:

```flux
import "sampledata"
```

### Functions

### `bool()`

bool returns a sample data set with boolean values.

#### Function type signature

```flux
(?includeNull: bool) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| includeNull | includeNull: Include null values in the returned dataset.
  Default is `false`. | No |
### `float()`

float returns a sample data set with float values.

#### Function type signature

```flux
(?includeNull: bool) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| includeNull | includeNull: Include null values in the returned dataset.
  Default is `false`. | No |
### `int()`

int returns a sample data set with integer values.

#### Function type signature

```flux
(?includeNull: bool) => stream[{A with _value: B, _value: int}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| includeNull | includeNull: Include null values in the returned dataset.
  Default is `false`. | No |
### `numericBool()`

numericBool returns a sample data set with numeric (integer) boolean values.

#### Function type signature

```flux
(?includeNull: bool) => stream[{A with _value: B, _value: int}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| includeNull | includeNull: Include null values in the returned dataset.
  Default is `false`. | No |
### `start()`

start represents the earliest time included in sample datasets.

#### Function type signature

```flux
time
```

#### Parameters

No parameters provided.

### `stop()`

stop represents the latest time included in sample datasets.

#### Function type signature

```flux
time
```

#### Parameters

No parameters provided.

### `string()`

string returns a sample data set with string values.

#### Function type signature

```flux
(?includeNull: bool) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| includeNull | includeNull: Include null values in the returned dataset.
  Default is `false`. | No |
### `uint()`

uint returns a sample data set with unsigned integer values.

#### Function type signature

```flux
(?includeNull: bool) => stream[{A with _value: B, _value: uint}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| includeNull | includeNull: Include null values in the returned dataset.
  Default is `false`. | No |
