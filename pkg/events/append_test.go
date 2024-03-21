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

func TestAppendEvent(t *testing.T) {
	tests := []struct {
		Event    *events.Event
		Expected error
	}{
		{
			Expected: events.AppendEventAircraftIDEmptyError,
			Event: &events.Event{
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

		result := events.AppendEvent(payload)
		assert.Equal(t, test.Expected, result)

		// if not success, then contiue
		if test.Expected != nil {
			continue
		}

		data := events.Index[payload.AircraftID]
		assert.NotEmpty(t, data)
	}
}
