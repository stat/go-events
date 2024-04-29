package backends

import (
	"context"
	"encoding/json"

	"events/pkg/db/redis"
	"events/pkg/env"
	"events/pkg/model"
	"events/pkg/utils"
)

type Redis[V model.Implementer] struct {
	*redis.Client
}

const (
	//TODO: move this to the env
	AircraftLocationsKey = "aircrafts-locations"
)

// func (backend *Redis[V]) Initialize(vars *env.Vars) (provider.Implementer[V], error) {
func (backend *Redis[V]) Initialize(vars *env.Vars) error {
	args := *vars
	args.RedisDB = vars.RedisDBCache

	// concrete := Redis[V]{}
	client, err := redis.NewWithEnv(utils.Ref(args))

	if err != nil {
		// return concrete, err
		return err
	}

	// concrete.Client = client
	backend.Client = client

	// return concrete, nil

	return nil
}

func (backend *Redis[V]) GetLatest(key string) (*V, error) {
	return backend.GetAircraftLocation(key)
}

func (backend *Redis[V]) GetAircraftLocation(key string) (*V, error) {
	cmd := backend.HGet(context.Background(), AircraftLocationsKey, key)
	err := cmd.Err()

	if err != nil {
		return nil, err
	}

	var result V

	value := cmd.Val()

	if err := json.Unmarshal([]byte(value), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (backend *Redis[V]) GetAircraftsLocations() (map[string]*V, error) {
	cmd := backend.HGetAll(context.Background(), AircraftLocationsKey)
	err := cmd.Err()

	if err != nil {
		return nil, err
	}

	result := map[string]*V{}
	value := cmd.Val()

	// TODO: rethink this...

	for k, v := range value {
		var event V

		if err := json.Unmarshal([]byte(v), &event); err != nil {
			return nil, err
		}

		result[k] = &event
	}

	return result, nil
}

func (backend *Redis[V]) UpsertAircraftLocation(key string, v *V) error {
	marshalled, err := json.Marshal(v)

	if err != nil {
		return err
	}

	cmd := backend.HSet(context.Background(), AircraftLocationsKey, key, marshalled)
	return cmd.Err()
}
