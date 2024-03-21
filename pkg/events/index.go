package events

var (
	// event buffer indexed by aircraftID
	Index map[string][]*Event = make(map[string][]*Event)
)
