package decorators

import (
	"context"

	"golang.org/x/exp/slog"
)

type QueryHandler[Q any, V any] interface {
	Handle(context.Context, Q) (V, error)
}

func ApplyQueryDecorators[Q any, V any](
	handler QueryHandler[Q, V],
	logger *slog.Logger,
	metrics MetricsClient,
) QueryHandler[Q, V] {
	return NewQueryLoggingDecorator[Q, V](
		NewQueryMetricsDecorator(handler, metrics),
		logger,
	)
}
