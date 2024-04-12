package provider

import (
	"grid/pkg/env"
	"grid/pkg/model"
)

// type Initializer = func[V any]

// type Initializer[V any] interface {
//   Initialize(vars *env.Vars) (Implementer[V], error)
// }

type Implementer[V model.Implementer] interface {
	GetLatest(key string) (*V, error)
	GetAircraftLocation(key string) (*V, error)
	GetAircraftsLocations() (map[string]*V, error)
	UpsertAircraftLocation(key string, v *V) error

	Initialize(vars *env.Vars) error
}

type Type string
