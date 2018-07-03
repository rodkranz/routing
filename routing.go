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

// New point of router struct
func New() *Router { return &Router{c: Config{}} }

// New point of router struct
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

// FnLambda type of lambda
type FnLambdaProxy func(context.Context, events.APIGatewayProxyRequest) (interface{}, error)

// Lambda trigger the events to find router and http verb.
func (r Router) LambdaProxy(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {
	// Find Method corresponding in our group router
	routingMethod, ok := r.rTable[request.HTTPMethod]
	if ok == false {
		return events.APIGatewayProxyResponse{}, ErrNoSupportForMethod{HTTPMethod: request.HTTPMethod}
	}

	// Find Path corresponding in our group router
	dispatcher, ok := routingMethod[request.Resource]
	if ok == false {
		return events.APIGatewayProxyResponse{}, ErrRouterNotFound{Resource: request.Resource}
	}

	// Get Lambda context
	c, _ := lambdacontext.FromContext(ctx)

	// execute dispatcher corresponding to Path and Method
	response, err := dispatcher(Context(c), RequestProxy(request))
	if err != nil {
		return events.APIGatewayProxyResponse{}, ErrDispatcher{Err: err}
	}

	// case response is already APIGatewayProxyResponse
	if r, ok := response.(events.APIGatewayProxyResponse); ok {
		return r, nil
	}

	// case response has methods to create an APIGatewayProxyResponse
	if r, ok := response.(ResponseWithStatusDispatcher); ok {
		return events.APIGatewayProxyResponse{
			Body:       r.ToJson(),
			StatusCode: r.GetStatusCode(),
			Headers:    r.GetHeaders(),
		}, err
	}

	// case response has methods to create an APIGatewayProxyResponse
	if r, ok := response.(ResponseDispatcher); ok {
		return events.APIGatewayProxyResponse{
			Body:       r.ToJson(),
			StatusCode: http.StatusOK,
			Headers:    r.GetHeaders(),
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
