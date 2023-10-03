## `boundaries` package

Package boundaries provides operators for finding the boundaries around certain days, months, and weeks.

Import the `boundaries` package:

```flux
import "boundaries"
```

### Functions

### `friday()`

friday returns a record with `start` and `stop` boundary timestamps for last Friday.

#### Function type signature

```flux
() => {stop: time, start: time}
```

#### Parameters

No parameters provided.

### `monday()`

monday returns a record with `start` and `stop` boundary timestamps of last Monday.
Last Monday is relative to `now()`. If today is Monday, the function returns boundaries for the previous Monday.

#### Function type signature

```flux
() => {stop: time, start: time}
```

#### Parameters

No parameters provided.

### `month()`

month returns a record with `start` and `stop` boundary timestamps for the current month.

#### Function type signature

```flux
(?month_offset: int) => {stop: time, start: time}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| month_offset | month_offset: Number of months to offset from the current month. Default is `0`. | No |
### `saturday()`

saturday returns a record with `start` and `stop` boundary timestamps for last Saturday.

#### Function type signature

```flux
() => {stop: time, start: time}
```

#### Parameters

No parameters provided.

### `sunday()`

sunday returns a record with `start` and `stop` boundary timestamps for last Sunday.

#### Function type signature

```flux
() => {stop: time, start: time}
```

#### Parameters

No parameters provided.

### `thursday()`

thursday returns a record with `start` and `stop` boundary timestamps for last Thursday.

#### Function type signature

```flux
() => {stop: time, start: time}
```

#### Parameters

No parameters provided.

### `tuesday()`

tuesday returns a record with `start` and `stop` boundary timestamps of last Tuesday.

#### Function type signature

```flux
() => {stop: time, start: time}
```

#### Parameters

No parameters provided.

### `wednesday()`

wednesday returns a record with `start` and `stop` boundary timestamps for last Wednesday.

#### Function type signature

```flux
() => {stop: time, start: time}
```

#### Parameters

No parameters provided.

### `week()`

week returns a record with `start` and `stop` boundary timestamps of the current week.
By default, weeks start on Monday.

#### Function type signature

```flux
(?start_sunday: bool, ?week_offset: int) => {stop: time, start: time}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| start_sunday | start_sunday: Indicate if the week starts on Sunday. Default is `false`. | No |
| week_offset | week_offset: Number of weeks to offset from the current week. Default is `0`. | No |
### `yesterday()`

yesterday returns a record with `start` and `stop` boundary timestamps for yesterday.

#### Function type signature

```flux
() => {stop: time, start: time}
```

#### Parameters

No parameters provided.

