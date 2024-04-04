package aircrafts_events_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"grid/pkg/models"
	"grid/pkg/repos/cache"
	"grid/pkg/transport/http/server"
	"grid/pkg/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "grid/testing"
)

func TestCreate(t *testing.T) {
	engine, err := server.Engine()
	require.Nil(t, err)

	aircraftID := uuid.NewString()
	timestamp := time.Now()

	event := &models.LocationEvent{
		AircraftID: aircraftID,
		Latitude:   37.6,
		Longitude:  -95.665,
		StationID:  "StationID",
		Timestamp:  utils.Ref(timestamp),
	}

	// create

	payload := utils.Must(json.Marshal(event))
	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("/v1.0/aircrafts/%s/events", aircraftID),
		bytes.NewBuffer(payload),
	)

	engine.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	// sleep

	attempts := 0
	attemptsMax := 5

	var location *models.LocationEvent

	for {
		time.Sleep(time.Second * 1)

		attempts++
		location, err = cache.Backend.GetAircraftLocation(aircraftID)

		if err != nil && attempts < attemptsMax {
			continue
		}

		break
	}

	require.NotEqual(t, attemptsMax, attempts)

	// check

	locationTimestamp := location.Timestamp

	event.Timestamp = nil
	location.Timestamp = nil

	assert.Equal(t, event, location)
	assert.WithinDuration(t, timestamp, *locationTimestamp, 0)
}
