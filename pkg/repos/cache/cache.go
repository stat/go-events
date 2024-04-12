package cache

import (
	"grid/pkg/env"
	"grid/pkg/model"
	"grid/pkg/repos/cache/backends"
	"grid/pkg/repos/cache/provider"
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

type Repo[V model.Implementer] struct {
	provider.Implementer[V]
}

type Options[V model.Implementer] struct {
	Backend provider.Implementer[V]
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

	implementation.Implementer = backend

	return fn, nil
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
