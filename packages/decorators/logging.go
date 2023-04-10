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

func NewCommandLogginDecorator[C any](base CommandHandler[C], logger *slog.Logger) *CommandLoggingDecorator[C] {
	return &CommandLoggingDecorator[C]{base, logger}
}

func (d CommandLoggingDecorator[C]) Handle(ctx context.Context, command C) (err error) {
	loggerWithContext := d.logger.With(
		slog.String("name", getCommandName(command)),
		slog.String("body", fmt.Sprintf("%#v", command)))

	loggerWithContext.InfoCtx(ctx, "command", command)

	defer func() {
		if err != nil {
			loggerWithContext.ErrorCtx(ctx, "command", slog.Bool("success", false), err)
		} else {
			loggerWithContext.Info("command", slog.Bool("success", true))
		}
	}()

	return d.base.Handle(ctx, command)

}
