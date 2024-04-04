package aircrafts_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"grid/pkg/models"
	"grid/pkg/repos/cache"
	"grid/pkg/transport/http/server"
	"grid/pkg/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
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
		"/v1.0/aircrafts",
		nil,
	)

	engine.ServeHTTP(recorder, req)
	require.Equal(t, http.StatusOK, recorder.Code)
}
