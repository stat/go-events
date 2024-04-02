package events_backends

import (
	"errors"
	"grid/pkg/env"
	"grid/pkg/models"
	"grid/pkg/repos/events/provider"
)

type Local struct {
	Data map[string][]*models.LocationEvent
}

var (
	LocalEventBackendKeyNotFound      = errors.New("Local event backend could not find key")
	LocalEventBackendIndexOutOfBounds = errors.New("Local event backend index out of bounds")
)

func (backend Local) Initialize(vars *env.Vars) (provider.Provider, error) {
	concrete := Local{}

	return concrete, nil
}

func (backend Local) Append(key string, v *models.LocationEvent) error {
	l, ok := backend.Data[key]

	if !ok {
		l = []*models.LocationEvent{}
	}

	l = append(l, v)

	backend.Data[key] = l

	return nil
}

func (backend Local) Del(key string) error {
	// TODO: implement or remove me
	return nil
}

func (backend Local) DelAtIndex(key string, index int64) error {
	// TODO: implement or remove me
	return nil
}

func (backend Local) DelHead(key string) error {
	l, ok := backend.Data[key]

	if !ok {
		return nil
	}

	l = l[:len(l)-1]

	backend.Data[key] = l

	return nil
}

func (backend Local) DelTail(key string) error {
	// TODO: implement or remove me
	return nil
}
func (backend Local) Get(key string) (*models.LocationEvent, error) {
	// TODO: implement or remove me
	return nil, nil
}

func (backend Local) GetAtIndex(key string, index int64) (*models.LocationEvent, error) {
	l, ok := backend.Data[key]

	if !ok {
		return nil, LocalEventBackendKeyNotFound
	}

	if int64(len(l)) < index {
		return nil, LocalEventBackendIndexOutOfBounds
	}

	return l[index], nil
}

func (backend Local) GetHead(key string) (*models.LocationEvent, error) {
	l, ok := backend.Data[key]

	if !ok {
		return nil, LocalEventBackendKeyNotFound
	}

	return backend.GetAtIndex(key, int64(len(l))-1)
}

func (backend Local) GetTail(key string) (*models.LocationEvent, error) {
	return backend.GetAtIndex(key, 0)
}
