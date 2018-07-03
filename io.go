package routing

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-lambda-go/events"
)

type (
	// Lambda Context
	Context *lambdacontext.LambdaContext

	// ResponseProxy
	ResponseProxy events.APIGatewayProxyResponse

	// RequestProxy
	RequestProxy events.APIGatewayProxyRequest

	// ResponseDispatcher interface basic to create an response with APIGatewayProxyResponse
	ResponseDispatcher interface {
		// ToJson return string to put in body's response
		ToJson() string
		// GetStatusCode return list of headers for put in response
		GetHeaders() map[string]string
	}

	// ResponseWithStatusDispatcher interface basic to create an response with APIGatewayProxyResponse with status code
	ResponseWithStatusDispatcher interface {
		ResponseDispatcher
		// GetStatusCode just to set status code of response
		GetStatusCode() int
	}


)
