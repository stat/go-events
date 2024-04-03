package provider

import (
	"grid/pkg/env"
	"grid/pkg/models"
)

type Provider interface {
	GetAircraftLocation() (*models.LocationEvent, error)
	GetAircraftsLocations() (map[string]*models.LocationEvent, error)
	UpsertAircraftLocation(key string, v *models.LocationEvent) error

	Initialize(vars *env.Vars) (Provider, error)
}
