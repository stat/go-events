package backends

import (
	"errors"
	"fmt"
	"sync"

	"events/pkg/env"
	"events/pkg/model"
	//"github.com/hashicorp/go-set/v2"
)

type Local[V model.Implementer] struct {
	// type Local struct {
	Data *sync.Map
}

const (
	LocalDLQSuffix = "-dlq"
	// LocalType      = backend.Type("local")
)

var (
	LocalEventBackendKeyNotFound      = errors.New("Local event backend could not find key")
	LocalEventBackendIndexOutOfBounds = errors.New("Local event backend index out of bounds")
)

func (backend *Local[V]) Initialize(vars *env.Vars) error {
	// concrete := Local[V]{
	//   Data: &sync.Map{},
	// }

	backend.Data = &sync.Map{}

	return nil
}

func (backend *Local[V]) Append(key string, v *V) error {
	// var l []interface{}

	// backend.Get(key, &l)

	l, err := backend.Get(key)

	if err != nil {
		l = []*V{}
	}

	l = append(l, v)

	backend.Data.Store(key, l)

	return nil
}

func (backend *Local[V]) AppendDLQ(key string, v *V) error {
	dlqKey := fmt.Sprintf("%s%s", key, RedisDLQSuffix)
	if err := backend.Append(dlqKey, v); err != nil {
		return err
	}

	return nil
}

func (backend *Local[V]) Del(key string) error {
	// TODO: implement or remove me
	return nil
}

func (backend *Local[V]) DelAtIndex(key string, index int64) error {
	// TODO: implement or remove me
	return nil
}

func (backend *Local[V]) DelHead(key string) error {
	l, err := backend.Get(key)

	if err != nil {
		return err
	}

	l = l[:len(l)-1]

	backend.Data.Store(key, l)

	return nil
}

func (backend *Local[V]) DelTail(key string) error {
	// TODO: implement or remove me
	return nil
}

func (backend *Local[V]) Get(key string) ([]*V, error) {
	// func (backend Local[V]) Get(key string) ([]interface{}, error) {
	// func (backend Local[V]) Get(key string, v interface{}) error {
	iface, ok := backend.Data.Load(key)

	if !ok {
		return nil, LocalEventBackendKeyNotFound
	}

	l, ok := iface.([]*V)

	if !ok {
		return nil, fmt.Errorf("event local get cast error")
	}

	return l, nil
}

func (backend *Local[V]) GetAtIndex(key string, index int64) (*V, error) {
	// func (backend Local) GetAtIndex(key string, index int64, v interface{}) error {
	l, err := backend.Get(key)

	if err != nil {
		return nil, err
	}

	if int64(len(l)) < index {
		return nil, LocalEventBackendIndexOutOfBounds
	}

	// v = l[index]

	// return nil
	return l[index], nil
}

func (backend *Local[V]) GetHead(key string) (*V, error) {
	l, err := backend.Get(key)

	if err != nil {
		return nil, err
	}

	return backend.GetAtIndex(key, int64(len(l))-1)
}

func (backend *Local[V]) GetTail(key string) (*V, error) {
	return backend.GetAtIndex(key, 0)
}

func (backend *Local[V]) Set(key string, l []*V) error {
	// TODO: implement or remove me
	return nil
}
