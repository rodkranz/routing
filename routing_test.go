package routing

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

type Test struct {
	Input       events.APIGatewayProxyRequest
	Action      func() *Router
	Expected    interface{} // events.APIGatewayProxyResponse
	ExpectedErr error
}

func TestNew(t *testing.T) {

	t.Run("New-With-Config", func(t *testing.T) {
		test := Test{
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodPatch,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return NewWithConfig(Config{DebugMode: true}).
					Register(http.MethodPatch,
						"/root",
						func(Context, RequestProxy) (interface{}, error) {
							return events.APIGatewayProxyResponse{
								StatusCode: http.StatusOK,
								Body:       "ovos",
							}, nil
						})
			},
			Expected: events.APIGatewayProxyResponse{
				Body:       "ovos",
				StatusCode: http.StatusOK,
			},
			ExpectedErr: nil,
		}

		resp, err := test.Action().LambdaProxy(context.Background(), test.Input)

		assert.Equal(t, test.ExpectedErr, err)
		assert.EqualValues(t, test.Expected, resp)
	})

	t.Run("Request-Method-"+http.MethodPatch, func(t *testing.T) {
		test := Test{
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodPatch,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Register(http.MethodPatch, "/root", func(Context, RequestProxy) (interface{}, error) {
					return events.APIGatewayProxyResponse{
						StatusCode: http.StatusOK,
						Body:       "ovos",
					}, nil
				})
			},
			Expected: events.APIGatewayProxyResponse{
				Body:       "ovos",
				StatusCode: http.StatusOK,
			},
			ExpectedErr: nil,
		}

		resp, err := test.Action().LambdaProxy(context.Background(), test.Input)

		assert.Equal(t, test.ExpectedErr, err)
		assert.EqualValues(t, test.Expected, resp)
	})

	t.Run("Request-Method-"+http.MethodGet, func(t *testing.T) {
		test := Test{
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodGet,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Get("/root", func(Context, RequestProxy) (interface{}, error) {
					return events.APIGatewayProxyResponse{
						StatusCode: http.StatusOK,
						Body:       "ovos",
					}, nil
				})
			},
			Expected: events.APIGatewayProxyResponse{
				Body:       "ovos",
				StatusCode: http.StatusOK,
			},
			ExpectedErr: nil,
		}

		resp, err := test.Action().LambdaProxy(context.Background(), test.Input)

		assert.Equal(t, test.ExpectedErr, err)
		assert.EqualValues(t, test.Expected, resp)
	})

	t.Run("Request-Method-"+http.MethodPost, func(t *testing.T) {
		test := Test{
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodPost,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Post("/root", func(Context, RequestProxy) (interface{}, error) {
					return events.APIGatewayProxyResponse{
						StatusCode: http.StatusOK,
						Body:       "ovos",
					}, nil
				})
			},
			Expected: events.APIGatewayProxyResponse{
				Body:       "ovos",
				StatusCode: http.StatusOK,
			},
			ExpectedErr: nil,
		}

		resp, err := test.Action().LambdaProxy(context.Background(), test.Input)

		assert.Equal(t, test.ExpectedErr, err)
		assert.EqualValues(t, test.Expected, resp)
	})

	t.Run("Request-Method-"+http.MethodPut, func(t *testing.T) {
		test := Test{
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodPut,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Put("/root", func(Context, RequestProxy) (interface{}, error) {
					return events.APIGatewayProxyResponse{
						StatusCode: http.StatusOK,
						Body:       "ovos",
					}, nil
				})
			},
			Expected: events.APIGatewayProxyResponse{
				Body:       "ovos",
				StatusCode: http.StatusOK,
			},
			ExpectedErr: nil,
		}

		resp, err := test.Action().LambdaProxy(context.Background(), test.Input)

		assert.Equal(t, test.ExpectedErr, err)
		assert.EqualValues(t, test.Expected, resp)
	})
	t.Run("Request-Method-"+http.MethodDelete, func(t *testing.T) {
		test := Test{
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodDelete,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Delete("/root", func(Context, RequestProxy) (interface{}, error) {
					return events.APIGatewayProxyResponse{
						StatusCode: http.StatusOK,
						Body:       "ovos",
					}, nil
				})
			},
			Expected: events.APIGatewayProxyResponse{
				Body:       "ovos",
				StatusCode: http.StatusOK,
			},
			ExpectedErr: nil,
		}

		resp, err := test.Action().LambdaProxy(context.Background(), test.Input)

		assert.Equal(t, test.ExpectedErr, err)
		assert.EqualValues(t, test.Expected, resp)
	})

	t.Run("Request-Method-NotFound", func(t *testing.T) {
		test := Test{
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodPut,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Option("/root", func(c Context, proxy RequestProxy) (interface{}, error) {
					return events.APIGatewayProxyResponse{}, nil
				})
			},
			Expected:    nil,
			ExpectedErr: ErrNoSupportForMethod{HTTPMethod: http.MethodPut},
		}

		resp, err := test.Action().LambdaProxy(context.Background(), test.Input)

		assert.Equal(t, test.ExpectedErr, err)
		assert.EqualValues(t, test.Expected, resp)
	})

	t.Run("Request-NotFound", func(t *testing.T) {
		test := Test{
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodGet,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Get("/", func(c Context, proxy RequestProxy) (interface{}, error) {
					return nil, nil
				})
			},
			Expected:    nil,
			ExpectedErr: ErrRouterNotFound{Resource: "/root"},
		}

		resp, err := test.Action().LambdaProxy(context.Background(), test.Input)

		assert.Equal(t, test.ExpectedErr, err)
		assert.EqualValues(t, test.Expected, resp)
	})

	t.Run("Request-Handler-Error", func(t *testing.T) {
		test := Test{
			Input: events.APIGatewayProxyRequest{
				Resource:   "/root",
				HTTPMethod: http.MethodGet,
				Body:       string("ovos"),
			},
			Action: func() *Router {
				return New().Get("/root", func(c Context, proxy RequestProxy) (interface{}, error) {
					return events.APIGatewayProxyResponse{}, io.ErrUnexpectedEOF
				})
			},
			Expected:    nil,
			ExpectedErr: ErrDispatcher{Err: io.ErrUnexpectedEOF},
		}

		resp, err := test.Action().LambdaProxy(context.Background(), test.Input)

		assert.Equal(t, test.ExpectedErr, err)
		assert.EqualValues(t, test.Expected, resp)
	})
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

			assert.Equal(t, test.Output, rsp)
		})
	}
}
