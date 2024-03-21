package events

func New() (map[string][]*Event, error) {
	return make(map[string][]*Event), nil
}
