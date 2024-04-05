package events_backends

import (
	"errors"
	"fmt"
	"sync"

	"grid/pkg/env"
	"grid/pkg/models"
	"grid/pkg/repos/events/provider"
	//"github.com/hashicorp/go-set/v2"
)

type Local struct {
	Data *sync.Map
}

const (
	LocalDLQSuffix = "-dlq"
)

var (
	LocalEventBackendKeyNotFound      = errors.New("Local event backend could not find key")
	LocalEventBackendIndexOutOfBounds = errors.New("Local event backend index out of bounds")
)

func (backend Local) Initialize(vars *env.Vars) (provider.Provider, error) {
	concrete := Local{
		Data: &sync.Map{},
	}

	backend.Data = &sync.Map{}

	return concrete, nil
}

func (backend Local) Append(key string, v *models.LocationEvent) error {
	l, err := backend.Get(key)

	if err != nil {
		l = []*models.LocationEvent{}
	}

	l = append(l, v)

	backend.Data.Store(key, l)

	return nil
}

func (backend Local) AppendDLQ(key string, v *models.LocationEvent) error {
	dlqKey := fmt.Sprintf("%s%s", key, RedisDLQSuffix)
	if err := backend.Append(dlqKey, v); err != nil {
		return err
	}

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
	l, err := backend.Get(key)

	if err != nil {
		return err
	}

	l = l[:len(l)-1]

	backend.Data.Store(key, l)

	return nil
}

func (backend Local) DelTail(key string) error {
	// TODO: implement or remove me
	return nil
}
func (backend Local) Get(key string) ([]*models.LocationEvent, error) {
	iface, ok := backend.Data.Load(key)

	if !ok {
		return nil, LocalEventBackendKeyNotFound
	}

	l, ok := iface.([]*models.LocationEvent)

	if !ok {
		return nil, fmt.Errorf("event local get cast error")
	}

	return l, nil
}

func (backend Local) GetAtIndex(key string, index int64) (*models.LocationEvent, error) {
	l, err := backend.Get(key)

	if err != nil {
		return nil, err
	}

	if int64(len(l)) < index {
		return nil, LocalEventBackendIndexOutOfBounds
	}

	return l[index], nil
}

func (backend Local) GetHead(key string) (*models.LocationEvent, error) {
	l, err := backend.Get(key)

	if err != nil {
		return nil, err
	}

	return backend.GetAtIndex(key, int64(len(l))-1)
}

func (backend Local) GetTail(key string) (*models.LocationEvent, error) {
	return backend.GetAtIndex(key, 0)
}
