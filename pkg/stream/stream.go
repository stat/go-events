package stream

import (
	"errors"
	"events/pkg/env"
	"events/pkg/model"
	"events/pkg/repos/cache"
	"events/pkg/repos/events"
	"events/workers"
)

var (

	//
	// Errors
	//

	// Options Errors

	OptionsCacheDefinitionsEmptyError    = errors.New("Stream cache options are empty")
	OptionsCacheEnvDefinitionsEmptyError = errors.New("Stream env options are empty")
	OptionsIndexDefinitionsEmptyError    = errors.New("Stream index options are empty")
	OptionsEmptyError                    = errors.New("Stream options are empty")
)

type Handler[T model.Implementer] func(payload *T) error

type Options[CacheModel, EventModel model.Implementer] struct {
	Env *env.Vars

	Cache *cache.Options[CacheModel]
	Index *events.Options[EventModel]

	Workers *workers.Options[CacheModel, EventModel]
}

func (definitions *Options[_, _]) Validate() error {
	if definitions.Env == nil {
		return OptionsCacheEnvDefinitionsEmptyError
	}
	if definitions.Cache == nil {
		return OptionsCacheDefinitionsEmptyError
	}

	if definitions.Index == nil {
		return OptionsIndexDefinitionsEmptyError
	}
	return nil
}

type Stream[CacheModel, EventModel model.Implementer] struct {
	Cache *cache.Repo[CacheModel]
	Index *events.Repo[EventModel]

	Handler Handler[EventModel]
	Workers *workers.Server[CacheModel, EventModel]

	Options *Options[CacheModel, EventModel]
}

func (ctx *Stream[_, E]) Handle(payload *E) error {
	return nil
}

func New[C, E model.Implementer](env *env.Vars) (*Stream[C, E], error) {
	return nil, nil
}

func NewWithOptions[C, E model.Implementer](options *Options[C, E]) (*Stream[C, E], error) {
	// ensure

	if options == nil {
		return nil, OptionsEmptyError
	}

	// validate

	if err := options.Validate(); err != nil {
		return nil, err
	}

	// cache

	cacheRepo, err := NewCacheWithOptions(options)

	if err != nil {
		return nil, err
	}

	// index

	indexRepo, err := NewIndexWithOptions(options)

	if err != nil {
		return nil, err
	}

	// stream

	s := &Stream[C, E]{
		Cache:   cacheRepo,
		Index:   indexRepo,
		Options: options,
	}

	// success

	return s, nil
}

func NewCacheWithOptions[C, E model.Implementer](options *Options[C, E]) (*cache.Repo[C], error) {
	repo := &cache.Repo[C]{}

	if err := cache.Initialize(repo, options.Cache); err != nil {
		return nil, err
	}

	return repo, nil
}

func NewIndexWithOptions[C, E model.Implementer](options *Options[C, E]) (*events.Repo[E], error) {
	repo := &events.Repo[E]{}

	if err := events.Initialize(repo, options.Index); err != nil {
		return nil, err
	}

	return repo, nil
}

// TODO: rename this to NewWithOptions and move this to pkg/repos/repo.go
// func NewRepoWithOptions[T any[V], V model.Implementer](options any[V]) (*T[V], error) {
//   //TODO: implement me
//   return nil, nil
// }
