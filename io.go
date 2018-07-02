package routing

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-lambda-go/events"
)

type (
	Context *lambdacontext.LambdaContext
	Response events.APIGatewayProxyResponse
	Request events.APIGatewayProxyRequest
)
