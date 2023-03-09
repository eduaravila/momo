package decorators

import "context"

type CommandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger Logger
}

func (d CommandLoggingDecorator[C]) Handle(ctx context.Context, command C) error {
	d.logger.Log("command", command)
	return d.base.Handle(ctx, command)
}
