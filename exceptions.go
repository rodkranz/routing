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

// Error return message of error for ErrNoSupportForMethod
func (e ErrNoSupportForMethod) Error() string {
	return fmt.Sprintf("this API doesnt support [%s] http method", e.HTTPMethod)
}

// Error return message of error for ErrRouterNotFound
func (e ErrRouterNotFound) Error() string {
	return fmt.Sprintf("route [%s] not found", e.Resource)
}

// Error return message of error for ErrDispatcher
func (e ErrDispatcher) Error() string {
	return fmt.Sprintf("dispatcher error: %s", e.Err)
}
