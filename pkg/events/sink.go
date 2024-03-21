package events

type SinkOptions interface{}

func Sink(events []*Event, options ...*SinkOptions) error {
	return nil
}
