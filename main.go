package grid

import (
	"errors"
	"fmt"
	"time"
)

//
// Consts
//

const (
	ReconsilationWindow = 2
)

//
// Vars
//

var (
	// errors
	AppendEventAircraftIDEmptyError = errors.New("must include aircraft id!")
	AppendEventStationIDEmptyError  = errors.New("must include station id!")
	AppendEventTimestampEmptyError  = errors.New("must include timestamp!")

	// event buffer indexed by aircraftID
	Index map[string][]*Event = make(map[string][]*Event)
)

type Event struct {
	AircraftID string
	Latitude   float64
	Longitude  float64
	StationID  string
	Timestamp  *time.Time
}

func AppendEvent(payload *Event, reconsileEvents bool) error {
	aircraftID := payload.AircraftID

	// sanity check

	if aircraftID == "" {
		return AppendEventAircraftIDEmptyError
	}

	if payload.StationID == "" {
		return AppendEventStationIDEmptyError
	}

	if payload.Timestamp == nil {
		return AppendEventTimestampEmptyError
	}

	// append

	data := append(Index[payload.AircraftID], payload)

	// check to see if we need to reconsile

	if reconsileEvents == true && len(data) > ReconsilationWindow {
		events, err := ReconsileEvents(aircraftID)

		if err != nil {
			return err
		}

		// flush

		data = []*Event{}

		err = storeEvents(events)

		if err != nil {
			return err
		}

	}

	// ensure index update

	Index[payload.AircraftID] = data

	// success

	return nil
}

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

func isEventEqual(e1, e2 *Event) bool {
	// sanity check

	if e1 == nil || e2 == nil {
		return false
	}

	// compute
	// TODO: implmenet additional comp algos

	result :=
		isTimeEqual(e1.Timestamp, e2.Timestamp)

	// success

	return result
}

func isTimeEqual(t1, t2 *time.Time) bool {
	// sanity check

	if t1 == nil || t2 == nil {
		return false
	}

	// compute
	// TODO: implmenet theshold algo

	result := t1.Sub(*t2) == 0

	// success

	return result
}

func storeEvents(events []*Event) error {
	return nil
}
