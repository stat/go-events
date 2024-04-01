package provider

import (
	"grid/pkg/env"
	"grid/pkg/models"
)

type Provider interface {
	GetAircraftLocation() (*models.ADSB, error)
	GetAircraftsLocations(key string) (map[string]*models.ADSB, error)
	UpsertAircraftLocation(key string, v interface{}) error

	Initialize(vars *env.Vars) (Provider, error)
}
