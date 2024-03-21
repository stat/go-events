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
	TestReconcileEventsAircraftID = "aircraftID"
	TestReconcileEventsStationID  = "stationID"
)

//
// Setup
//

func testReconcileEventsSetup() error {
	// clear

	if err := events.Clear(); err != nil {
		return err
	}

	// setup

	timestamp := time.Now()

	options := &events.AppendOptions{
		ReconcileEvents: false,
	}

	events.AppendWithOptions(&events.Event{
		AircraftID: TestReconcileEventsAircraftID,
		StationID:  TestReconcileEventsStationID,
		Latitude:   1,
		Longitude:  1,
		Timestamp:  &timestamp,
	}, options)

	events.AppendWithOptions(&events.Event{
		AircraftID: TestReconcileEventsAircraftID,
		StationID:  TestReconcileEventsStationID,
		Latitude:   1,
		Longitude:  1,
		Timestamp:  &timestamp,
	}, options)

	return nil
}

//
// Teardown
//

func testReconcileEventsTeardown() error {
	return events.Clear()
}

//
// Test
//

func TestReconcileEvents(t *testing.T) {
	// setup

	if err := testReconcileEventsSetup(); err != nil {
		panic(err)
	}

	// reconcile

	events, err := events.Reconcile(TestReconcileEventsAircraftID)

	require.Nil(t, err)
	assert.Equal(t, 1, len(events))

	// teardown

	if err := testReconcileEventsTeardown(); err != nil {
		panic(err)
	}

}
