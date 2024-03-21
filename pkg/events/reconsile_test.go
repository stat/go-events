package events_test

import (
	"testing"
	"time"

	"grid/pkg/events"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//
// Vars
//

var (
	TestReconsileEventsAircraftID = "aircraftID"
	TestReconsileEventsStationID  = "stationID"
)

//
// Setup
//

func testReconsileEventsSetup() error {
	// clear

	if err := events.Clear(); err != nil {
		return err
	}

	// setup

	timestamp := time.Now()

	options := &events.AppendOptions{
		ReconsileEvents: false,
	}

	events.AppendWithOptions(&events.Event{
		AircraftID: TestReconsileEventsAircraftID,
		StationID:  TestReconsileEventsStationID,
		Latitude:   1,
		Longitude:  1,
		Timestamp:  &timestamp,
	}, options)

	events.AppendWithOptions(&events.Event{
		AircraftID: TestReconsileEventsAircraftID,
		StationID:  TestReconsileEventsStationID,
		Latitude:   1,
		Longitude:  1,
		Timestamp:  &timestamp,
	}, options)

	return nil
}

//
// Teardown
//

func testReconsileEventsTeardown() error {
	return events.Clear()
}

//
// Test
//

func TestReconsileEvents(t *testing.T) {
	// setup

	if err := testReconsileEventsSetup(); err != nil {
		panic(err)
	}

	// reconsile

	events, err := events.Reconsile(TestReconsileEventsAircraftID)

	require.Nil(t, err)
	assert.Equal(t, 1, len(events))

	// teardown

	if err := testReconsileEventsTeardown(); err != nil {
		panic(err)
	}

}
