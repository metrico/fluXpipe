## `timezone` package

Package timezone defines functions for setting timezones
on the location option in package universe.

Import the `timezone` package:

```flux
import "timezone"
```

### Functions

### `fixed()`

fixed returns a location record with a fixed offset.

#### Function type signature

```flux
(offset: A) => {zone: string, offset: A}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| offset | offset: Fixed duration for the location offset.
  This duration is the offset from UTC. | Yes |
### `location()`

location returns a location record based on a location or timezone name.

#### Function type signature

```flux
(name: string) => {zone: string, offset: duration}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| name | name: Location name (as defined by your operating system timezone database). | Yes |
### `utc()`

utc is the default location with a completely linear clock and no offset.
It is used as the default for location-related options.

#### Function type signature

```flux
{zone: string, offset: duration}
```

#### Parameters

No parameters provided.

