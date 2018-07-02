package routing

import (
	"fmt"
)

type (
	// ErrNoSupportForMethod returns when http verb is not registered.
	ErrNoSupportForMethod struct {
		HTTPMethod string
	}

	// ErrRouterNotFound returns when path doesn't match with any one registered.
	ErrRouterNotFound struct {
		Resource string
	}

	// ErrDispatcher returns when dispatcher returns any error.
	ErrDispatcher struct {
		Err error
	}
)

func (e ErrNoSupportForMethod) Error() string {
	return fmt.Sprintf("this API doesnt support [%s] http method", e.HTTPMethod)
}

func (e ErrRouterNotFound) Error() string {
	return fmt.Sprintf("route [%s] not found", e.Resource)
}

func (e ErrDispatcher) Error() string {
	return fmt.Sprintf("dispatcher error: %s", e.Err)
}