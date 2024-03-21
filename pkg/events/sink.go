package events

type SinkOptions interface{}

func SinkEvents(events []*Event, options ...*SinkOptions) error {
	return nil
}
