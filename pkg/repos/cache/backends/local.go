package backends

import (
	"errors"
	"sync"

	"grid/pkg/env"
	"grid/pkg/models"
	"grid/pkg/repos/cache/provider"
)

type Local struct {
	// AircraftsLocations map[string]*models.LocationEvent
	AircraftsLocations *sync.Map
}

var (
	LocalCacheKeyNotFoundError      = errors.New("Local cache could not find key")
	LocalCacheKeyValueCoersionError = errors.New("Local cache could not cast value")
)

func (backend Local) Initialize(vars *env.Vars) (provider.Provider, error) {
	concrete := Local{
		// AircraftsLocations: map[string]*models.LocationEvent{},
		AircraftsLocations: &sync.Map{},
	}

	return concrete, nil
}

func (backend Local) GetAircraftLocation(key string) (*models.LocationEvent, error) {
	iface, ok := backend.AircraftsLocations.Load(key)

	if !ok {
		return nil, LocalCacheKeyNotFoundError
	}

	location, ok := iface.(*models.LocationEvent)

	if !ok {
		return nil, LocalCacheKeyValueCoersionError
	}

	return location, nil
}

func (backend Local) GetAircraftsLocations() (map[string]*models.LocationEvent, error) {
	// TODO: implement or remove me
	return nil, nil
}

func (backend Local) UpsertAircraftLocation(key string, v *models.LocationEvent) error {
	// backend.AircraftsLocations[key] = v

	backend.AircraftsLocations.Store(key, v)
	return nil
}
