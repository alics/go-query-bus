package core

import (
	"context"
	"reflect"
	"testing"
)

type QueryFilter struct {
	id int
}

type QueryHandler struct {
}

func (q QueryHandler) Handle(ctx context.Context, filter interface{}) (result interface{}, err error) {
	return 1, nil
}

func NewAQueryHandler() IQueryHandler {
	return &QueryHandler{}
}

func TestBus_Register(t *testing.T) {
	bus := New()

	if err := bus.Register(&QueryFilter{}, NewAQueryHandler()); err != nil {
		t.Errorf("Failed to register query handler: %v", err)
	}

	if len(bus.GetHandlers()) != 1 {
		t.Errorf("Expected to have 1 query handler, got %d", len(bus.GetHandlers()))
	}

	for _, handler := range bus.GetHandlers() {
		if reflect.ValueOf(handler) != reflect.ValueOf(NewAQueryHandler()) {
			t.Error("Registered query handler is different from the expected one")
		}
	}

	// duplicated
	if err := bus.Register(&QueryFilter{}, NewAQueryHandler()); err == nil {
		t.Error("Bus must not accept duplicated query handler")
	}
}

func TestBus_Execute(t *testing.T) {
	bus := New()

	if err := bus.Register(&QueryFilter{}, NewAQueryHandler()); err != nil {
		t.Errorf("Failed to register query handler: %v", err)
	}

	c := &QueryFilter{}
	expected := 1

	result, err := bus.Execute(context.Background(), c)
	if err != nil {
		t.Errorf("Failed to execute query handler: %v", err)
	}

	if result != expected {
		t.Errorf("Execution result is wrong. Expected %d, got %d", expected, result)
	}
}
