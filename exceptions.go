package routing

import (
	"fmt"
)

type (
	ErrNoSupportForMethod struct {
		HTTPMethod string
	}

	ErrRouterNotFound struct {
		Resource string
	}

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
