package routing

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

type Test struct {
	Name        string
	Input       events.APIGatewayProxyRequest
	Action      func() *Router
	Expected    events.APIGatewayProxyResponse
	ExpectedErr error
}

func TestNew(t *testing.T) {
	tests := []Test{
		{
			Name: "Method ",
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodPatch,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Register(http.MethodPatch, "/root", func(request events.APIGatewayProxyRequest) (interface{}, error) {
					return "ovos", nil
				})
			},
			Expected: events.APIGatewayProxyResponse{
				Body:       "ovos",
				StatusCode: http.StatusOK,
			},
			ExpectedErr: nil,
		},
		// Test method GET
		{
			Name: "Method GET",
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodGet,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Get("/root", func(request events.APIGatewayProxyRequest) (interface{}, error) {
					return "ovos", nil
				})
			},
			Expected: events.APIGatewayProxyResponse{
				Body:       "ovos",
				StatusCode: http.StatusOK,
			},
			ExpectedErr: nil,
		},
		// Test method POST
		{
			Name: "Method POST",
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodPost,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Post("/root", func(request events.APIGatewayProxyRequest) (interface{}, error) {
					return "ovos", nil
				})
			},
			Expected: events.APIGatewayProxyResponse{
				Body:       "ovos",
				StatusCode: http.StatusOK,
			},
			ExpectedErr: nil,
		},
		// Test method PUT
		{
			Name: "Method PUST",
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodPut,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Put("/root", func(request events.APIGatewayProxyRequest) (interface{}, error) {
					return "ovos", nil
				})
			},
			Expected: events.APIGatewayProxyResponse{
				Body:       "ovos",
				StatusCode: http.StatusOK,
			},
			ExpectedErr: nil,
		},
		// Method not found
		{
			Name: "Method not found",
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodGet,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Option("/root", func(request events.APIGatewayProxyRequest) (interface{}, error) {
					return "ovos", nil
				})
			},
			Expected:    events.APIGatewayProxyResponse{},
			ExpectedErr: ErrNoSupportForMethod{HTTPMethod: http.MethodGet},
		},
		// Router not found
		{
			Name: "Router not found",
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodDelete,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Delete("/not-found", func(request events.APIGatewayProxyRequest) (interface{}, error) {
					return "ovos", nil
				})
			},
			Expected:    events.APIGatewayProxyResponse{},
			ExpectedErr: ErrRouterNotFound{Resource: "/root"},
		},
		// Handler error
		{
			Name: "Handler error",
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodGet,
			},
			Action: func() *Router {
				return New().Get("/root", func(request events.APIGatewayProxyRequest) (interface{}, error) {
					return nil, io.ErrUnexpectedEOF
				})
			},
			Expected:    events.APIGatewayProxyResponse{},
			ExpectedErr: ErrDispatcher{Err: io.ErrUnexpectedEOF},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.Name), func(t *testing.T) {
			resp, err := test.Action().Lambda(test.Input)
			if err != test.ExpectedErr {
				t.Errorf("expected error [%s], but got: [%s]", test.ExpectedErr, err)
			}

			if !reflect.DeepEqual(resp, test.Expected) {
				t.Errorf("expected error [%v], but got: [%v]", test.Expected, resp)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "lorem ipsum   dolor amet    lot",
			Output: "loremipsumdolorametlot",
		},
		{
			Input:  "lorem ipsum dolor amet lot",
			Output: "loremipsumdolorametlot",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			rsp := trim(test.Input)
			if test.Output != rsp {
				t.Errorf("expected error [%s], but got: [%s]", test.Output, rsp)
			}
		})
	}
}
