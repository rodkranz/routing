Router
---

Example: 

```go
package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rodkranz/routing"
)

const (
	hello = "/hello"
)

func main() {
	r := routing.New()

	r.Register(http.MethodGet, hello, func(context routing.Context, proxy routing.RequestProxy) (i interface{}, e error) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string("Hello world"),
		}, nil
	})

	lambda.Start(r.Lambda)
}
```