package cache

import (
	"errors"
	"events/pkg/env"
	"events/pkg/model"
	"events/pkg/repos/cache/backends"
	"events/pkg/repos/cache/provider"
)

// type Item[V any] struct {
//   Value *V `json:"value"`
// }

// type Provider[T provider.Implementer[V], V any] struct {
//   Concrete T
// }

// type Options[T provider.Implementer[V], V any] struct {
// type Options[T provider.Initializer, V model.Implementer] struct {
//
//	type Options[T any, V any] struct {
//	  Provider *T
//	  Model    *V
//	}

var (
	//
	// Errors
	//

	OptionsEmptyError    = errors.New("Cache repository options are empty")
	OptionsEnvEmptyError = errors.New("Cache repository env definitions are empty")
)

//
// Options
//

type Options[V model.Implementer] struct {
	Env     *env.Vars
	Backend provider.Implementer[V]
}

func (definitions *Options[V]) Validate() error {
	if definitions.Env == nil {
		return OptionsEnvEmptyError
	}

	return nil
}

//
// Repository
//

type Repo[V model.Implementer] struct {
	provider.Implementer[V]
}

func (implementation *Repo[V]) InitializeFn(options *Options[V]) (func(vars *env.Vars) error, error) {
	// if options == nil {
	//   options = &Options[V]{}
	// }

	// if options.Backend == nil {
	//   options.Backend = &backends.Local[V]{}
	// }

	// backend := options.Backend

	fn := func(vars *env.Vars) error {
		// return backend.Initialize(vars)
		return Initialize(implementation, options)
	}

	// implementation.Implementer = backend

	return fn, nil
}

//
// Static Fns
//

func DefaultOptions[V model.Implementer]() (*Options[V], error) {
	// TODO: implement me
	return nil, nil
}

// TODO: refine generics and move into pkg/repo/repo.go
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

	repo.Implementer = backend

	return nil
}

// type V = model.Implementer

// var (
//   // Backend provider.Implementer[model.Implementer]
//   Backend interface{}
// )

// func Initialize[T provider.Implementer[models.CacheEntry[V]], V model.Implementer](vars *env.Vars) error {
// func Initialize[T provider.Implementer](vars *env.Vars) error {
//   var provider T

//   concrete, err := provider.Initialize(vars)

//   if err != nil {
//     return err
//   }

//   Backend = concrete

//   return nil
// }

// func InitializeFnWithOptions(options *Options) func(*env.Vars) error {
// func InitializeFn[T provider.Implementer, V model.Implementer]() func(*env.Vars) error {
// func InitializeFn[T provider.Implementer, V model.Implementer](*env.Vars) error {
//   return nil
// }

func InitializeFn[V model.Implementer](options *Options[V]) (func(vars *env.Vars) error, error) {
	// func InitializeFn[V model.Implementer](implementer provider.Provider[V]) (func(vars *env.Vars) error, error) {
	return nil, nil
}
