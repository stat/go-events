package events

type StoreOptions interface{}

func StoreEvents(events []*Event, options ...*StoreOptions) error {
	return nil
}
