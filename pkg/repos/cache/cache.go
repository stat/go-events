package cache

import (
	"grid/pkg/env"
	"grid/pkg/repos/cache/provider"
)

var (
	Backend provider.Provider
)

func Initialize[T provider.Provider](vars *env.Vars) error {
	var provider T

	concrete, err := provider.Initialize(vars)

	if err != nil {
		return err
	}

	Backend = concrete

	return nil
}
