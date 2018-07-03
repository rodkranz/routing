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
		// ToJSON return string to put in body's response
		ToJSON() string
	}

	// ResponseWithStatusDispatcher interface basic to create an response with APIGatewayProxyResponse with status code
	ResponseWithStatusDispatcher interface {
		ResponseDispatcher
		// GetStatusCode return list of headers for put in response
		GetHeaders() map[string]string
		// GetStatusCode just to set status code of response
		GetStatusCode() int
	}


)
