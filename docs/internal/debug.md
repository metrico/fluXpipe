## `debug` package

Package debug provides methods for debugging the Flux engine.

Import the `debug` package:

```flux
import "debug"
```

### Functions

### `feature()`

feature returns the value associated with the given feature flag.

#### Function type signature

```flux
(key: string) => A
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| key | key: Feature flag name. | Yes |
### `getOption()`

getOption gets the value of an option using a form of reflection.

#### Function type signature

```flux
(name: string, pkg: string) => A
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| pkg | pkg: Full path of the package. | Yes |
| name | name: Option name. | Yes |
### `null()`

null returns the null value with a given type.

#### Function type signature

```flux
(?type: string) => A where A: Basic
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| type | type: Null type. | No |
### `opaque()`

opaque works like `pass` in that it passes any incoming tables directly to the
following transformation, save for its type signature does not indicate that the
input type has any correlation with the output type.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Stream to pass unmodified to next transformation. | No |
### `pass()`

pass will pass any incoming tables directly next to the following transformation.
It is best used to interrupt any planner rules that rely on a specific ordering.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Stream to pass unmodified to next transformation. | No |
### `sink()`

sink will discard all data that comes into it.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Stream to discard. | No |
### `slurp()`

slurp will read the incoming tables and concatenate buffers with the same group key
into a single in memory table buffer. This is useful for testing the performance impact of multiple
buffers versus a single buffer.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Stream to consume into single buffers per table. | No |
### `vectorize()`

vectorize controls whether the compiler attempts to vectorize Flux functions.

#### Function type signature

```flux
bool
```

#### Parameters

No parameters provided.

