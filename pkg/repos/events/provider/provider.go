package provider

import (
	"events/pkg/env"
	"events/pkg/model"
)

type Provider[V model.Implementer] interface {
	Append(key string, v *V) error
	// AppendDLQ(key string, v *V) error
	Del(key string) error
	DelAtIndex(key string, index int64) error
	DelHead(key string) error
	DelTail(key string) error
	Get(key string) ([]*V, error)
	GetAtIndex(key string, index int64) (*V, error)
	GetHead(key string) (*V, error)
	GetTail(key string) (*V, error)
	// Set(key string, l []*V) error

	Initialize(vars *env.Vars) error
}
