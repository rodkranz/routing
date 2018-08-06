package routing

import (
	"net/http"
	"strings"
	"unicode"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type (
	// Dispatcher is the method will invoke when router has been match
	Dispatcher func(Context, RequestProxy) (interface{}, error)

	// DispatchGroup is map of routers available in your lambda.
	DispatchGroup map[string]Dispatcher

	// DispatchTable is map of methods available in your lambda.
	DispatchTable map[string]DispatchGroup

	// Router redirect the correct path to correct dispatcher
	Router struct {
		c      Config        // configuration of router
		rTable DispatchTable // rTable methods and routers available in you lambda
	}
)

// Config for help developers knows what is going on inside router.
type Config struct {
	DebugMode bool
}

// New returns new pointer of router struct with config default
func New() *Router { return &Router{c: Config{}} }

// NewWithConfig returns new pointer of router instance with custom configuration
func NewWithConfig(c Config) *Router { return &Router{c: c} }

// register all routers and methods and centralize all registers
func (r *Router) register(method, path string, dispatcher Dispatcher) *Router {
	if r.rTable == nil {
		r.rTable = make(DispatchTable)
	}

	method = strings.ToUpper(method)
	if _, has := r.rTable[method]; !has {
		r.rTable[method] = make(DispatchGroup)
	}

	r.rTable[method][trim(path)] = dispatcher
	return r
}

// Register any kind of method and path for dispatcher
func (r *Router) Register(method, path string, dispatcher Dispatcher) *Router {
	return r.register(method, path, dispatcher)
}

// Get Register router for requests method was GET
func (r *Router) Get(path string, dispatcher Dispatcher) *Router {
	return r.register(http.MethodGet, path, dispatcher)
}

// Post Register router for requests method was Post
func (r *Router) Post(path string, dispatcher Dispatcher) *Router {
	return r.register(http.MethodPost, path, dispatcher)
}

// Put Register router for requests method was Put
func (r *Router) Put(path string, dispatcher Dispatcher) *Router {
	return r.register(http.MethodPut, path, dispatcher)
}

// Delete Register router for requests method was Delete
func (r *Router) Delete(path string, dispatcher Dispatcher) *Router {
	return r.register(http.MethodDelete, path, dispatcher)
}

// Option Register router for requests method was Option
func (r *Router) Option(path string, dispatcher Dispatcher) *Router {
	return r.register(http.MethodOptions, path, dispatcher)
}

// FnLambdaProxy type of lambda function
type FnLambdaProxy func(context.Context, events.APIGatewayProxyRequest) (interface{}, error)

// LambdaProxy trigger the events to find router and http verb.
func (r Router) LambdaProxy(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {
	// Find Method corresponding in our group router
	routingMethod, ok := r.rTable[request.HTTPMethod]
	if ok == false {
		return nil, ErrNoSupportForMethod{HTTPMethod: request.HTTPMethod}
	}

	// Find Path corresponding in our group router
	dispatcher, ok := routingMethod[request.Resource]
	if ok == false {
		return nil, ErrRouterNotFound{Resource: request.Resource}
	}

	// Get Lambda context
	c, _ := lambdacontext.FromContext(ctx)
	ctxLambda := Context{LambdaContext: c, Context: ctx}

	// execute dispatcher corresponding to Path and Method
	response, err := dispatcher(ctxLambda, RequestProxy(request))
	if err != nil {
		return nil, ErrDispatcher{Err: err}
	}

	// case response is already APIGatewayProxyResponse
	if r, ok := response.(events.APIGatewayProxyResponse); ok {
		return r, nil
	}

	// case response has methods to create an APIGatewayProxyResponse
	if r, ok := response.(ResponseWithStatusDispatcher); ok {
		return events.APIGatewayProxyResponse{
			Body:       r.ToJSON(),
			StatusCode: r.GetStatusCode(),
			Headers:    r.GetHeaders(),
		}, err
	}

	// case response has methods to create an APIGatewayProxyResponse
	if r, ok := response.(ResponseDispatcher); ok {
		return events.APIGatewayProxyResponse{
			Body:       r.ToJSON(),
			StatusCode: http.StatusOK,
		}, err
	}

	// if anything match then it will return the response
	return response, nil
}

// trim remove spaces from string
func trim(in string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, in)
}
