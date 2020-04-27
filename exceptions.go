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

// IsErrNoSupportForMethod return if error type is the same of ErrNoSupportForMethod
func IsErrNoSupportForMethod(err error) bool {
	_, ok := err.(ErrNoSupportForMethod)
	return ok
}

// Error return message of error for ErrRouterNotFound
func (e ErrRouterNotFound) Error() string {
	return fmt.Sprintf("route [%s] not found", e.Resource)
}

// IsErrRouterNotFound return if error type is the same of ErrRouterNotFound
func IsErrRouterNotFound(err error) bool {
	_, ok := err.(ErrRouterNotFound)
	return ok
}


// Error return message of error for ErrDispatcher
func (e ErrDispatcher) Error() string {
	return fmt.Sprintf("dispatcher error: %s", e.Err)
}

// IsErrDispatcher return if error type is the same of ErrDispatcher
func IsErrDispatcher(err error) bool {
	_, ok := err.(ErrDispatcher)
	return ok
}
