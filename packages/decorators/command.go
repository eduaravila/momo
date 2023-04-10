package decorators

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/exp/slog"
)

type Logger interface {
	Log(key string, val interface{})
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, command C) error
}

func ApplyCommandDecorators[C any](
	handler CommandHandler[C],
	logger *slog.Logger,
	metricsClient MetricsClient) CommandHandler[C] {
	return CommandLoggingDecorator[C]{
		base: CommandMetricsDecorator[C]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

func getCommandName[C any](command C) string {
	return strings.Split(fmt.Sprintf("%T", command), ".")[1]
}
