package backends_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"grid/pkg/models"
	events_backends "grid/pkg/repos/events/backends"
	"grid/pkg/utils"

	_ "grid/testing"
)

func TestLocal(t *testing.T) {
	// setup

	provider := events_backends.Local{}
	backend, err := provider.Initialize(nil)

	// require.Nil(t, err)

	key := uuid.NewString()

	event := &models.LocationEvent{
		AircraftID: key,
		Timestamp:  utils.Ref(time.Now()),
	}

	// l := []*models.LocationEvent{event}
	// tests

	// backend.Data.Store(key, l)

	err = backend.Append(key, event)
	require.Nil(t, err)

	// iface, ok := backend.Data.Load(key)

	// fmt.Println(iface, ok, key)

	// result, err := backend.Get(key)
	// require.Nil(t, err)
	// assert.EqualValues(t, event, result)

	result, err := backend.GetHead(key)
	require.Nil(t, err)
	assert.EqualValues(t, event, result)
	// result, ok := backend.Data.Load(key)
	// require.True(t, ok)
	// assert.EqualValues(t, l, result)

	event = &models.LocationEvent{
		AircraftID: key,
		Timestamp:  utils.Ref(time.Now()),
	}

	err = backend.Append(key, event)
	require.Nil(t, err)

	result, err = backend.GetHead(key)
	require.Nil(t, err)
	assert.EqualValues(t, event, result)

	// teardown

}
