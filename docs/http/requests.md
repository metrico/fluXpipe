## `requests` package

Package requests provides functions for transferring data using the HTTP protocol.

Import the `requests` package:

```flux
import "requests"
```

### Functions

### `defaultConfig()`

defaultConfig is the global default for all http requests using the requests package.
Changing this config will affect all other packages using the requests package.
To change the config for a single request, pass a new config directly into the corresponding function.

#### Function type signature

```flux
{timeout: duration, insecureSkipVerify: bool}
```

#### Parameters

No parameters provided.

### `do()`

do makes an http request.

#### Function type signature

```flux
(
    method: string,
    url: string,
    ?body: bytes,
    ?config: {A with timeout: duration, insecureSkipVerify: bool},
    ?headers: [string:string],
    ?params: [string:[string]],
) => {statusCode: int, headers: [string:string], duration: duration, body: bytes}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| method | method: method of the http request.
     Supported methods: DELETE, GET, HEAD, PATCH, POST, PUT. | Yes |
| url | url: URL to request. This should not include any query parameters. | Yes |
| params | params: Set of key value pairs to add to the URL as query parameters.
    Query parameters will be URL encoded.
    All values for a key will be appended to the query. | No |
| headers | headers: Set of key values pairs to include on the request. | No |
| body | body: Data to send with the request. | No |
| config | config: Set of options to control how the request should be performed. | No |
### `get()`

get makes a http GET request. This identical to calling `request.do(method: "GET", ...)`.

#### Function type signature

```flux
(
    url: string,
    ?body: bytes,
    ?config: {A with timeout: duration, insecureSkipVerify: bool},
    ?headers: [string:string],
    ?params: [string:[string]],
) => {statusCode: int, headers: [string:string], duration: duration, body: bytes}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL to request. This should not include any query parameters. | Yes |
| params | params: Set of key value pairs to add to the URL as query parameters.
    Query parameters will be URL encoded.
    All values for a key will be appended to the query. | No |
| headers | headers: Set of key values pairs to include on the request. | No |
| body | body: Data to send with the request. | No |
| config | config: Set of options to control how the request should be performed. | No |
### `peek()`

peek converts an HTTP response into a table for easy inspection.

#### Function type signature

```flux
(
    response: {A with statusCode: E, headers: D, duration: C, body: B},
) => stream[{statusCode: E, headers: string, duration: int, body: string}]
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| response | response: Response data from an HTTP request. | Yes |
### `post()`

post makes a http POST request. This identical to calling `request.do(method: "POST", ...)`.

#### Function type signature

```flux
(
    url: string,
    ?body: bytes,
    ?config: {A with timeout: duration, insecureSkipVerify: bool},
    ?headers: [string:string],
    ?params: [string:[string]],
) => {statusCode: int, headers: [string:string], duration: duration, body: bytes}
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL to request. This should not include any query parameters. | Yes |
| params | params: Set of key value pairs to add to the URL as query parameters.
    Query parameters will be URL encoded.
    All values for a key will be appended to the query. | No |
| headers | headers: Set of key values pairs to include on the request. | No |
| body | body: Data to send with the request. | No |
| config | config: Set of options to control how the request should be performed. | No |
