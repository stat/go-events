package tasks

import (
	"grid/pkg/db/redis"
	"grid/pkg/env"

	"github.com/jinzhu/copier"
)

var (
	AircraftDB *redis.Client
	CacheDB    *redis.Client
)

func Initialize(vars *env.Vars) error {
	// // aircraft db

	// client, err := InitializeClient(vars.RedisDBAircraft, vars)

	// if err != nil {
	//   return err
	// }

	// AircraftDB = client

	// // cache db

	// client, err = InitializeClient(vars.RedisDBCache, vars)

	// if err != nil {
	//   return err
	// }

	// CacheDB = client

	// // success

	return nil
}

func InitializeClient(db int, vars *env.Vars) (*redis.Client, error) {
	options := &redis.Options{}
	copier.Copy(options, vars)

	// force db

	options.RedisDB = db

	// create client

	client, err := redis.NewWithOptions(options)

	if err != nil {
		return nil, err
	}

	return client, nil
}
