package routing

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/aws/aws-lambda-go/events"
)

// Dispatcher is the method will invoke when router has been match
type Dispatcher func(events.APIGatewayProxyRequest) (interface{}, error)

// DispatchGroup is map of routers available in your lambda.
type DispatchGroup map[string]Dispatcher

// DispatchTable is map of methods available in your lambda.
type DispatchTable map[string]DispatchGroup

// Router redirect the correct path to correct dispatcher
type Router struct {
	rTable DispatchTable // rTable methods and routers available in you lambda
}

// New point of router struct
func New() *Router { return &Router{} }

// Register any kind of method and path for dispatcher
func (r *Router) Register(method, path string, dispatcher Dispatcher) *Router {
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

// Get Register router for requests method was GET
func (r *Router) Get(path string, dispatcher Dispatcher) *Router {
	return r.Register(http.MethodGet, path, dispatcher)
}

// Post Register router for requests method was Post
func (r *Router) Post(path string, dispatcher Dispatcher) *Router {
	return r.Register(http.MethodPost, path, dispatcher)
}

// Put Register router for requests method was Put
func (r *Router) Put(path string, dispatcher Dispatcher) *Router {
	return r.Register(http.MethodPut, path, dispatcher)
}

// Delete Register router for requests method was Delete
func (r *Router) Delete(path string, dispatcher Dispatcher) *Router {
	return r.Register(http.MethodDelete, path, dispatcher)
}

// Option Register router for requests method was Option
func (r *Router) Option(path string, dispatcher Dispatcher) *Router {
	return r.Register(http.MethodOptions, path, dispatcher)
}

// Lambda trigger the events to find router and http verb.
func (r Router) Lambda(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	routingMethod, ok := r.rTable[request.HTTPMethod]
	if ok == false {
		return events.APIGatewayProxyResponse{}, ErrNoSupportForMethod{HTTPMethod: request.HTTPMethod}
	}

	dispatcher, ok := routingMethod[request.Resource]
	if ok == false {
		return events.APIGatewayProxyResponse{}, ErrRouterNotFound{Resource: request.Resource}
	}

	response, err := dispatcher(request)
	if err != nil {
		return events.APIGatewayProxyResponse{}, ErrDispatcher{Err: err}
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint(response),
		StatusCode: http.StatusOK,
	}, nil
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
