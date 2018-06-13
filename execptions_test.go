package routing

import (
	"testing"
	"fmt"
	"io"
	"net/http"
)

func TestErrNoSupportForMethod_Error(t *testing.T) {
	MockHTTPMethod := http.MethodGet

	actual := ErrNoSupportForMethod{HTTPMethod: MockHTTPMethod}.Error()
	expected := fmt.Sprintf("this API doesnt support [%s] http method", MockHTTPMethod)

	if actual != expected {
		t.Errorf("expected error message [%s], but got [%s]", expected, actual)
	}
}
func TestErrRouterNotFound_Error(t *testing.T) {
	MockResource := "/root"

	actual := ErrRouterNotFound{Resource: MockResource}.Error()
	expected := fmt.Sprintf("route [%s] not found", MockResource)

	if actual != expected {
		t.Errorf("expected error message [%s], but got [%s]", expected, actual)
	}
}
func TestErrDispatcher_Error(t *testing.T) {
	MockErr := io.ErrUnexpectedEOF

	actual := ErrDispatcher{Err: MockErr}.Error()
	expected := fmt.Sprintf("dispatcher error: %s", MockErr)

	if actual != expected {
		t.Errorf("expected error message [%s], but got [%s]", expected, actual)
	}
}
