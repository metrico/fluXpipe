## `sample` package

Package sample provides functions for downloading and outputting InfluxDB sample datasets.

Import the `sample` package:

```flux
import "sample"
```

### Functions

### `alignToNow()`

alignToNow shifts time values in input data to align the chronological last point to _now_.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Defaults to piped-forward data (`<-`). | No |
### `data()`

data downloads a specified InfluxDB sample dataset.

#### Function type signature

```flux
(set: string) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| set | set: Sample data set to download and output. | Yes |
### `list()`

list outputs information about available InfluxDB sample datasets.

#### Function type signature

```flux
(
    
) => stream[{
    url: string,
    type: string,
    size: string,
    name: string,
    description: string,
}]
```

#### Parameters

No parameters provided.

