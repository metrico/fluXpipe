## `join` package

Package join provides functions that join two table streams together.

Import the `join` package:

```flux
import "join"
```

### Functions

### `full()`

full performs a full outer join on two table streams.

#### Function type signature

```flux
(<-left: stream[A], as: (l: A, r: B) => C, on: (l: A, r: B) => bool, right: stream[B]) => stream[C] where A: Record, B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| left | left: Left input stream. Default is piped-forward data (<-). | No |
| right | right: Right input stream. | Yes |
| on | on: Function that takes a left and right record (`l`, and `r` respectively), and returns a boolean. | Yes |
| as | as: Function that takes a left and a right record (`l` and `r` respectively), and returns a record.
  The returned record is included in the final output. | Yes |
### `inner()`

inner performs an inner join on two table streams.

#### Function type signature

```flux
(<-left: stream[A], as: (l: A, r: B) => C, on: (l: A, r: B) => bool, right: stream[B]) => stream[C] where A: Record, B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| left | left: Left input stream. Default is piped-forward data (<-). | No |
| right | right: Right input stream. | Yes |
| on | on: Function that takes a left and right record (`l`, and `r` respectively), and returns a boolean. | Yes |
| as | as: Function that takes a left and a right record (`l` and `r` respectively), and returns a record.
  The returned record is included in the final output. | Yes |
### `left()`

left performs a left outer join on two table streams.

#### Function type signature

```flux
(<-left: stream[A], as: (l: A, r: B) => C, on: (l: A, r: B) => bool, right: stream[B]) => stream[C] where A: Record, B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| left | left: Left input stream. Default is piped-forward data (<-). | No |
| right | right: Right input stream. | Yes |
| on | on: Function that takes a left and right record (`l`, and `r` respectively), and returns a boolean. | Yes |
| as | as: Function that takes a left and a right record (`l` and `r` respectively), and returns a record.
  The returned record is included in the final output. | Yes |
### `right()`

right performs a right outer join on two table streams.

#### Function type signature

```flux
(<-left: stream[A], as: (l: A, r: B) => C, on: (l: A, r: B) => bool, right: stream[B]) => stream[C] where A: Record, B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| left | left: Left input stream. Default is piped-forward data (<-). | No |
| right | right: Right input stream. | Yes |
| on | on: Function that takes a left and right record (`l`, and `r` respectively), and returns a boolean. | Yes |
| as | as: Function that takes a left and a right record (`l` and `r` respectively), and returns a record.
  The returned record is included in the final output. | Yes |
### `tables()`

tables joins two input streams together using a specified method, predicate, and a function to join two corresponding records, one from each input stream.

#### Function type signature

```flux
(
    <-left: stream[A],
    as: (l: A, r: B) => C,
    method: string,
    on: (l: A, r: B) => bool,
    right: stream[B],
) => stream[C] where A: Record, B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| left | left: Left input stream. Default is piped-forward data (`<-`). | No |
| right | right: Right input stream. | Yes |
| on | on: Function that takes a left and right record (`l`, and `r` respectively), and returns a boolean. | Yes |
| as | as: Function that takes a left and a right record (`l` and `r` respectively), and returns a record.
  The returned record is included in the final output. | Yes |
| method | method: String that specifies the join method. | Yes |
### `time()`

time joins two table streams together exclusively on the `_time` column.

#### Function type signature

```flux
(
    <-left: stream[{A with _time: B}],
    as: (l: {A with _time: B}, r: {C with _time: D}) => E,
    right: stream[{C with _time: D}],
    ?method: string,
) => stream[E] where B: Equatable, D: Equatable, E: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| left | left: Left input stream. Default is piped-forward data (<-). | No |
| right | right: Right input stream. | Yes |
| as | as: Function that takes a left and a right record (`l` and `r` respectively), and returns a record.
  The returned record is included in the final output. | Yes |
| method | method: String that specifies the join method. Default is `inner`. | No |
