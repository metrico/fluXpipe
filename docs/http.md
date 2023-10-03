## `http` package

Package http provides functions for transferring data using the HTTP protocol.

Import the `http` package:

```flux
import "http"
```

### Functions

### `basicAuth()`

basicAuth returns a Base64-encoded basic authentication header
using a specified username and password combination.

#### Function type signature

```flux
(p: string, u: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| u | u: Username to use in the basic authentication header. | Yes |
| p | p: Password to use in the basic authentication header. | Yes |
### `endpoint()`

endpoint iterates over input data and sends a single POST request per input row to
a specficied URL.

#### Function type signature

```flux
(
    url: string,
) => (mapFn: (r: A) => {B with headers: C, data: bytes}) => (<-tables: stream[A]) => stream[{A with _sent: string}] where C: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL to send the POST reqeust to. | Yes |
### `pathEscape()`

pathEscape escapes special characters in a string (including `/`)
and replaces non-ASCII characters with hexadecimal representations (`%XX`).

#### Function type signature

```flux
(inputString: string) => string
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| inputString | inputString: String to escape. | Yes |
### `post()`

post sends an HTTP POST request to the specified URL with headers and data
and returns the HTTP status code.

#### Function type signature

```flux
(url: string, ?data: bytes, ?headers: A) => int where A: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL to send the POST request to. | Yes |
| headers | headers: Headers to include with the POST request. | No |
| data | data: Data body to include with the POST request. | No |
