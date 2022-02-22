package bus

import (
	"context"
	"errors"
	"fmt"
	"reflect"


)

// HandlerFunc defines a handler function interface
type HandlerFunc interface{}

// Msg defines a message interface.
type Msg interface{}

var ErrHandlerNotFound = errors.New("handler not found")

// TransactionManager defines a transaction interface
type TransactionManager interface {
	InTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

// Bus type defines the bus interface structure
type Bus interface {
	Dispatch(ctx context.Context, msg Msg) error

	Publish(ctx context.Context, msg Msg) error


}

// InProcBus defines the bus structure
type InProcBus struct {
	handlers map[string]HandlerFunc
}

// New initialize the bus
func New() *InProcBus {
	bus := &InProcBus{
		handlers:         make(map[string]HandlerFunc),
	}
	return bus
}

