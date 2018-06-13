package routing

import (
	"net/http"
	"fmt"
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

type Router struct {
	rTable DispatchTable // rTable methods and routers available in you lambda
}

func New() *Router { return &Router{} }

func (r *Router) Register(method, path string, dispatcher Dispatcher) (*Router) {
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

func (r *Router) Get(path string, dispatcher Dispatcher) (*Router)    { return r.Register(http.MethodGet, path, dispatcher) }
func (r *Router) Post(path string, dispatcher Dispatcher) (*Router)   { return r.Register(http.MethodPost, path, dispatcher) }
func (r *Router) Put(path string, dispatcher Dispatcher) (*Router)    { return r.Register(http.MethodPut, path, dispatcher) }
func (r *Router) Delete(path string, dispatcher Dispatcher) (*Router) { return r.Register(http.MethodDelete, path, dispatcher) }
func (r *Router) Option(path string, dispatcher Dispatcher) (*Router) { return r.Register(http.MethodOptions, path, dispatcher) }

// Lambda
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

func trim(in string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, in)
}

