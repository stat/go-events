package app

import (
	"events/pkg/env"
	"events/pkg/lifecycle"
)

var (
	Terminators = lifecycle.Fns[env.Vars]{
		// postgres.terminate,
		// redis.terminate,
	}
)

func Terminate(vars *env.Vars) error {
	return Terminators.Execute(vars)
}
