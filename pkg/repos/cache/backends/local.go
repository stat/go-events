package backends

import (
	"sync"

	"grid/pkg/env"
	"grid/pkg/models"
	"grid/pkg/repos/cache/provider"
)

type Local struct {
	// AircraftsLocations map[string]*models.LocationEvent
	AircraftsLocations *sync.Map
}

func (backend Local) Initialize(vars *env.Vars) (provider.Provider, error) {
	concrete := Local{
		// AircraftsLocations: map[string]*models.LocationEvent{},
		AircraftsLocations: &sync.Map{},
	}

	return concrete, nil
}

func (backend Local) GetAircraftLocation() (*models.LocationEvent, error) {
	// TODO: implement or remove me
	return nil, nil
}

func (backend Local) GetAircraftsLocations(key string) (map[string]*models.LocationEvent, error) {
	// TODO: implement or remove me
	return nil, nil
}

func (backend Local) UpsertAircraftLocation(key string, v *models.LocationEvent) error {
	// backend.AircraftsLocations[key] = v

	backend.AircraftsLocations.Store(key, v)
	return nil
}
