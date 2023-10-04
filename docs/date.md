## `date` package

Package date provides date and time constants and functions.

Import the `date` package:

```flux
import "date"
```

### Functions

### `April()`

April is a constant that represents the month of April.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `August()`

August is a constant that represents the month of August.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `December()`

December is a constant that represents the month of December.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `February()`

February is a constant that represents the month of February.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `Friday()`

Friday is a constant that represents Friday as a day of the week.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `January()`

January is a constant that represents the month of January.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `July()`

July is a constant that represents the month of July.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `June()`

June is a constant that represents the month of June.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `March()`

March is a constant that represents the month of March.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `May()`

May is a constant that represents the month of May.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `Monday()`

Monday is a constant that represents Monday as a day of the week.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `November()`

November is a constant that represents the month of November.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `October()`

October is a constant that represents the month of October.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `Saturday()`

Saturday is a constant that represents Saturday as a day of the week.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `September()`

September is a constant that represents the month of September.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `Sunday()`

Sunday is a constant that represents Sunday as a day of the week.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `Thursday()`

Thursday is a constant that represents Thursday as a day of the week.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `Tuesday()`

Tuesday is a constant that represents Tuesday as a day of the week.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `Wednesday()`

Wednesday is a constant that represents Wednesday as a day of the week.

#### Function type signature

```flux
int
```

#### Parameters

No parameters provided.

### `add()`

add adds a duration to a time value and returns the resulting time value.

#### Function type signature

```flux
(d: duration, to: A, ?location: {zone: string, offset: duration}) => time where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| d | d: Duration to add. | Yes |
| to | to: Time to add the duration to. | Yes |
| location | location: Location to use for the time value. | No |
### `hour()`

hour returns the hour of a specified time. Results range from `[0 - 23]`.

#### Function type signature

```flux
(t: A, ?location: {zone: string, offset: duration}) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
### `microsecond()`

microsecond returns the microseconds for a specified time.
Results range `from [0-999999]`.

#### Function type signature

```flux
(t: A) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
### `millisecond()`

millisecond returns the milliseconds for a specified time.
Results range from `[0-999]`.

#### Function type signature

```flux
(t: A) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
### `minute()`

minute returns the minute of a specified time. Results range from `[0 - 59]`.

#### Function type signature

```flux
(t: A, ?location: {zone: string, offset: duration}) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
### `month()`

month returns the month of a specified time. Results range from `[1 - 12]`.

#### Function type signature

```flux
(t: A, ?location: {zone: string, offset: duration}) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
### `monthDay()`

monthDay returns the day of the month for a specified time.
Results range from `[1 - 31]`.

#### Function type signature

```flux
(t: A, ?location: {zone: string, offset: duration}) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
### `nanosecond()`

nanosecond returns the nanoseconds for a specified time.
Results range from `[0-999999999]`.

#### Function type signature

```flux
(t: A) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
### `quarter()`

quarter returns the quarter for a specified time. Results range from `[1-4]`.

#### Function type signature

```flux
(t: A, ?location: {zone: string, offset: duration}) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
### `scale()`

scale will multiply the duration by the given value.

#### Function type signature

```flux
(d: duration, n: int) => duration
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| d | d: Duration to scale. | Yes |
| n | n: Amount to scale the duration by. | Yes |
### `second()`

second returns the second of a specified time. Results range from `[0 - 59]`.

#### Function type signature

```flux
(t: A) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
### `sub()`

sub subtracts a duration from a time value and returns the resulting time value.

#### Function type signature

```flux
(d: duration, from: A, ?location: {zone: string, offset: duration}) => time where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| from | from: Time to subtract the duration from. | Yes |
| d | d: Duration to subtract. | Yes |
| location | location: Location to use for the time value. | No |
### `time()`

time returns the time value of a specified relative duration or time.

#### Function type signature

```flux
(t: A, ?location: {zone: string, offset: duration}) => time where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Duration or time value. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
### `truncate()`

truncate returns a time truncated to the specified duration unit.

#### Function type signature

```flux
(t: A, unit: duration, ?location: {zone: string, offset: duration}) => time where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
| unit | unit: Unit of time to truncate to. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
### `week()`

week returns the ISO week of the year for a specified time.
Results range from `[1 - 53]`.

#### Function type signature

```flux
(t: A, ?location: {zone: string, offset: duration}) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
### `weekDay()`

weekDay returns the day of the week for a specified time.
Results range from `[0 - 6]`.

#### Function type signature

```flux
(t: A, ?location: {zone: string, offset: duration}) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
### `year()`

year returns the year of a specified time.

#### Function type signature

```flux
(t: A, ?location: {zone: string, offset: duration}) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
### `yearDay()`

yearDay returns the day of the year for a specified time.
Results can include leap days and range from `[1 - 366]`.

#### Function type signature

```flux
(t: A, ?location: {zone: string, offset: duration}) => int where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| t | t: Time to operate on. | Yes |
| location | location: Location used to determine timezone.
  Default is the `location` option. | No |
