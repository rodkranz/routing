package routing

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrNoSupportForMethod_Error(t *testing.T) {
	MockHTTPMethod := http.MethodGet

	actual := ErrNoSupportForMethod{HTTPMethod: MockHTTPMethod}.Error()
	expected := fmt.Sprintf("this API doesnt support [%s] http method", MockHTTPMethod)

	assert.Equal(t, expected, actual)
}
func TestIsErrNoSupportForMethod(t *testing.T) {
	t.Run("Should-Return-Ok", func(t *testing.T) {
		myErr := ErrNoSupportForMethod{
			HTTPMethod: http.MethodPost,
		}

		assert.True(t, IsErrNoSupportForMethod(myErr))
	})

	t.Run("Should-Return-Fail", func(t *testing.T) {
		myErr := errors.New("lorem Bacon")

		assert.False(t, IsErrNoSupportForMethod(myErr))
	})
}

func TestErrRouterNotFound_Error(t *testing.T) {
	MockResource := "/root"

	actual := ErrRouterNotFound{Resource: MockResource}.Error()
	expected := fmt.Sprintf("route [%s] not found", MockResource)

	assert.Equal(t, expected, actual)
}
func TestIsErrRouterNotFound(t *testing.T) {
	t.Run("Should-Return-Ok", func(t *testing.T) {
		myErr := ErrRouterNotFound{
			Resource: "/root",
		}

		assert.True(t, IsErrRouterNotFound(myErr))
	})

	t.Run("Should-Return-Fail", func(t *testing.T) {
		myErr := errors.New("lorem Bacon")

		assert.False(t, IsErrRouterNotFound(myErr))
	})
}

func TestErrDispatcher_Error(t *testing.T) {
	MockErr := io.ErrUnexpectedEOF

	actual := ErrDispatcher{Err: MockErr}.Error()
	expected := fmt.Sprintf("dispatcher error: %s", MockErr)

	assert.Equal(t, expected, actual)
}
func TestIsErrDispatcher(t *testing.T) {
	t.Run("Should-Return-Ok", func(t *testing.T) {
		myErr := ErrDispatcher{
			Err: io.ErrUnexpectedEOF,
		}

		assert.True(t, IsErrDispatcher(myErr))
	})

	t.Run("Should-Return-Fail", func(t *testing.T) {
		myErr := errors.New("lorem Bacon")

		assert.False(t, IsErrDispatcher(myErr))
	})
}
