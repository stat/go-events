package backends

import (
	"context"
	"grid/pkg/db/redis"
	"grid/pkg/env"
	"grid/pkg/models"
	"grid/pkg/repos/cache/provider"
	"grid/pkg/utils"
)

type Redis struct {
	*redis.Client
}

const (
	//TODO: move this to the env
	AircraftLocationsKey = "aircrafts-locations"
)

func (backend Redis) Initialize(vars *env.Vars) (provider.Provider, error) {
	args := *vars
	args.RedisDB = vars.RedisDBEvents

	concrete := Redis{}
	client, err := redis.NewWithEnv(utils.Ref(args))

	if err != nil {
		return concrete, err
	}

	concrete.Client = client

	return concrete, nil
}

func (backend Redis) GetAircraftLocation() (*models.LocationEvent, error) {
	return nil, nil
}

func (backend Redis) GetAircraftsLocations(key string) (map[string]*models.LocationEvent, error) {
	return nil, nil
}

func (backend Redis) UpsertAircraftLocation(key string, v interface{}) error {
	cmd := backend.HSet(context.Background(), AircraftLocationsKey, key, v)
	return cmd.Err()
}
