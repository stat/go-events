package provider

import (
	"grid/pkg/env"
	"grid/pkg/models"
)

type Provider interface {
	Append(key string, v interface{}) error
	Del(key string) error
	DelAtIndex(key string, index int64) error
	DelHead(key string) error
	DelTail(key string) error
	Get(key string) (*models.ADSB, error)
	GetAtIndex(key string, index int64) (*models.ADSB, error)
	GetHead(key string) (*models.ADSB, error)
	GetTail(key string) (*models.ADSB, error)

	Initialize(vars *env.Vars) (Provider, error)
}
