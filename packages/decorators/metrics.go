package decorators

import (
	"context"
	"time"
)

type MetricsClient interface {
	In(key string, val int)
}

type CommandMetricsDecorator[C any] struct {
	base   CommandHandler[C]
	client MetricsClient
}

// Handle is a decorator that measures the time it takes to execute a command
func (d CommandMetricsDecorator[C]) Handle(ctx context.Context, command C) error {

	start := time.Now()
	end := time.Since(start)

	d.client.In("command", int(end.Seconds()))
	return d.base.Handle(ctx, command)
}
