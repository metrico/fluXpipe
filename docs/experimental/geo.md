## `geo` package

Package geo provides tools for working with geotemporal data, such as
filtering and grouping by geographic location.

Import the `geo` package:

```flux
import "geo"
```

### Functions

### `ST_Contains()`

ST_Contains returns boolean indicating whether the defined region contains a
specified GIS geometry.

#### Function type signature

```flux
(geometry: A, region: B, ?units: {distance: string}) => bool where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| region | region: Region to test. Specify record properties for the shape. | Yes |
| geometry | geometry: GIS geometry to test. Can be either point or linestring geometry. | Yes |
| units | units: Record that defines the unit of measurement for distance.
  Default is the `geo.units` option. | No |
### `ST_DWithin()`

ST_DWithin tests if the specified region is within a defined distance from
the specified GIS geometry and returns `true` or `false`.

#### Function type signature

```flux
(distance: A, geometry: B, region: C, ?units: {distance: string}) => bool where A: Comparable + Equatable, B: Record, C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| region | region: Region to test. Specify record properties for the shape. | Yes |
| geometry | geometry: GIS geometry to test. Can be either point or linestring geometry. | Yes |
| distance | distance: Maximum distance allowed between the region and geometry.
  Define distance units with the `geo.units` option. | Yes |
| units | units: Record that defines the unit of measurement for distance.
  Default is the `geo.units` option. | No |
### `ST_Distance()`

ST_Distance returns the distance from a given region to a specified GIS geometry.

#### Function type signature

```flux
(geometry: A, region: B, ?units: {distance: string}) => float where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| region | region: Region to test. Specify record properties for the shape. | Yes |
| geometry | geometry: GIS geometry to test. Can be either point or linestring geometry. | Yes |
| units | units: Record that defines the unit of measurement for distance.
  Default is the `geo.units` option. | No |
### `ST_Intersects()`

ST_Intersects tests if the specified GIS geometry intersects with the
specified region and returns `true` or `false`.

#### Function type signature

```flux
(geometry: A, region: B, ?units: {distance: string}) => bool where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| region | region: Region to test. Specify record properties for the shape. | Yes |
| geometry | geometry: GIS geometry to test. Can be either point or linestring geometry. | Yes |
| units | units: Record that defines the unit of measurement for distance.
  Default is the `geo.units` option. | No |
### `ST_Length()`

ST_Length returns the [spherical length or distance](https://mathworld.wolfram.com/SphericalDistance.html)
of the specified GIS geometry.

#### Function type signature

```flux
(geometry: A, ?units: {distance: string}) => float where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| geometry | geometry: GIS geometry to test. Can be either point or linestring geometry.
  Point geometry will always return `0.0`. | Yes |
| units | units: Record that defines the unit of measurement for distance. | No |
### `ST_LineString()`

ST_LineString converts a series of geographic points into linestring.

#### Function type signature

```flux
(<-tables: stream[{A with lon: C, lat: B}]) => stream[D] where D: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `asTracks()`

asTracks groups rows into tracks (sequential, related data points).

#### Function type signature

```flux
(<-tables: stream[A], ?groupBy: [string], ?orderBy: [string]) => stream[A] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| groupBy | groupBy: Columns to group by. These columns should uniquely identify each track.
  Default is `["id","tid"]`. | No |
| orderBy | orderBy: Columns to order results by. Default is `["_time"]`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `filterRows()`

filterRows filters data by a specified geographic region with the option of strict filtering.

#### Function type signature

```flux
(
    <-tables: stream[{B with s2_cell_id: string, lon: D, lat: C}],
    region: A,
    ?level: int,
    ?maxSize: int,
    ?minSize: int,
    ?s2cellIDLevel: int,
    ?strict: bool,
) => stream[{B with s2_cell_id: string, lon: D, lat: C}] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| region | region: Region containing the desired data points. | Yes |
| minSize | minSize: Minimum number of cells that cover the specified region.
  Default is `24`. | No |
| maxSize | maxSize: Maximum number of cells that cover the specified region.
  Default is `-1` (unlimited). | No |
| level | level: [S2 cell level](https://s2geometry.io/resources/s2cell_statistics.html)
  of grid cells. Default is `-1`. | No |
| s2cellIDLevel | s2cellIDLevel: [S2 cell level](https://s2geometry.io/resources/s2cell_statistics.html)
  used in the `s2_cell_id` tag. Default is `-1` (detects S2 cell level from the `s2_cell_id` tag). | No |
| strict | strict: Enable strict geographic data filtering. Default is `true`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `getGrid()`

getGrid calculates a grid or set of cell ID tokens for a specified region.

#### Function type signature

```flux
(
    region: A,
    units: {distance: string},
    ?level: int,
    ?maxLevel: int,
    ?maxSize: int,
    ?minSize: int,
) => {set: [string], level: int} where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| region | region: Region used to return S2 cell ID tokens.
  Specify record properties for the region shape. | Yes |
| minSize | minSize: Minimum number of cells that cover the specified region. | No |
| maxSize | maxSize: Minimum number of cells that cover the specified region. | No |
| level | level: S2 cell level of grid cells. | No |
| maxLevel | maxLevel: Maximumn S2 cell level of grid cells. | No |
| units | units: Record that defines the unit of measurement for distance. | Yes |
### `getLevel()`

getLevel returns the S2 cell level of specified cell ID token.

#### Function type signature

```flux
(token: string) => int
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| token | token: S2 cell ID token. | Yes |
### `gridFilter()`

gridFilter filters data by a specified geographic region.

#### Function type signature

```flux
(
    <-tables: stream[{B with s2_cell_id: string}],
    region: A,
    ?level: int,
    ?maxSize: int,
    ?minSize: int,
    ?s2cellIDLevel: int,
    ?units: {distance: string},
) => stream[{B with s2_cell_id: string}] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| region | region: Region containing the desired data points. | Yes |
| minSize | minSize: Minimum number of cells that cover the specified region.
  Default is `24`. | No |
| maxSize | maxSize: Maximum number of cells that cover the specified region.
  Default is `-1` (unlimited). | No |
| level | level: [S2 cell level](https://s2geometry.io/resources/s2cell_statistics.html)
  of grid cells. Default is `-1`. | No |
| s2cellIDLevel | s2cellIDLevel: [S2 cell level](https://s2geometry.io/resources/s2cell_statistics.html)
  used in the `s2_cell_id` tag. Default is `-1` (detects S2 cell level from the S2 cell ID token). | No |
| units | units: Record that defines the unit of measurement for distance.
  Default is the `geo.units` option. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `groupByArea()`

groupByArea groups rows by geographic area.

#### Function type signature

```flux
(
    <-tables: stream[{A with s2_cell_id: string, lon: float, lat: float}],
    level: int,
    newColumn: string,
    ?s2cellIDLevel: int,
) => stream[B] where B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| newColumn | newColumn: Name of the new column for the unique identifier for each geographic area. | Yes |
| level | level: [S2 Cell level](https://s2geometry.io/resources/s2cell_statistics.html)
  used to determine the size of each geographic area. | Yes |
| s2cellIDLevel | s2cellIDLevel: [S2 Cell level](https://s2geometry.io/resources/s2cell_statistics.html)
  used in the `s2_cell_id` tag. Default is `-1` (detects S2 cell level from the `s2_cell_id` tag). | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `s2CellIDToken()`

s2CellIDToken returns and S2 cell ID token for given cell or point at a
specified S2 cell level.

#### Function type signature

```flux
(level: int, ?point: {lon: float, lat: float}, ?token: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| token | token: S2 cell ID token to update. | No |
| point | point: Record with `lat` and `lon` properties that specify the latitude and
  longitude in decimal degrees (WGS 84) of a point. | No |
| level | level: S2 cell level to use when generating the S2 cell ID token. | Yes |
### `s2CellLatLon()`

s2CellLatLon returns the latitude and longitude of the center of an S2 cell.

#### Function type signature

```flux
(token: string) => {lon: float, lat: float}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| token | token: S2 cell ID token. | Yes |
### `shapeData()`

shapeData renames existing latitude and longitude fields to **lat** and **lon**
and adds an **s2\_cell\_id** tag.

#### Function type signature

```flux
(
    <-tables: stream[{C with _field: string}],
    latField: A,
    level: int,
    lonField: B,
) => stream[{D with s2_cell_id: string, lon: float, lat: float}] where A: Equatable, B: Equatable
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| latField | latField: Name of the existing field that contains the latitude value in decimal degrees (WGS 84). | Yes |
| lonField | lonField: Name of the existing field that contains the longitude value in decimal degrees (WGS 84). | Yes |
| level | level: [S2 cell level](https://s2geometry.io/resources/s2cell_statistics.html)
  to use when generating the S2 cell ID token. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `stContains()`

stContains returns boolean indicating whether the defined region contains a specified GIS geometry.

#### Function type signature

```flux
(geometry: A, region: B, units: {distance: string}) => bool where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| region | region: Region to test. Specify record properties for the shape. | Yes |
| geometry | geometry: GIS geometry to test. Can be either point or linestring geometry. | Yes |
| units | units: Record that defines the unit of measurement for distance. | Yes |
### `stDistance()`

stDistance returns the distance from a given region to a specified GIS geometry.

#### Function type signature

```flux
(geometry: A, region: B, units: {distance: string}) => float where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| region | region: Region to test. Specify record properties for the shape. | Yes |
| geometry | geometry: GIS geometry to test. Can be either point or linestring geometry. | Yes |
| units | units: Record that defines the unit of measurement for distance. | Yes |
### `stLength()`

stLength returns the [spherical length or distance](https://mathworld.wolfram.com/SphericalDistance.html)
of the specified GIS geometry.

#### Function type signature

```flux
(geometry: A, units: {distance: string}) => float where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| geometry | geometry: GIS geometry to test. Can be either point or linestring geometry.
  Point geometry will always return `0.0`. | Yes |
| units | units: Record that defines the unit of measurement for distance. | Yes |
### `strictFilter()`

strictFilter filters data by latitude and longitude in a specified region.

#### Function type signature

```flux
(<-tables: stream[{B with lon: D, lat: C}], region: A) => stream[{B with lon: D, lat: C}] where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| region | region: Region containing the desired data points. | Yes |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `toRows()`

toRows pivots fields into columns based on time.

#### Function type signature

```flux
(<-tables: stream[A]) => stream[B] where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `totalDistance()`

totalDistance calculates the total distance covered by subsequent points
in each input table.

#### Function type signature

```flux
(<-tables: stream[{B with lon: float, lat: float}], ?outputColumn: A) => stream[C] where C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| outputColumn | outputColumn: Total distance output column. Default is `_value`. | No |
| tables | tables: Input data. Default is piped-forward data (`<-`). | No |
### `units()`

units defines the unit of measurment used in geotemporal operations.

#### Function type signature

```flux
{distance: string}
```

#### Parameters

No parameters provided.

