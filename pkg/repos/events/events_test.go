package events_test

import (
	"testing"

	"events/pkg/env"
	"events/pkg/models"
	"events/pkg/repos/events/backends"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	_ "events/testing"
)

func TestEventsRepo(t *testing.T) {
	vars, err := env.Load()
	assert.Nil(t, err)

	backend := backends.Local[models.LocationEvent]{}
	backend.Initialize(vars)

	key := uuid.NewString()

	// append

	value := &models.LocationEvent{AircraftID: "hi"}
	err = backend.Append(key, value)
	assert.Nil(t, err)

	// get

	values, err := backend.Get(key)
	assert.Nil(t, err)

	// get cache

	// cached, err :=

	// assert

	assert.EqualValues(t, value, values[len(values)-1])
	// assert.EqualValues(t,
}
