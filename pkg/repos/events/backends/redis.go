package events_backends

import (
	"context"
	"encoding/json"

	"grid/pkg/db/redis"
	"grid/pkg/env"
	"grid/pkg/models"
	"grid/pkg/repos/events/provider"
	"grid/pkg/utils"
)

type Redis struct {
	*redis.Client
}

func (backend Redis) Initialize(vars *env.Vars) (provider.Provider, error) {
	args := *vars
	args.RedisDB = vars.RedisDBEvents

	concrete := Redis{}
	client, err := redis.NewWithEnv(utils.Ref(args))

	if err != nil {
		return Redis{}, err
	}

	concrete.Client = client

	return concrete, nil
}

func (backend Redis) Append(key string, v interface{}) error {
	cmd := backend.LPush(context.Background(), key, v)

	return cmd.Err()
}

func (backend Redis) Del(key string) error {
	return nil
}

func (backend Redis) DelAtIndex(key string, index int64) error {
	return nil
}

func (backend Redis) DelHead(key string) error {
	cmd := backend.LPop(context.Background(), key)

	return cmd.Err()
}

func (backend Redis) DelTail(key string) error {
	return nil
}

func (backend Redis) Get(key string) (*models.ADSB, error) {
	return nil, nil
}

func (backend Redis) GetAtIndex(key string, index int64) (*models.ADSB, error) {
	cmd := backend.LIndex(context.Background(), key, index)

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	// unmarshal

	event := &models.ADSB{}
	err := json.Unmarshal([]byte(cmd.Val()), event)

	// delete event if mangled

	if err != nil {
		err = backend.DelHead(key)
	}

	// success

	return event, err
}

func (backend Redis) GetHead(key string) (*models.ADSB, error) {
	return backend.GetAtIndex(key, 0)
}

func (backend Redis) GetTail(key string) (*models.ADSB, error) {
	return backend.GetAtIndex(key, -1)
}
