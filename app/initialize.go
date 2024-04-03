package app

import (
	_ "grid/pkg/db/postgres"

	"grid/pkg/db/queue"
	"grid/pkg/env"
	"grid/pkg/lifecycle"
	"grid/pkg/repos/cache"
	"grid/pkg/repos/events"
	"grid/pkg/tasks"
	"grid/pkg/transport/http/server"
	"grid/pkg/transport/ws/socket"
	"grid/workers"

	cache_backends "grid/pkg/repos/cache/backends"
	events_backends "grid/pkg/repos/events/backends"
)

var (
	Initializers = lifecycle.Fns[env.Vars]{
		// db client connections

		// postgres.Initialize,
		queue.Initialize,

		// repos

		cache.Initialize[cache_backends.Redis],
		events.Initialize[events_backends.Redis],

		// transport

		server.Initialize,
		socket.Initialize,

		// asynq task registry

		tasks.Initialize,

		// asynq server

		workers.Initialize,
	}
)

func Initialize() (*env.Vars, error) {
	// load env

	vars, err := env.Load()

	if err != nil {
		return nil, err
	}

	return vars, InitializeWithVars(vars)
}

func InitializeWithVars(vars *env.Vars) error {
	return Initializers.Execute(vars)
}
