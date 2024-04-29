package app

import (
	"grid/config"
	"grid/pkg/db/queue"
	"grid/pkg/env"
	"grid/pkg/lifecycle"
	"grid/pkg/repos/cache"
	"grid/pkg/repos/events"
	"grid/pkg/transport/http/server"
	"grid/pkg/transport/ws/socket"
	"grid/pkg/utils"
	"grid/workers"

	// "grid/pkg/db/postgres"

	cache_backends "grid/pkg/repos/cache/backends"
	events_backends "grid/pkg/repos/events/backends"
)

// TODO: move type aliases into config

// TODO: scope these vars to Stream
// TODO: think about how to implement multiple streams

var (
	Initializers = lifecycle.Fns[env.Vars]{
		// db client connections

		// postgres.Initialize,
		queue.Initialize,

		// repos

		// cache.InitializeFn(
		//   &cache_backends.Local{},
		//   &models.LocationEvent{},
		// ),

		// events.InitializeFn(&events.Options{
		//   Provider: events_backends.LocalType,
		// }),

		// utils.Must(
		//   events.InitializeFn(&events.Options{
		//     Backend: &events_backends.Local[models.LocationEvent]{},
		//   }),
		// ),

		// events.Initialize[events_backends.Local[models.LocationEvent]],
		// utils.Must(
		//   events.InitializeFn(EventsOptions),
		//   // events.InitializeFn(
		//   //   &backends.Local[models.LocationEvent]{},
		//   // ),
		// ),

		utils.Must(
			Cache.InitializeFn(CacheOptions),
		),

		utils.Must(
			Index.InitializeFn(IndexOptions),
		),

		// events.Backend.InitializeFnWithOptions(events.Options{
		//   KeyFn: func(event *models.LocationEvent) (key string, err error) {
		//     return "", nil
		//   },
		// }),

		// transport

		server.Initialize,

		utils.Must(
			Socket.InitializeFn(SocketOptions),
		),

		// asynq server

		// workers.Initialize,
		utils.Must(
			Workers.InitializeFn(WorkersOptions),
		),
	}

	Index        = &events.Repo[config.EventModel]{}
	IndexOptions = &events.Options[config.EventModel]{
		Backend: &events_backends.Redis[config.EventModel]{},
	}

	Cache        = &cache.Repo[config.CacheModel]{}
	CacheOptions = &cache.Options[config.CacheModel]{
		Backend: &cache_backends.Redis[config.CacheModel]{},
	}

	// ServerOptions = &server.Options{}
	Socket        = &socket.Server[config.CacheModel]{}
	SocketOptions = &socket.Options[config.CacheModel]{
		Cache: Cache,
	}

	Workers        = &workers.Server[config.CacheModel, config.EventModel]{}
	WorkersOptions = &workers.Options[config.CacheModel, config.EventModel]{
		Cache:  Cache,
		Events: Index,
	}
)

// func Initialize(vars *env.Vars) error {
//   err := Initializers.Execute(vars)

//   if err != nil {
//     return err
//   }

//   return nil
// }

func Initialize(vars *env.Vars) error {

	return nil
}
