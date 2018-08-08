package routing

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-xray-sdk-go/xray"
)

type (
	// ResponseProxy response struct from AWS APIGatewayProxyResponse
	ResponseProxy events.APIGatewayProxyResponse

	// RequestProxy request struct from AWS APIGatewayProxyRequest
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

	// Context Context
	Context struct {
		context.Context
		*lambdacontext.LambdaContext
		XRaySegment *xray.Segment
	}
)
