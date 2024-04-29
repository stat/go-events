package backends

import (
	"errors"
	"sync"

	"events/pkg/env"
	"events/pkg/model"
)

type Local[V model.Implementer] struct {
	// Data map[string]*models.LocationEvent
	Data *sync.Map
}

var (
	// Local = provider.Type("local")

	LocalCacheKeyNotFoundError      = errors.New("Local cache could not find key")
	LocalCacheKeyValueCoersionError = errors.New("Local cache could not cast value")
)

// func InitializeLocal[T provider.Implementer[V], V model.Implementer](vars *env.Vars) (provider.Implementer[V], error) {

//   return nil, nil
// }

// func (backend Local) Initialize(vars *env.Vars) (provider.Implementer, error) {
func (backend *Local[V]) Initialize(vars *env.Vars) error {
	// concrete := Local{
	//   // AircraftsLocations: map[string]*models.LocationEvent{},
	//   Latests: &sync.Map{},
	// }

	backend.Data = &sync.Map{}

	return nil
}

func (backend *Local[V]) GetLatest(key string) (*V, error) {
	iface, ok := backend.Data.Load(key)

	if !ok {
		return nil, LocalCacheKeyNotFoundError
	}

	event, ok := iface.(*V)

	if !ok {
		return nil, LocalCacheKeyValueCoersionError
	}

	return event, nil
	// return iface, nil
}

func (backend *Local[V]) GetAircraftLocation(key string) (*V, error) {
	iface, ok := backend.Data.Load(key)

	if !ok {
		return nil, LocalCacheKeyNotFoundError
	}

	location, ok := iface.(*V)

	if !ok {
		return nil, LocalCacheKeyValueCoersionError
	}

	return location, nil
}

func (backend *Local[V]) GetAircraftsLocations() (map[string]*V, error) {
	// TODO: implement or remove me
	return nil, nil
}

func (backend *Local[V]) UpsertAircraftLocation(key string, v *V) error {
	// backend.AircraftsLocations[key] = v

	backend.Data.Store(key, v)

	return nil
}
