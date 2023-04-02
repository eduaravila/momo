package decorators

import "context"

type QueryHandler[Q any, V any] interface {
	Handle(context.Context, Q) (V, error)
}

func ApplyQueryDecorators[Q any, V any](handler QueryHandler[Q, V]) QueryHandler[Q, V] {
	// TODO: Add, log, metrics and auth decorators
	return handler
}
