## `http` package

Package http provides functions for transferring data using HTTP protocol.

Import the `http` package:

```flux
import "http"
```

### Functions

### `get()`

get submits an HTTP GET request to the specified URL and returns the HTTP
status code, response body, and response headers.

#### Function type signature

```flux
(url: string, ?headers: A, ?timeout: duration) => {statusCode: int, headers: B, body: bytes} where A: Record, B: Record
```

#### Parameters

| Parameter | Description | Required |
| --- | --- | --- |
| url | url: URL to send the GET request to. | Yes |
| headers | headers: Headers to include with the GET request. | No |
| timeout | timeout: Timeout for the GET request. Default is `30s`. | No |
