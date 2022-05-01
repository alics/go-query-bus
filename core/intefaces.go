package core

import "context"

type IQueryFilter interface {
}
type IQueryHandler interface {
	Handle(ctx context.Context, filter interface{}) (result interface{}, err error)
}

type QueryBus interface {
	Register(interface{}, IQueryHandler) error
	GetHandlers() handlers
	Execute(context.Context, interface{}) (interface{}, error)
}

type ISortablePropertyCollection interface {
	GetDefault() string
}
