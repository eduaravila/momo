package metrics

type NoOpt struct{}

// TODO: Implement the metrics client.
func (NoOpt) In(key string, val int) {}
