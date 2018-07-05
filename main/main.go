package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/rodkranz/routing"
)

func main() {
	r := routing.New()

	r.Register(http.MethodGet, "/hello", disp)
	r.Get("/method_get", disp)

	lambda.Start(r.Lambda)
}

func disp(req events.APIGatewayProxyRequest) (interface{}, error){
	m := map[string]interface{}{
		"path":   req.Path,
		"body":   req.Body,
		"method": req.HTTPMethod,
	}

	bs, err := json.Marshal(m)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("error to parse body: %s", err),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(bs),
	}, nil
}
