package grid_test

import (
	"os"
	"testing"
	"time"

	"grid"
	"grid/pkg/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	setup()
	result := m.Run()
	teardown()

	os.Exit(result)
}

// Setup

func setup() {
}

// Teardown

func teardown() {
}

//
// Tests
//

func TestAppendEvent(t *testing.T) {
	tests := []struct {
		Event    *grid.Event
		Expected error
	}{
		{
			Expected: grid.AppendEventAircraftIDEmptyError,
			Event: &grid.Event{
				AircraftID: "",
				Latitude:   0,
				Longitude:  0,
				StationID:  "stationID",
				Timestamp:  utils.Ref(time.Now()),
			},
		},
	}

	for _, test := range tests {
		payload := test.Event

		result := grid.AppendEvent(payload, true)
		assert.Equal(t, test.Expected, result)

		if test.Expected != nil {
			continue
		}

		// assumptions:
		// aircraftID is always present

		data := grid.Index[payload.AircraftID]
		assert.NotEmpty(t, data)
	}
}

func TestReconsileEvents(t *testing.T) {
	aircraftID := "aircraftID"
	stationID := "stationID"

	timestamp := time.Now()

	grid.AppendEvent(&grid.Event{
		AircraftID: aircraftID,
		Latitude:   1,
		Longitude:  1,
		StationID:  stationID,
		Timestamp:  &timestamp,
	}, false)

	grid.AppendEvent(&grid.Event{
		AircraftID: aircraftID,
		Latitude:   1,
		Longitude:  1,
		StationID:  stationID,
		Timestamp:  &timestamp,
	}, false)

	events, err := grid.ReconsileEvents(aircraftID)

	require.Nil(t, err)
	assert.Equal(t, 1, len(events))
}
