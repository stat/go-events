package aircrafts_events_test

import (
	"bytes"
	"encoding/json"
	"grid/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"grid/pkg/http/server"
	"grid/pkg/utils"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"

	_ "grid/testing"
)

func TestCreate(t *testing.T) {
	engine, err := server.Engine()
	require.Nil(t, err)

	event := &models.LocationEvent{
		AircraftID: "AircraftID",
		Latitude:   37.6,
		Longitude:  -95.665,
		StationID:  "StationID",
		Timestamp:  utils.Ref(time.Now()),
	}

	payload := utils.Must(json.Marshal(event))
	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"POST",
		"/v1.0/aircrafts/aircraft_id/events",
		bytes.NewBuffer(payload),
	)

	engine.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}
