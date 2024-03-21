package events

import "fmt"

func ReconsileEvents(aircraftID string) ([]*Event, error) {
	// get events by aircraftID

	data := Index[aircraftID]

	// ensure data

	if data == nil {
		return nil, fmt.Errorf("no data available for %s", aircraftID)
	}

	// initialize reconsiled events

	reconsiledEvents := []*Event{}

	var event *Event = nil
	for _, item := range data {
		// determine if equal, then resume
		if isEventEqual(event, item) {
			continue
		}

		// append

		reconsiledEvents = append(reconsiledEvents, item)

		// rotate

		event = item
	}

	// success

	return reconsiledEvents, nil
}
