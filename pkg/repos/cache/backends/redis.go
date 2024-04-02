package backends

import (
	"context"
	"encoding/json"
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
	args.RedisDB = vars.RedisDBCache

	concrete := Redis{}
	client, err := redis.NewWithEnv(utils.Ref(args))

	if err != nil {
		return concrete, err
	}

	concrete.Client = client

	return concrete, nil
}

func (backend Redis) GetAircraftLocation() (*models.LocationEvent, error) {
	// TODO: implement or remove me
	return nil, nil
}

func (backend Redis) GetAircraftsLocations(key string) (map[string]*models.LocationEvent, error) {
	// TODO: implement or remove me
	return nil, nil
}

func (backend Redis) UpsertAircraftLocation(key string, v *models.LocationEvent) error {
	marshalled, err := json.Marshal(v)

	if err != nil {
		return err
	}

	cmd := backend.HSet(context.Background(), AircraftLocationsKey, key, marshalled)
	return cmd.Err()
}
