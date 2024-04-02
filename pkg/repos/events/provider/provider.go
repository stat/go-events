package provider

import (
	"grid/pkg/env"
	"grid/pkg/models"
)

type Provider interface {
	Append(key string, v *models.LocationEvent) error
	Del(key string) error
	DelAtIndex(key string, index int64) error
	DelHead(key string) error
	DelTail(key string) error
	Get(key string) ([]*models.LocationEvent, error)
	GetAtIndex(key string, index int64) (*models.LocationEvent, error)
	GetHead(key string) (*models.LocationEvent, error)
	GetTail(key string) (*models.LocationEvent, error)
	// Set(key string, l []*models.LocationEvent) error

	Initialize(vars *env.Vars) (Provider, error)
}
