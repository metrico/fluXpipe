import "experimental/http/requests"

 response = requests.do(
     method: "GET",
     url: "http://example.com",
     params: ["start": ["100"]],
 )

 requests.peek(response: response) 
