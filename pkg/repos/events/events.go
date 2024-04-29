package events

import (
	"errors"
	"events/pkg/env"
	"events/pkg/model"

	// "events/pkg/repo/backend"

	// "events/pkg/repos/events/backend"
	"events/pkg/repos/events/backends"
	"events/pkg/repos/events/provider"
)

type Repo[V model.Implementer] struct {
	provider.Provider[V]
}

// type V model.Implementer

var (
	// Backend provider.Provider[V]
	// Backend interface{}

	OptionsEmptyError    = errors.New("Index repository options are empty")
	OptionsEnvEmptyError = errors.New("Index repository env definitions are empty")
)

// var (
// Backend provider.Provider[model.Implementer]
// Backend backend.Provider
// )

type Options[V model.Implementer] struct {
	Env *env.Vars

	// Backend interface{}
	// Backend backend.Type
	Backend provider.Provider[V]
}

func (definitions *Options[V]) Validate() error {
	if definitions.Env == nil {
		return OptionsEnvEmptyError
	}

	return nil
}

func (implementation *Repo[V]) InitializeFn(options *Options[V]) (func(vars *env.Vars) error, error) {
	if options == nil {
		options = &Options[V]{}
	}

	if options.Backend == nil {
		options.Backend = &backends.Local[V]{}
	}

	backend := options.Backend

	fn := func(vars *env.Vars) error {
		return backend.Initialize(vars)
	}

	implementation.Provider = backend

	return fn, nil
}

// // func InitializeFn(options *Options) func(vars *env.Vars) error {
func InitializeFn[V model.Implementer](options *Options[V]) (func(vars *env.Vars) error, error) {
	// func InitializeFn[V model.Implementer](implementer provider.Provider[V]) (func(vars *env.Vars) error, error) {
	// backend, ok := options.Backend.(provider.Provider[model.Implementer])

	// if !ok {
	//   return nil, fmt.Errorf("could not create init function for events repo")
	// }

	backend := options.Backend

	fn := func(vars *env.Vars) error {
		return backend.Initialize(vars)
	}

	// Backend = backend

	return fn, nil
}

func Initialize[V model.Implementer](repo *Repo[V], options *Options[V]) error {
	// TODO: move this to default options merge
	if options == nil {
		return OptionsEmptyError
	}

	if err := options.Validate(); err != nil {
		return err
	}

	// TODO: move this to default options merge
	if options.Backend == nil {
		options.Backend = &backends.Local[V]{}
	}

	// backend

	backend := options.Backend

	// TODO: implement env vars within options
	if err := backend.Initialize(options.Env); err != nil {
		return err
	}

	repo.Provider = backend

	return nil
}

// func Initialize[T backend.Provider](vars *env.Vars) error {
//   var backend T

//   err := backend.Initialize(vars)

//   if err != nil {
//     return err
//   }

//   Backend = backend

//   return nil
// }

// func Initialize[T backend.Type, V model.Implementer](vars *env.Vars) error {
// func Initialize(vars *env.Vars) error {

//   provider

//   concrete, err := provider.Initialize(vars)

//   if err != nil {
//     return err
//   }

//   Backend = concrete

//   return nil
// }

// func Append[V model.Implementer](key string, v *V) error {
//   return Backend.Append(key, v)
// }

// func Del[V model.Implementer](key string) error {
//   return Backend.Del(key)
// }

// func DelAtIndex[V model.Implementer](key string, index int64) error {
//   return Backend.DelAtIndex(key, index)
// }

// func DelHead[V model.Implementer](key string) (*V, error) {
//   // TODO: implement me
//   return nil, nil
// }

// func DelTail[V model.Implementer](key string) (*V, error) {
//   // TODO: implement me
//   return nil, nil
// }

// func Get[V model.Implementer](key string) (*V, error) {
//   iface, err := Backend.Get(key)

//   return nil, nil
// }

// func GetAtIndex[V model.Implementer](key string) (*V, error) {
//   // TODO: implement me
//   return nil, nil
// }

// func GetHead[V model.Implementer](key string) (*V, error) {
//   // TODO: implement me
//   return nil, nil
// }

// func GetTail[V model.Implementer](key string) (*V, error) {
//   // TODO: implement me
//   return nil, nil
// }
