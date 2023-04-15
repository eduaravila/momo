package decorators

import (
	"context"
	"fmt"

	"golang.org/x/exp/slog"
)

type CommandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger *slog.Logger
}

func NewCommandLogginDecorator[C any](
	base CommandHandler[C],
	logger *slog.Logger,
) *CommandLoggingDecorator[C] {
	return &CommandLoggingDecorator[C]{base, logger}
}

func (d CommandLoggingDecorator[C]) Handle(
	ctx context.Context,
	command C,
) (err error) {
	loggerWithContext := d.logger.With(
		slog.String("name", generateActionName(command)),
		slog.String("body", fmt.Sprintf("%#v", command)))

	loggerWithContext.InfoCtx(ctx, "Executing command", command)

	defer func() {
		if err != nil {
			loggerWithContext.ErrorCtx(ctx, "Command goes Boom!", slog.Bool("success", false), err)
		} else {
			loggerWithContext.Info("Command  execuded successfully", slog.Bool("success", true))
		}
	}()

	return d.base.Handle(ctx, command)

}

type QueryLoggingDecorator[Q any, V any] struct {
	base   QueryHandler[Q, V]
	logger *slog.Logger
}

func (q QueryLoggingDecorator[Q, V]) Handle(
	ctx context.Context,
	query Q,
) (res V, err error) {
	loggerWithContext := q.logger.With(
		slog.String("name", generateActionName(query)),
		slog.String("body", fmt.Sprintf("%#v", query)))

	loggerWithContext.InfoCtx(ctx, "Executing query", query)

	defer func() {
		if err != nil {
			loggerWithContext.Error("Query goes Boom!", slog.Bool("success", false), err)
		} else {
			loggerWithContext.Info("Query  execuded successfully", slog.Bool("success", true))
		}
	}()

	return q.base.Handle(ctx, query)
}

func NewQueryLoggingDecorator[Q any, V any](
	base QueryHandler[Q, V], logger *slog.Logger) *QueryLoggingDecorator[Q, V] {
	return &QueryLoggingDecorator[Q, V]{base, logger}
}
