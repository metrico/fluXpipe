## `tasks` package

Package tasks provides tools for working with InfluxDB tasks.

Import the `tasks` package:

```flux
import "tasks"
```

### Functions

### `lastSuccess()`

lastSuccess returns the time of the last successful run of the InfluxDB task
or the value of the `orTime` parameter if the task has never successfully run.

#### Function type signature

```flux
(orTime: A) => time where A: Timeable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| orTime | orTime: Default time value returned if the task has never successfully run. | Yes |
### `lastSuccessTime()`

lastSuccessTime is the last time this task ran successfully.

#### Function type signature

```flux
time
```

#### Parameters

No parameters provided.

