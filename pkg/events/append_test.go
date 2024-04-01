package events_test

import (
	"testing"
	"time"

	"grid/pkg/events"
	"grid/pkg/utils"

	"github.com/stretchr/testify/assert"
)

//
// Test
//

func TestConsume(t *testing.T) {
	// clear

	// require.Nil(t, events.Clear())

	// tests

	tests := []struct {
		Event    *events.Location
		Expected error
	}{
		{
			Expected: events.AppendAircraftIDEmptyError,
			Event: &events.Location{
				AircraftID: "",
				Latitude:   0,
				Longitude:  0,
				StationID:  "stationID",
				Timestamp:  utils.Ref(time.Now()),
			},
		},
		{
			Expected: nil,
			Event: &events.Location{
				AircraftID: "aircraftID",
				Latitude:   0,
				Longitude:  0,
				StationID:  "stationID",
				Timestamp:  utils.Ref(time.Now()),
			},
		},
	}

	for _, test := range tests {
		payload := test.Event

		result := events.Append(payload)
		assert.Equal(t, test.Expected, result)

		// if not success, then contiue
		if test.Expected != nil {
			continue
		}

		data := events.Index[payload.AircraftID]
		assert.NotEmpty(t, data)
	}
}
