package aircrafts_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"grid/pkg/models"
	"grid/pkg/repos/cache"
	"grid/pkg/transport/http/respond"
	"grid/pkg/transport/http/server"
	"grid/pkg/utils"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "grid/testing"
)

func TestGet(t *testing.T) {
	// http engine

	engine, err := server.Engine()
	require.Nil(t, err)

	// upsert location

	aircraftID := uuid.NewString()
	timestamp := time.Now()

	event := &models.LocationEvent{
		AircraftID: aircraftID,
		Latitude:   37.6,
		Longitude:  -95.665,
		StationID:  "StationID",
		Timestamp:  utils.Ref(timestamp),
	}

	err = cache.Backend.UpsertAircraftLocation(aircraftID, event)
	require.Nil(t, err)

	// get location

	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("/v1.0/aircrafts/%s", aircraftID),
		nil,
	)

	engine.ServeHTTP(recorder, req)
	require.Equal(t, http.StatusOK, recorder.Code)

	// response

	responseString := recorder.Body.String()

	// unmarshal

	response := &respond.Response[models.LocationEvent]{}

	err = json.Unmarshal([]byte(responseString), response)
	require.Nil(t, err)

	responseTimestamp := response.Data.Timestamp
	event.Timestamp = nil
	response.Data.Timestamp = nil

	// assert

	assert.EqualValues(t, event, response.Data)
	assert.WithinDuration(t, timestamp, *responseTimestamp, 0)
}
