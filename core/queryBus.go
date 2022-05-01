package core

import (
	"context"
	"errors"
	"reflect"
)

// HandlerFunc defines a function to execute the query handler.
type HandlerFunc func(context.Context, interface{}) (interface{}, error)
type handlers map[string]IQueryHandler

type bus struct {
	handlers handlers
}

// Register assign a query filter to a query handler for future executions.
func (b *bus) Register(queryFilter interface{}, handlerFunc IQueryHandler) error {
	if err := b.validate(queryFilter); err != nil {
		return err
	}

	queryFilterName := reflect.TypeOf(queryFilter).String()

	if _, err := b.handler(queryFilterName); err == nil {
		return errors.New("query filter already assigned to a handler")
	}

	b.handlers[queryFilterName] = handlerFunc

	return nil
}

// GetHandlers returns all registered query handlers.
func (b *bus) GetHandlers() handlers {
	return b.handlers
}

// Execute send a given query filter to its assigned query handler.
func (b *bus) Execute(ctx context.Context, queryFilter interface{}) (interface{}, error) {
	if err := b.validate(queryFilter); err != nil {
		return nil, err
	}

	handler, err := b.handler(reflect.TypeOf(queryFilter).String())

	if err != nil {
		return nil, err
	}

	return handler.Handle(ctx, queryFilter)
}

// New creates a new query bus.
func New() QueryBus {
	return &bus{
		handlers: make(handlers),
	}
}

func (b *bus) handler(cmdName string) (IQueryHandler, error) {
	if h, ok := b.handlers[cmdName]; ok {
		return h, nil
	}
	return nil, errors.New("handler not found for query filter")
}

func (b *bus) validate(cmd interface{}) error {
	value := reflect.ValueOf(cmd)

	if value.Kind() != reflect.Ptr || !value.IsNil() && value.Elem().Kind() != reflect.Struct {
		return errors.New("only pointer to commands are allowed")
	}
	return nil
}
