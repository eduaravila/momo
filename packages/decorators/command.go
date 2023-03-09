package decorators

import "context"

type Logger interface {
	Log(key string, val interface{})
}

type CommandHandler[C any] interface {
	Handle(ctx context.Context, command C) error
}

func ApplyCommandDecorators[C any](handler CommandHandler[C], logger Logger, metricsClient MetricsClient) CommandHandler[C] {
	return CommandLoggingDecorator[C]{
		base: CommandMetricsDecorator[C]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}
