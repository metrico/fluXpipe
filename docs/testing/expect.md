## `expect` package

Package expect includes functions to mark
any expectations for a testcase to be satisfied
before the testcase finishes running.

Import the `expect` package:

```flux
import "expect"
```

### Functions

### `planner()`

planner will cause the present testcase to
expect the given planner rules will be invoked
exactly as many times as the number given.

#### Function type signature

```flux
(rules: [string:int]) => {}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| rules | rules: Mapping of rules names to expected counts. | Yes |
