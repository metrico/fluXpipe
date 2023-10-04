## `hash` package

Package hash provides functions that convert string values to hashes.

Import the `hash` package:

```flux
import "hash"
```

### Functions

### `b64()`

b64 converts a string value to a Base64 string.

#### Function type signature

```flux
(v: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to hash. | Yes |
### `cityhash64()`

cityhash64 converts a string value to a 64-bit hexadecimal hash using the CityHash64 algorithm.

#### Function type signature

```flux
(v: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to hash. | Yes |
### `hmac()`

hmac converts a string value to an MD5-signed SHA-1 hash.

#### Function type signature

```flux
(k: A, v: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to hash. | Yes |
| k | k: Key to sign hash. | Yes |
### `md5()`

md5 converts a string value to an MD5 hash.

#### Function type signature

```flux
(v: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to hash. | Yes |
### `sha1()`

sha1 converts a string value to a hexadecimal hash using the SHA-1 hash algorithm.

#### Function type signature

```flux
(v: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to hash. | Yes |
### `sha256()`

sha256 converts a string value to a hexadecimal hash using the SHA 256 hash algorithm.

#### Function type signature

```flux
(v: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to hash. | Yes |
### `xxhash64()`

xxhash64 converts a string value to a 64-bit hexadecimal hash using the xxHash algorithm.

#### Function type signature

```flux
(v: A) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| v | v: String to hash. | Yes |
