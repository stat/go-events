package tasks_test

import (
	"testing"
	"time"

	"events/pkg/models"
	"events/pkg/tasks/consumer"
	"events/pkg/tasks/producer"
	"events/pkg/utils"

	"github.com/stretchr/testify/assert"

	_ "events/testing"
)

func TestPipeline(t *testing.T) {
	aircraftID := "aircraft-01"
	stationID := "station-01"

	start := time.Now()

	tests := []struct {
		Event    *models.LocationEvent
		Expected error
	}{
		{
			Event: &models.LocationEvent{
				AircraftID: aircraftID,
				Latitude:   1.0,
				Longitude:  1.0,
				StationID:  stationID,
				Timestamp:  utils.Ref(start),
			},
			Expected: nil,
		},
		{
			Event: &models.LocationEvent{
				AircraftID: aircraftID,
				Latitude:   1.0,
				Longitude:  1.0,
				StationID:  stationID,
				Timestamp:  utils.Ref(start),
			},
			Expected: consumer.ProcessEventTimestampEqualError,
		},
		{
			Event: &models.LocationEvent{
				AircraftID: aircraftID,
				Latitude:   1.0,
				Longitude:  1.0,
				StationID:  stationID,
				Timestamp:  utils.Ref(start.Add(time.Second)),
			},
			Expected: consumer.ProcessEventTimestampAfterError,
		},
		{
			Event: &models.LocationEvent{
				AircraftID: aircraftID,
				Latitude:   1.0,
				Longitude:  1.0,
				StationID:  stationID,
				Timestamp:  utils.Ref(start.Add(time.Second * -1)),
			},
			Expected: consumer.ProcessEventTimestampBeforeError,
		},
		{
			Event: &models.LocationEvent{
				AircraftID: aircraftID,
				Latitude:   1.0,
				Longitude:  1.0,
				StationID:  stationID,
				Timestamp:  nil,
			},
			Expected: models.ADSBValidateTimestampError,
		},
	}

	// process

	for _, test := range tests {
		event := test.Event

		// consumer

		err := consumer.Process(event)

		// assert

		assert.Equal(t, test.Expected, err)

		if test.Expected != nil {
			continue
		}

		// producer

		err = producer.Process(event)

		assert.Equal(t, nil, err)
	}
}
